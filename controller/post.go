package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jheader/golang_blog/config"
	"github.com/jheader/golang_blog/model"
	"github.com/jheader/golang_blog/utils"
)

type PostController struct{}

type PostCreatOrUpdateRequest struct {
	ID      *uint  `json:"postID"` // 指针类型，nil 表示未传
	Title   string `json:"title" binding:"required,min=1"`
	Content string `json:"content" binding:"required,min=1"`
}

func (p *PostController) CreateOrUpdate(c *gin.Context) {

	var req PostCreatOrUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if req.ID != nil {
		postID := *req.ID

		if !checkIsOwner(uint64(postID), c) {
			// 检查不通过，直接返回
			return
		}
		//方式1：使用结构体更新（仅更新非零值字段） 2使用 map 更新（更新所有指定字段，包括零值）
		if config.DB.Model(&model.Post{}).Where("id = ?", postID).Updates(model.Post{
			Title:   req.Title,
			Content: req.Content,
		}).Error != nil {
			utils.BadRequest(c, "更新的数据失败")
			return
		}
		utils.Success(c, "更新数据成功")
		return
	}
	//新增
	user_id, _ := c.Get("user_id")
	newPost := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user_id.(uint),
	}
	//忽略数据中的主键（即使设置了也会生成新主键，除非禁用自增）
	if err := config.DB.Create(&newPost).Error; err != nil {
		utils.BadRequest(c, "新增文章失败："+err.Error())
		return
	}

	utils.Success(c, "新增数据成功")

}

// 查询全部文章 支持分页
func (p *PostController) GetAllPosts(c *gin.Context) {

	var posts []model.Post
	pageStr := c.DefaultQuery("page", "1") // 默认第1页
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // 解析失败或页码无效，默认第1页
	}

	sizeStr := c.DefaultQuery("size", "10") // 默认每页10条
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 || size > 100 {
		size = 10 // 解析失败或条数无效，默认10条（限制最大100）
	}
	var total int64
	// 步骤1：统计总条数（不含 LIMIT/OFFSET）
	if err := config.DB.Model(&model.Post{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "统计文章总数失败：" + err.Error(),
			"data": nil,
		})
		return
	}
	// 步骤2：使用分页中间件查询当前页数据
	if err := config.DB.Preload("User").Preload("Comments").Scopes(utils.Paginate(page, size)).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "查询文章列表失败：" + err.Error(),
			"data": nil,
		})
		return
	}

	// 3. 构建分页响应（使用工具类统一格式）
	pageResp := utils.NewPageResponse(posts, total, page, size)

	// 4. 成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": pageResp,
	})
}

// 查询单个文章的详细信息。
func (p *PostController) GetPostById(c *gin.Context) {

	// 1. 解析路由参数（post_id）
	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文章ID格式错误（必须为数字）",
			"data": nil,
		})
		return
	}
	var post model.Post
	if err := config.DB.Preload("User").Preload("Comments").First(&post, postID).Error; err != nil {
		utils.BadRequest(c, "查询文章失败："+err.Error())
		return
	}

	utils.Success(c, post)

}

func (p *PostController) DeletedById(c *gin.Context) {

	// 1. 解析路由参数（post_id）
	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文章ID格式错误（必须为数字）",
			"data": nil,
		})
		return
	}

	if !checkIsOwner(postID, c) {
		// 检查不通过，直接返回
		return
	}
	var post model.Post
	result := config.DB.Delete(&post, postID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除文章失败：" + result.Error.Error(),
			"data": nil,
		})
		return
	}

	// 检查是否有数据被删除
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "文章不存在或已被删除",
			"data": nil,
		})
		return
	}
	utils.Success(c, "删除数据成功")

}

func checkIsOwner(postID uint64, c *gin.Context) bool {

	//检查是否是自己的文章
	var post model.Post
	// 用 First 替代 Find（按主键查询单条数据，未找到会返回 ErrRecordNotFound）
	if err := config.DB.Preload("User").First(&post, postID).Error; err != nil {
		utils.BadRequest(c, "not found post id"+strconv.FormatUint(postID, 10)+"的数据")
		return false
	}
	currentUsername, _ := c.Get("current_username")
	userna, _ := currentUsername.(string) // 断言为结构体类型
	if userna != post.User.Username {
		utils.BadRequest(c, "不能修改用户"+post.User.Username+"的数据")
		return false
	}
	return true
}
