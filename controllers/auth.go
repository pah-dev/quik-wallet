package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pah-dev/quik-wallet/configs"
	"github.com/pah-dev/quik-wallet/models"
	"github.com/pah-dev/quik-wallet/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Auth *AuthController

type AuthController struct {
	
}

type User struct {
	Username	string `gorm:"type:varchar(255);unique;not null" json:"username" form:"username"`
	Password  	string `gorm:"type:varchar(255);not null" json:"password" form:"password"`
}


func (a *AuthController) GetAuth(c *gin.Context) {
	session := configs.Default(c)
	code := http.StatusOK
	msg := ""
	token := ""
	user := &User{}
	if err := c.ShouldBind(user); err != nil {
		logrus.Warn(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	auth := models.Auth{Username: user.Username, Password: user.Password}
	isExist, err := checkAuth(c, auth)
	if err != nil {
		logrus.Warn(err.Error())
		code = http.StatusInternalServerError
		msg = "Internal server error"
	}else{
		if !isExist {
			code = http.StatusNotFound
			msg = "User not found, incorrect credentials"
		}else{
			newToken, err := utils.GenerateToken(auth.Username, auth.Password)
			if err != nil {
				logrus.Warn(err.Error())
				code = http.StatusInternalServerError
				msg = "Internal server error"
			}
			session.Set("token", newToken)
			token = newToken
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
		"token": token,
	})
}

func checkAuth(c *gin.Context, auth models.Auth) (bool, error) {
	db := c.MustGet("db").(*gorm.DB)
	return models.CheckAuth(db, auth.Username, auth.Password)
}
