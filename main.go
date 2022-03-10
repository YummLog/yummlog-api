package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"yummlog/internal/bootstrap"
	"yummlog/internal/model"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	ctx := context.Background()
	app, err := bootstrap.NewApplication(ctx)
	if err != nil {
		log.Fatalf("error initializing application: %s", err)
	}

	//api := controller.RESTController{}

	router = model.RegisterHandlersWithOptions(router, &app.RESTController, model.GinServerOptions{
		BaseURL: "/v1",
	})

	_ = router.Run(":3000")
}
