package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yummlog/internal/db"
	"yummlog/internal/model"
	"yummlog/internal/service"
)

type RESTController struct {
	FoodPostsService service.FoodPostsCRUD
}

func (f *RESTController) ListFoodPosts(c *gin.Context, params model.ListFoodPostsParams) {
	p := "Parsippany"
	nj := "New Jersey"
	usa := "USA"
	t := time.Now()
	id := "121"
	add1 := "1008 Route 36"
	notes := "overall great experience"
	pg := 1
	pgSize := 10
	total := 100
	fpl := model.FoodPostsList{
		FoodPosts: &[]model.FoodPost{
			{
				Id:             &id,
				RestaurantName: "Adayar Ananda Bhavan",
				City:           &p,
				State:          &nj,
				Country:        &usa,
				Date:           &t,
				Address1:       &add1,
				FoodItems: []model.FoodItems{
					{
						Name:       "Full Meals",
						Experience: model.FoodItemsExperienceLike,
					},
					{
						Name:       "Chikoo Milk Shake",
						Experience: model.FoodItemsExperienceFavorite,
					},
					{
						Name:       "Chole Bature",
						Experience: model.FoodItemsExperienceDislike,
					},
				},
				Notes: &notes,
			},
		},
		Page:     &pg,
		PageSize: &pgSize,
		Total:    &total,
	}
	c.JSON(200, fpl)
}

func (f *RESTController) CreateFoodPost(c *gin.Context) {
	var foodPost model.FoodPost

	err := c.BindJSON(&foodPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Message: err.Error(),
		})
	}

	dbFoodPost, err := MapAPIFoodPostToDBFoodpost(foodPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: err.Error(),
		})
	}

	var dbPostDetails []db.Postdetail
	for _, fi := range foodPost.FoodItems {
		pd, pdErr := MapAPIFoodPostToDBPostDetails(fi, dbFoodPost.ID)
		if pdErr != nil {
			c.JSON(http.StatusInternalServerError, model.Error{
				Message: pdErr.Error(),
			})
		}
		dbPostDetails = append(dbPostDetails, pd)
	}

	createdPost, createdPostDetails, err := f.FoodPostsService.CreateFoodPosts(c, dbFoodPost, dbPostDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: err.Error(),
		})
	}

	result, err := MapDBPostAndDetailsToAPIFoodpost(createdPost, createdPostDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, result)

}
