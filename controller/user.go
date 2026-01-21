package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jheader/golang_blog/config"
	"github.com/jheader/golang_blog/model"
	"github.com/jheader/golang_blog/utils"
)

type User struct{}

func (u *User) GetProfile(c *gin.Context) {

	var user model.User
	userID, b := c.GetQuery("userID")
	if b && userID != "" {

		if err := config.DB.First(&user, userID).Error; err != nil {
			utils.UserNotExsit(c)
			return
		}

	} else {
		utils.InternalServerError(c, "userID 为空")
		return
	}

	utils.Success(c, user)

}
