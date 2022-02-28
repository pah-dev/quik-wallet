package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/pah-dev/quik-wallet/models"
)


func VerifyCache(c *gin.Context, api bool) {
	id := c.Params.ByName("id")
	cache := c.MustGet("cache").(*redis.Client)
	redisWallet, err := cache.Get(c, id).Bytes()
	if err != nil {
		c.Next()
	}else{
		var wallet models.Wallet
		json.Unmarshal(redisWallet, &wallet)
		if api {
			c.JSON(http.StatusOK, wallet)
		}else{
			data := gin.H{
				"title":  "Quik Wallets",
				"subtitle":  "BALANCE",
				"operation": "balance",
				"wallet": wallet,
			}
			c.HTML(http.StatusOK, "manage.html", data)
		}
		c.Abort()
	}
}