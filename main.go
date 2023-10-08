package main

import (
	"github.com/gin-gonic/gin"
	"go-search-info/controller"
)

func main() {
	r := gin.Default()
	r.GET("/search", controller.GetInfo)
	err := r.Run(":8081")
	if err != nil {
		panic(err.Error())
	}
}
