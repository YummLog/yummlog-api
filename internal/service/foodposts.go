package service

import (
	"context"

	"yummlog/internal/db"
)

type FoodPostsCRUD interface {
	ListFoodPosts(ctx context.Context) ([]db.ListFoodPostsRow, error)
	CreateFoodPosts(ctx context.Context, fp db.Foodpost, pd []db.Postdetail) (db.Foodpost, []db.Postdetail, error)
}

type FoodPostsService struct {
	//needs reader connection
	Reader *db.Queries
	//needs writer connection
	Writer *db.Queries
}

func NewFoodPostsCRUD(ctx context.Context, reader *db.Queries, writer *db.Queries) FoodPostsCRUD {
	return &FoodPostsService{
		Reader: reader,
		Writer: writer,
	}
}

func (fps *FoodPostsService) ListFoodPosts(ctx context.Context) ([]db.ListFoodPostsRow, error) {
	return fps.Reader.ListFoodPosts(ctx)
}

func (fps *FoodPostsService) CreateFoodPosts(ctx context.Context, fp db.Foodpost, pd []db.Postdetail) (db.Foodpost, []db.Postdetail, error) {

	cfp, err := MapDBFoodPostToDBCreateFoodPostParams(fp)
	if err != nil {
		return db.Foodpost{}, nil, err
	}

	createdFoodPost, err := fps.Writer.CreateFoodPost(ctx, cfp)
	if err != nil {
		return db.Foodpost{}, nil, err
	}

	//var createdPostDetails []db.Postdetail
	//for _, postDetail := range pd {
	//	a, e := MapDBPostDetailsToDBCreatePostDetailsParams(postDetail)
	//	if e != nil {
	//		return db.Foodpost{}, nil, err
	//	}
	//
	//	a.PostID = createdFoodPost.ID
	//	cpd, e := fps.Writer.CreatePostDetails(ctx, a)
	//	if e != nil {
	//		return db.Foodpost{}, nil, err
	//	}
	//	createdPostDetails = append(createdPostDetails, cpd)
	//}

	return createdFoodPost, nil, nil
}
