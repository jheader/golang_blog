package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jheader/golang_blog/config"
	"github.com/jheader/golang_blog/model"
	"github.com/jheader/golang_blog/utils"
)

type AuthController struct{}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (ac *AuthController) Register(c *gin.Context) {

	var req RegisterRequest

	err := c.Bind(&req)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	//查询用户是否存在
	u, err := model.GetUserByUsername(req.Username, config.DB)
	if err == nil {
		utils.BadRequest(c, *&u.Username+"already exists")
		return
	}

	var existingUser model.User
	// 检查邮箱是否已存在
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.BadRequest(c, "Email already exists")
		return
	}

	// 创建新用户
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // 密码会在BeforeCreate钩子中自动加密
	}

	errorMsgs := (&user).Save(config.DB)
	if errorMsgs != nil {
		utils.InternalServerError(c, errorMsgs.Error())
		return
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "Failed to generate token")
		return
	}

	utils.Success(c, map[string]any{
		"Token": token,
		"User":  user,
	})

}
