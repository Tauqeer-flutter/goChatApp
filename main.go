package main

import (
	"fmt"
	"goChatApp/container"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	cont := container.NewContainer()
	cont.SetupRoutes(router)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	server := cont.Config.SERVER
	port := cont.Config.PORT
	fmt.Println(server + ":" + port)
	err := router.Run(server + ":" + port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
