package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pah-dev/quik-wallet/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Log *LogController

type LogController struct {
	
}

func (w *LogController) CreateLog(c *gin.Context, info map[string]interface{}) {
	detail := "Wallet [" + fmt.Sprint(info["id"])  + "] " + fmt.Sprint(info["operation"]) + " - " + fmt.Sprint(info["amount"])
	log := models.Log{Hash: fmt.Sprint(info["hash"]), Operation: fmt.Sprint(info["operation"]), Amount: info["amount"].(float32), Detail: detail}
	db := c.MustGet("db").(*gorm.DB)
	err := models.LogModel.CreateLog(db, &log)
	if err != nil {
		logrus.Warn("Error saving transaction log")
	}
}
