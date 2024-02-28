package controller

import (
	"database/sql"
	"time"

	uuid2 "github.com/google/uuid"
	"yummlog/internal"
	"yummlog/internal/db"
	"yummlog/internal/model"
	"yummlog/internal/service"
)

func MapAPIFoodPostToDBFoodpost(fp model.FoodPost) (db.Foodpost, error) {

	address1, err := service.MapStringToSqlNullString(*fp.Address1)
	if err != nil {
		return db.Foodpost{}, err
	}

	address2, err := service.MapStringToSqlNullString(*fp.Address2)
	if err != nil {
		return db.Foodpost{}, err
	}

	city, err := service.MapStringToSqlNullString(*fp.City)
	if err != nil {
		return db.Foodpost{}, err
	}

	state, err := service.MapStringToSqlNullString(*fp.State)
	if err != nil {
		return db.Foodpost{}, err
	}

	zipCode, err := service.MapStringToSqlNullString(*fp.Zip)
	if err != nil {
		return db.Foodpost{}, err
	}

	country, err := service.MapStringToSqlNullString(*fp.Country)
	if err != nil {
		return db.Foodpost{}, err
	}

	notes, err := service.MapStringToSqlNullString(*fp.Notes)
	if err != nil {
		return db.Foodpost{}, err
	}

	if fp.Id == nil {
		foodPostsUUID, err := internal.GetNewUUID()
		if err != nil {
			return db.Foodpost{}, err
		}
		fp.Id = &foodPostsUUID
	}

	if fp.Date == nil {
		t := time.Now()
		fp.Date = &t
	}

	return db.Foodpost{
		ID:             fp.Id.String(),
		RestaurantName: fp.RestaurantName,
		Address1:       address1,
		Address2:       address2,
		City:           city,
		State:          state,
		Country:        country,
		Zipcode:        zipCode,
		UserID:         sql.NullString{},
		CreatedBy:      sql.NullString{},
		CreatedDate: sql.NullTime{
			Time:  *fp.Date,
			Valid: true,
		},
		UpdatedBy: sql.NullString{},
		UpdatedDate: sql.NullTime{
			Time:  *fp.Date,
			Valid: true,
		},
		Notes: notes,
	}, nil
}

func MapAPIFoodPostToDBPostDetails(pd model.FoodItems, foodPostsId string) (db.Postdetail, error) {
	postDetailsUUID, err := internal.GetNewUUID()
	if err != nil {
		return db.Postdetail{}, err
	}

	return db.Postdetail{
		ID:         postDetailsUUID.String(),
		PostID:     foodPostsId,
		Item:       pd.Name,
		Experience: string(pd.Experience),
	}, nil
}

func MapDBPostAndDetailsToAPIFoodpost(foodpost db.Foodpost, postDetails []db.Postdetail) (model.FoodPost, error) {
	var apiFoodPost model.FoodPost

	apiFoodPost, err := MapDBFoodPostToAPIFoodPost(foodpost)
	if err != nil {
		return apiFoodPost, nil
	}

	var apiPostDetails []model.FoodItems
	for _, fooditem := range postDetails {
		apd, apdErr := MapDBPostDetailsToAPIPostDetails(fooditem)
		if apdErr != nil {
			return apiFoodPost, apdErr
		}
		apiPostDetails = append(apiPostDetails, apd)
	}
	apiFoodPost.FoodItems = apiPostDetails

	return apiFoodPost, nil
}

func MapDBFoodPostToAPIFoodPost(foodpost db.Foodpost) (model.FoodPost, error) {

	uuid, err := uuid2.Parse(foodpost.ID)
	if err != nil {
		return model.FoodPost{}, err
	}

	return model.FoodPost{
		Address1:       &foodpost.Address1.String,
		Address2:       &foodpost.Address2.String,
		City:           &foodpost.City.String,
		Country:        &foodpost.Country.String,
		Date:           &foodpost.CreatedDate.Time,
		Id:             &uuid,
		Notes:          &foodpost.Notes.String,
		RestaurantName: foodpost.RestaurantName,
		State:          &foodpost.State.String,
		Zip:            &foodpost.Zipcode.String,
	}, nil
}

func MapDBPostDetailsToAPIPostDetails(postDetails db.Postdetail) (model.FoodItems, error) {
	return model.FoodItems{
		Experience: model.FoodItemsExperience(postDetails.Experience),
		Name:       postDetails.Item,
	}, nil
}

func MapListFoodPostsRowToAPIFoodPost(listFoodPosts []db.ListFoodPostsRow) (*[]model.FoodPost, error) {

	idToPostMap := make(map[string]*model.FoodPost)
	var foodPosts []model.FoodPost

	for _, foodPostsRow := range listFoodPosts {
		foodItem := model.FoodItems{
			Experience: model.FoodItemsExperience(foodPostsRow.Experience.String),
			Name:       foodPostsRow.Item.String,
		}
		var foodPost *model.FoodPost
		var ok bool

		u, err := uuid2.Parse(foodPostsRow.ID)
		if err != nil {
			return nil, err
		}

		//check if record was already created
		if foodPost, ok = idToPostMap[foodPostsRow.ID]; !ok {
			foodPost = &model.FoodPost{
				Id:             &u,
				Address1:       &foodPostsRow.Address1.String,
				Address2:       &foodPostsRow.Address2.String,
				City:           &foodPostsRow.City.String,
				Country:        &foodPostsRow.Country.String,
				Date:           &foodPostsRow.CreatedDate.Time,
				Notes:          &foodPostsRow.Notes.String,
				RestaurantName: foodPostsRow.RestaurantName,
				State:          &foodPostsRow.State.String,
				Zip:            &foodPostsRow.Zipcode.String,
				FoodItems:      []model.FoodItems{foodItem},
			}
			idToPostMap[foodPostsRow.ID] = foodPost
		} else {
			foodPost.FoodItems = append(foodPost.FoodItems, foodItem)
		}

	}

	for _, post := range idToPostMap {
		foodPosts = append(foodPosts, *post)
	}

	return &foodPosts, nil
}
