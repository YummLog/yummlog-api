package controller

import (
	"database/sql"
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

	if fp.Id == nil || *fp.Id == "" {
		foodPostsUUID, err := internal.GetNewUUID()
		if err != nil {
			return db.Foodpost{}, err
		}
		fp.Id = &foodPostsUUID
	}

	return db.Foodpost{
		ID:             *fp.Id,
		RestaurantName: fp.RestaurantName,
		Address1:       address1,
		Address2:       address2,
		City:           city,
		State:          state,
		Country:        country,
		Zipcode:        zipCode,
		UserID:         sql.NullString{},
		CreatedBy:      sql.NullString{},
		CreatedDate:    sql.NullTime{},
		UpdatedBy:      sql.NullString{},
		UpdatedDate:    sql.NullTime{},
		Notes:          notes,
	}, nil
}

func MapAPIFoodPostToDBPostDetails(pd model.FoodItems, foodPostsId string) (db.Postdetail, error) {
	postDetailsUUID, err := internal.GetNewUUID()
	if err != nil {
		return db.Postdetail{}, err
	}

	return db.Postdetail{
		ID:         postDetailsUUID,
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
	return model.FoodPost{
		Address1:       &foodpost.Address1.String,
		Address2:       &foodpost.Address2.String,
		City:           &foodpost.City.String,
		Country:        &foodpost.Country.String,
		Date:           &foodpost.CreatedDate.Time,
		Id:             &foodpost.ID,
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
