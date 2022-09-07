package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct{}

func main() {
	router := gin.Default()
	api := router.Group("api")
	v1 := api.Group("v1")
	setupV1Api(v1)
	router.Run(":8080")
}

func setupV1Api(v1 *gin.RouterGroup) {
	setupShoppingCart(v1)
}

func setupShoppingCart(v1 *gin.RouterGroup) {}
