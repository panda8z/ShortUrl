package main

import (
	"github.com/panda8z/shorturl/model"
	"github.com/panda8z/shorturl/router"
)

func main() {
	model.InitDb()
	router.Start()
}
