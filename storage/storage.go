package storage

import (
	"context"
	"pulishla/models"
)

type StorageI interface {
	Category() CategoryRepoI
	Comment() CommentRepoI
	Links() LinksRepoI
	Post() PostRepoI
	User() UserRepoI
	Views() ViewsRepoI
}

type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *models.CreateCategory) (*models.Category, error)
	UpdateCategory(ctx context.Context, req *models.UptadeCategory) (*models.Category, error)
	DeleteCategory(ctx context.Context, req *models.DeleteCategory) error
}

type CommentRepoI interface {
	CreateComment(ctx context.Context, req *models.CreateComment) (*models.Comments, error)
}

type LinksRepoI interface {
	CreateLink(ctx context.Context, req *models.CreateLinks) (*models.Links, error)
}

type PostRepoI interface {
	CreatePost(ctx context.Context, req *models.CreatePost) (*models.Post, error)
	UpdatePost(ctx context.Context, req *models.UpdatePost) (*models.Post, error)
	DeletePost(ctx context.Context, req *models.DeletePost) error
	GetListPost(ctx context.Context) (*models.GetListPostResponse, error)
}

type UserRepoI interface {
	UserCreate(ctx context.Context, req *models.UserCreate) (*models.User, error)
	UserUpdate(ctx context.Context, req *models.UserUpdate) (*models.User, error)
	UserDelete(ctx context.Context, req *models.UserDelete) error
	GetListUser(ctx context.Context) (*models.UserGetListResponse, error)
}
type ViewsRepoI interface {
	CreateView(ctx context.Context, req *models.CreateView) (*models.Views, error)
}
