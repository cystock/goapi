package main

import (
	"./controllers/miapi"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main()  {
	router.GET("/users/:userId", miapi.GetUser)
	router.GET("/sites/:siteId", miapi.GetSite)
	router.GET("/countries/:countryId", miapi.GetCountry)
	router.GET("/result/:userId", miapi.GetResult)
	router.GET("/resultgoroutine/:userId", miapi.GetResultGoroutine)
	router.GET("/resultchannel/:userId",miapi.GetResultChannel)

	router.Run(port)
}