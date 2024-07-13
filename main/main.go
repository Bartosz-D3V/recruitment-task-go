package main

import (
	"fmt"
	appConfig "github.com/Bartosz-D3V/recruitment-task-go/config"
	"github.com/Bartosz-D3V/recruitment-task-go/service"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	config := appConfig.New()

	engine := gin.New()
	searchSvc := service.New(config)

	log := config.Logger
	log.Info("Application starting")

	engine.GET("/endpoint/:number", func(context *gin.Context) {
		log.Debug("Endpoint /endpoint/:number called")
		HandleGetNumber(searchSvc, context)
	})

	err := engine.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
