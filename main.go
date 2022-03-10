package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"yummlog/internal/controller"
	"yummlog/internal/model"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	api := controller.RESTController{}

	router = model.RegisterHandlersWithOptions(router, &api, model.GinServerOptions{
		BaseURL: "/v1",
	})

	_ = router.Run(":3000")
}
