package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"yummlog/internal/model"
)

type FoodPostsService struct {
}

func (f *FoodPostsService) ListFoodPosts(c *gin.Context, params model.ListFoodPostsParams) {
	p := "Parsippany"
	nj := "New Jersey"
	t := time.Now()
	fpl := model.FoodPostsList{
		{
			RestaurantName: "Adayar Ananda Bhavan",
			City:           &p,
			State:          &nj,
			FoodItems: []string{
				"Full Meals",
				"Chikoo Milkshake",
			},
			Date:       &t,
			Experience: "Yummy",
		},
	}
	c.JSON(200, fpl)
}

func (f *FoodPostsService) CreateFoodPost(c *gin.Context, params model.CreateFoodPostParams) {

}
