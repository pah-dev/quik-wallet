package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pah-dev/quik-wallet/controllers"
)

func Wallet(e *gin.Engine) {
	
	rAuth := e.Group("/api/v1/auth")
	rAuth.POST("/login", controllers.Auth.GetAuth)

	rWallet := e.Group("/api/v1/wallets")
	// rWallet.Use(midl.JWT())

	rWallet.GET("/", controllers.Wallet.ShowAll)
	rWallet.GET("/:id", controllers.Wallet.GetWalletByID)
	rWallet.GET("/:id/balance", controllers.Wallet.Balance)
	rWallet.GET("/:id/manage", controllers.Wallet.Manage)
	rWallet.GET("/:id/manage/:op", controllers.Wallet.Manage)
	rWallet.POST("/:id/credit", controllers.Wallet.Credit)
	rWallet.POST("/:id/debit", controllers.Wallet.Debit)


}