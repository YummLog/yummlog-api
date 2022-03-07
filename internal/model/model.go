// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package model

import (
	"time"
)

const (
	BasicAuthScopes = "BasicAuth.Scopes"
)

// Defines values for FoodItemsExperience.
const (
	FoodItemsExperienceDislike FoodItemsExperience = "dislike"

	FoodItemsExperienceFavorite FoodItemsExperience = "favorite"

	FoodItemsExperienceLike FoodItemsExperience = "like"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// FoodItems defines model for FoodItems.
type FoodItems struct {
	// like, dislike or loved it
	Experience FoodItemsExperience `json:"experience"`

	// name of the food item
	Name string `json:"name"`
}

// like, dislike or loved it
type FoodItemsExperience string

// FoodPost defines model for FoodPost.
type FoodPost struct {
	Address1       *string     `json:"address1,omitempty"`
	Address2       *string     `json:"address2,omitempty"`
	City           *string     `json:"city,omitempty"`
	Country        *string     `json:"country,omitempty"`
	Date           *time.Time  `json:"date,omitempty"`
	FoodItems      []FoodItems `json:"foodItems"`
	Id             *string     `json:"id,omitempty"`
	Notes          *string     `json:"notes,omitempty"`
	RestaurantName string      `json:"restaurantName"`
	State          *string     `json:"state,omitempty"`
	Zip            *string     `json:"zip,omitempty"`
}

// FoodPostsList defines model for FoodPostsList.
type FoodPostsList struct {
	FoodPosts *[]FoodPost `json:"foodPosts,omitempty"`
	Page      *int        `json:"page,omitempty"`
	PageSize  *int        `json:"pageSize,omitempty"`
	Total     *int        `json:"total,omitempty"`
}

// ListFoodPostsParams defines parameters for ListFoodPosts.
type ListFoodPostsParams struct {
	RestaurantName *string `json:"restaurantName,omitempty"`
}

// CreateFoodPostJSONBody defines parameters for CreateFoodPost.
type CreateFoodPostJSONBody FoodPost

// CreateFoodPostJSONRequestBody defines body for CreateFoodPost for application/json ContentType.
type CreateFoodPostJSONRequestBody CreateFoodPostJSONBody
