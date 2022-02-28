package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	config "github.com/pah-dev/quik-wallet/configs"
	route "github.com/pah-dev/quik-wallet/routes"
	util "github.com/pah-dev/quik-wallet/utils"
)

func main() {
	router := SetupRouter()

	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	
	db := config.Connection()
	
	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger())

	router.Static("/static","./public/static")
	router.LoadHTMLGlob("./public/templates/*.html")

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	store := sessions.NewCookieStore([]byte("wallets"))
	router.Use(config.Sessions("wallets", store))

	route.Wallet(router)

	return router
}

