package service

import (
	"database/sql"
	"yummlog/internal/db"
)

func MapDBFoodPostToDBCreateFoodPostParams(fp db.Foodpost) (db.CreateFoodPostParams, error) {
	return db.CreateFoodPostParams{
		ID:             fp.ID,
		RestaurantName: fp.RestaurantName,
		Address1:       fp.Address1,
		Address2:       fp.Address2,
		City:           fp.City,
		State:          fp.State,
		Country:        fp.Country,
		Zipcode:        fp.Zipcode,
		UserID:         fp.UserID,
		CreatedBy:      fp.CreatedBy,
		CreatedDate:    fp.CreatedDate,
		UpdatedBy:      fp.UpdatedBy,
		UpdatedDate:    fp.UpdatedDate,
		Notes:          fp.Notes,
	}, nil
}

func MapDBPostDetailsToDBCreatePostDetailsParams(pd db.Postdetail) (db.CreatePostDetailsParams, error) {
	return db.CreatePostDetailsParams{
		ID:         pd.ID,
		PostID:     pd.PostID,
		Item:       pd.Item,
		Experience: pd.Experience,
	}, nil
}

func MapStringToSqlNullString(s string) (sql.NullString, error) {
	ns := sql.NullString{}
	err := ns.Scan(s)
	if err != nil {
		return ns, err
	}
	return ns, nil
}
