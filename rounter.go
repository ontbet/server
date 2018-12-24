package main

import (
	. "ontbet/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", IndexApi)

	router.GET("/querylastgame", QueryLastInfo)

	router.GET("/queryforaddress/:address", QueryForAddress)

	return router
}
