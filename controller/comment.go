package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jheader/golang_blog/config"
	"github.com/jheader/golang_blog/model"
	"github.com/jheader/golang_blog/utils"
)

type CommentController struct{}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

func (com *CommentController) CreateComment(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "User not authenticated")
		return
	}
	postIdString := c.Param("post_id")
	postID, err := strconv.ParseUint(postIdString, 10, 32)
	if err != nil {
		utils.BadRequest(c, "Invalid post ID")
		return
	}
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// 检查文章是否存在
	var post model.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "Post not found")
		return
	}
	comment := model.Comment{
		Content: req.Content,
		UserID:  userID.(uint),
		PostID:  uint(postID),
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		utils.InternalServerError(c, "Failed to create comment")
		return
	}

	// 预加载用户信息
	config.DB.Preload("User").First(&comment, comment.ID)

	utils.Success(c, comment)
}

func (com *CommentController) GetComments(c *gin.Context) {

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
	//文章id
	postIDStr := c.Query("postId")
	if postIDStr == "" {
		utils.BadRequest(c, "文章ID不能为空")
		return

	}
	postId, err := strconv.Atoi(postIDStr)
	if err != nil || size < 0 {
		utils.BadRequest(c, "文章ID格式不对")
		return
	}
	var total int64
	if err := config.DB.Model(&model.Comment{}).Where("post_id = ?", postId).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "统计文章总数失败：" + err.Error(),
			"data": nil,
		})
		return
	}

	//分页查询
	var comments []model.Comment
	if err := config.DB.Scopes(utils.Paginate(page, size)).Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "查询文章列表失败：" + err.Error(),
			"data": nil,
		})
		return
	}

	// 3. 构建分页响应（使用工具类统一格式）
	pageResp := utils.NewPageResponse(comments, total, page, size)

	// 4. 成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": pageResp,
	})

}
