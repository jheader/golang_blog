package controller

import (
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

}
