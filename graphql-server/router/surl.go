package router

import (
	v1 "github.com/panda8z/shorturl/api/v1"
)

func Start() {
	r := SetupRouter()
	r.POST("/surl", v1.GenShorturl)
	r.POST("/surl/:short", v1.QuaryOriginUrl)
	r.Run(":8080")
}
