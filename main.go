package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goChatApp/container"
	"log"
)

func main() {
	router := gin.Default()

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
