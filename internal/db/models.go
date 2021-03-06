// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Foodpost struct {
	ID             string
	RestaurantName string
	Address1       sql.NullString
	Address2       sql.NullString
	City           sql.NullString
	State          sql.NullString
	Country        sql.NullString
	Zipcode        sql.NullString
	UserID         sql.NullString
	CreatedBy      sql.NullString
	CreatedDate    sql.NullTime
	UpdatedBy      sql.NullString
	UpdatedDate    sql.NullTime
	Notes          sql.NullString
}

type Postdetail struct {
	ID         string
	PostID     string
	Item       string
	Experience string
}
