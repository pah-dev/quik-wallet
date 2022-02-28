package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pah-dev/quik-wallet/configs"
	"github.com/pah-dev/quik-wallet/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Wallet *WalletController

type WalletController struct {
	
}

type OpForm struct {
    Amount  float32 `form:"amount"`
}

func (w *WalletController) GetWalletByID(c *gin.Context) {
	var wallet models.Wallet
	id := c.Params.ByName("id")
	db := c.MustGet("db").(*gorm.DB)
	err := models.WalletModel.GetWalletByID(db, &wallet, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, wallet)
	}
}

func (w *WalletController) Balance(c *gin.Context){
	var wallet models.Wallet
	session := configs.Default(c)
	id := c.Params.ByName("id")
	db := c.MustGet("db").(*gorm.DB)
	err := models.WalletModel.GetWalletByID(db, &wallet, id)
	if err != nil {
		logrus.Warn(err.Error())
		session.AddFlash(err, "Warn")
		session.Save()
	} else {
		info := gin.H{
			"operation": "Balance",
			"amount": wallet.Balance,
			"id": wallet.ID,
			"hash": wallet.Hash,
		}
		Log.CreateLog(c, info)
		data := gin.H{
			"title":  "Quik Wallets",
			"subtitle":  "BALANCE",
			"operation": "balance",
			"wallet": wallet,
		}
		c.HTML(http.StatusOK, "manage.html", data)
	}
}

func (w *WalletController) Manage(c *gin.Context){
	var wallet models.Wallet
	id := c.Params.ByName("id")
	op := c.Params.ByName("op")
	db := c.MustGet("db").(*gorm.DB)
	err := models.WalletModel.GetWalletByID(db, &wallet, id)
	if err != nil {
		logrus.Warn(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		data := gin.H{
			"title":  "Quik Wallets",
			"subtitle":  "MANAGE",
			"operation": op,
			"wallet": wallet,
		}
		if op != "" {
			c.HTML(http.StatusOK, "operation.html", data)
		}else{
			c.HTML(http.StatusOK, "manage.html", data)
		}
	}
}

func (w *WalletController) Credit(c *gin.Context){
	var wallet models.UpdateWallet
	var oldWallet models.Wallet
	session := configs.Default(c)
	id := c.Params.ByName("id")
	db := c.MustGet("db").(*gorm.DB)
	form := &OpForm{}
	if err := c.ShouldBind(form); err != nil {
		logrus.Warn(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := models.WalletModel.GetWalletByID(db, &oldWallet, id)
	if err != nil {
		session.AddFlash("Wallet not found", "Warn")
		session.Save()
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	wallet = models.UpdateWallet{ Credit: form.Amount, Debit: 0 }
	err = models.WalletModel.CreditWallet(db, &oldWallet, &wallet)
	if err != nil {
		logrus.Warn(err.Error())
		session.AddFlash(err, "Warn")
		session.Save()
	}else{
		 info := gin.H{
			"operation": "Credit",
			"amount": form.Amount,
			"id": oldWallet.ID,
			"hash": oldWallet.Hash,
		}
		Log.CreateLog(c, info)
	}
	data := gin.H{
		"title":  "Quik Wallets",
		"subtitle":  "CREDIT",
		"error": err,
		"operation": "credit",
		"wallet": oldWallet,
		"MsgInfo": session.Flashes("Info"),
		"MsgWarn": session.Flashes("Warn"),
	}
	c.HTML(http.StatusOK, "operation.html", data)
}

func (w *WalletController) Debit(c *gin.Context){
	var wallet models.UpdateWallet
	var oldWallet models.Wallet
	session := configs.Default(c)
	id := c.Params.ByName("id")
	db := c.MustGet("db").(*gorm.DB)
	form := &OpForm{}
	if err := c.ShouldBind(form); err != nil {
		logrus.Warn(err.Error())
		session.AddFlash(err, "Warn")
		session.Save()
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := models.WalletModel.GetWalletByID(db, &oldWallet, id)
	if err != nil {
		logrus.Warn(err.Error())
		session.AddFlash(err, "Warn")
		session.Save()
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	wallet = models.UpdateWallet{ Credit: 0, Debit: form.Amount }
	err = models.WalletModel.DebitWallet(db, &oldWallet, &wallet)
	if err != nil {
		logrus.Warn(err.Error())
		session.AddFlash(err, "Warn")
		session.Save()
	}else{
		info := gin.H{
		   "operation": "Debit",
		   "amount": form.Amount,
		   "wallet": oldWallet,
	   }
	   Log.CreateLog(c, info)
   	}
	data := gin.H{
		"title":  "Quik Wallets",
		"subtitle":  "DEBIT",
		"error": err,
		"operation": "debit",
		"wallet": oldWallet,
		"MsgInfo": session.Flashes("Info"),
		"MsgWarn": session.Flashes("Warn"),
	}
	c.HTML(http.StatusOK, "operation.html", data)
}

func (w *WalletController) ShowAll(c *gin.Context) {
	var wallets []models.Wallet
	session := configs.Default(c)
	db := c.MustGet("db").(*gorm.DB)
	err := models.WalletModel.GetAllWallets(db, &wallets)
	if err != nil {
		logrus.Warn(err.Error())
		session.AddFlash(err, "Warn")
		session.Save()
		c.AbortWithStatus(http.StatusNotFound)
	}
	data := gin.H{
		"title":  "Quik Wallets",
		"wallets": wallets,
	}
	c.HTML(http.StatusOK, "index.html", data)
}
