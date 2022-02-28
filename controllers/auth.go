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

func (a *AuthController) GetAuth(c *gin.Context) {
	session := configs.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	auth := models.Auth{Username: username, Password: password}
	isExist, err := checkAuth(c, auth)
	if err != nil {
		logrus.Warn(err.Error())
		return
	}
	if !isExist {
		return
	}
	token, err := utils.GenerateToken(username, password)
	if err != nil {
		logrus.Warn(err.Error())
		return
	}
	session.Set("token", token)
	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func checkAuth(c *gin.Context, auth models.Auth) (bool, error) {
	db := c.MustGet("db").(*gorm.DB)
	return models.CheckAuth(db, auth.Username, auth.Password)
}
