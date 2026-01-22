package controller

import (
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
		//检查是否是自己的文章
		var post model.Post
		postID := *req.ID
		// 用 First 替代 Find（按主键查询单条数据，未找到会返回 ErrRecordNotFound）
		if err := config.DB.Preload("User").First(&post, postID).Error; err != nil {
			utils.BadRequest(c, "not found post id"+strconv.FormatUint(uint64(postID), 10)+"的数据")
			return
		}
		currentUsername, _ := c.Get("current_username")
		userna, _ := currentUsername.(string) // 断言为结构体类型
		if userna != post.User.Username {
			utils.BadRequest(c, "不能修改用户"+post.User.Username+"的数据")
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
	return

}
