# 极简风格前后端分离短网址生成器——server



## 功能TODO

- 长链接转短链接
- 短链接过期时间
- 日志管理
- Dev文档
- API文档
- 后端管理模块
  - 权限系统
  - 系统监控和健康管理
- 用户
  - 注册登录等
  - 用户链接管理
- 高可用
- 高并发
- 微服务


```txt
package main

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/shorturl/model"
)

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}


var db = make(map[string]string)

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
			shortUrl := CreateShortUrl(long.Url)
			c.JSON(http.StatusOK, gin.H{
				"surl": shortUrl.Short,
			})
		}

	})
	r.GET("/url/:short", func(c *gin.Context) {
		user := c.Params.ByName("short")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}

	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
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


```