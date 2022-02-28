package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	config "github.com/pah-dev/quik-wallet/configs"
	route "github.com/pah-dev/quik-wallet/routes"
	"github.com/pah-dev/quik-wallet/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	router := SetupRouter()

	logrus.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	
	db := config.Connection()
	
	cache := redis.NewClient(&redis.Options{
		Addr: utils.GodotEnv("REDIS_URL"),
	})

	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger())

	router.Static("/static","./public/static")
	router.LoadHTMLGlob("./public/templates/*.html")

	if utils.GodotEnv("GO_ENV") != "production" && utils.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if utils.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("cache", cache)
	})
	store := sessions.NewCookieStore([]byte("wallets"))
	router.Use(config.Sessions("wallets", store))

	route.Wallet(router)

	return router
}

