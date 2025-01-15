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
	CreateComment(ctx context.Context, req *models.CreateComment) (*models.Comment, error)
	DeleteComment(ctx context.Context, req *models.DeleteComment) error
	GetListComments(ctx context.Context) (*models.GetListCommentResponse, error)
}

type LinksRepoI interface {
	CreateLink(ctx context.Context, req *models.CreateLinks) (*models.Links, error)
	UpdateLink(ctx context.Context, req *models.UpdateLinks) (*models.Links, error)
	DeleteLinks(ctx context.Context, req *models.DeleteLinks) error
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
	GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetListUser(ctx context.Context) (*models.UserGetListResponse, error)
	UserDelete(ctx context.Context, req *models.UserPrimaryKey) error
}
type ViewsRepoI interface {
	CreateView(ctx context.Context, req *models.CreateView) (*models.Views, error)
	UpdateView(ctx context.Context, req *models.UpdateView) (*models.Views, error)
	GetByIdView(ctx context.Context,req *models.PrimaryKeyView)(*models.Views,error)
	GetListView(ctx context.Context)(*models.GetListViewResponse,error)	
}
