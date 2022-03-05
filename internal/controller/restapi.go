package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"yummlog/internal/model"
)

type RESTController struct {
}

func (f *RESTController) ListFoodPosts(c *gin.Context, params model.ListFoodPostsParams) {
	p := "Parsippany"
	nj := "New Jersey"
	t := time.Now()
	fpl := model.FoodPostsList{
		{
			RestaurantName: "Adayar Ananda Bhavan",
			City:           &p,
			State:          &nj,
			FoodItems: []model.FoodItems{
				{
					Name:       "Full Meals",
					Experience: "like",
				},
				{
					Name:       "Chikoo Milshake",
					Experience: "love",
				},
			},
			Date: &t,
		},
	}
	c.JSON(200, fpl)
}

func (f *RESTController) CreateFoodPost(c *gin.Context, params model.CreateFoodPostParams) {

}
