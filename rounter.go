package main

import (
	. "ontbet/api"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/querylastgame", QueryLastInfo)

	router.GET("/queryforaddress/:address", QueryForAddress)

	return router
}
