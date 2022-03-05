package main

import (
	"github.com/gin-gonic/gin"
	"yummlog/internal/controller"
)

func main() {
	router := gin.Default()
	api := controller.FoodPostsService{}

	router = controller.RegisterHandlersWithOptions(router, &api, controller.GinServerOptions{
		BaseURL: "/v1",
	})

	_ = router.Run(":3000")
}
