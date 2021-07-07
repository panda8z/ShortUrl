package router

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/shorturl/model"
)

var db = make(map[string]string)

// TODO: test model
type LongUrl struct {
	Url string `json:"url"`
}

func CreateShortUrl(long string) model.Url {
	return model.Url{
		Origin: long,
		Short:  base64.StdEncoding.EncodeToString([]byte(long)),
	}
}
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	r.POST("/url", func(c *gin.Context) {
		// long := &LongUrl{}
		// c.ShouldBindJSON(long)
		long := LongUrl{}
		if c.Bind(&long) == nil {
			// 1. create short url
			// 2. quary is exist
			// 3. exist return false
			// 4. unexist save Url Model return Url.Short
			shortUrl := CreateShortUrl(long.Url)
			c.JSON(http.StatusOK, gin.H{
				"surl": shortUrl.Short,
			})
		}

	})

	r.GET("/url/:short", func(c *gin.Context) {
		short := c.Params.ByName("short")
		// 1. quary to check.
		c.JSON(http.StatusOK, gin.H{"url": "simple=" + short})

	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		var json struct {
			Value string `json:"value" binding:"required"`
		}
		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
	return r
}
