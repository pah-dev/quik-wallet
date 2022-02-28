package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pah-dev/quik-wallet/controllers"
	midl "github.com/pah-dev/quik-wallet/middlewares"
)

func Wallet(e *gin.Engine) {
	rPub := e.Group("/public/wallets")
	rPub.GET("/", func(c *gin.Context) { controllers.Wallet.List(c, false) })
	rPub.GET("/:id/manage", controllers.Wallet.Manage)
	rPub.GET("/:id/manage/:op", controllers.Wallet.Manage)
	rPub.GET("/:id/balance",  func(c *gin.Context) { midl.VerifyCache(c, false) }, 
		func(c *gin.Context) { controllers.Wallet.Balance(c, false) })
	rPub.POST("/:id/credit", func(c *gin.Context) { controllers.Wallet.Credit(c, false) })
	rPub.POST("/:id/debit", func(c *gin.Context) { controllers.Wallet.Debit(c, false) })

	rAuth := e.Group("/api/v1/auth")
	rAuth.POST("/login", controllers.Auth.GetAuth)

	rWallet := e.Group("/api/v1/wallets")
	rWallet.Use(midl.JWT())

	rWallet.GET("/", func(c *gin.Context) { controllers.Wallet.List(c, true) })
	rWallet.GET("/:id", controllers.Wallet.GetWalletByID)
	rWallet.GET("/:id/balance",  func(c *gin.Context) { midl.VerifyCache(c, true) }, 
		func(c *gin.Context) { controllers.Wallet.Balance(c, true) })
	rWallet.POST("/:id/credit", func(c *gin.Context) { controllers.Wallet.Credit(c, true) })
	rWallet.POST("/:id/debit", func(c *gin.Context) { controllers.Wallet.Debit(c, true) })


}