package postgres

import (
	"context"
	"pulishla/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (p *postRepo) CreatePost(ctx context.Context, req *models.CreatePost) (*models.Post, error) {
	return nil, nil
}

func (p *postRepo) UpdatePost(ctx context.Context, req *models.UpdatePost) (*models.Post, error) {
	return nil, nil
}

func (p *postRepo) DeletePost(ctx context.Context, req *models.DeletePost) error {
	return nil
}

func (p *postRepo) GetListPost(ctx context.Context) (*models.GetListPostResponse, error) {
	return nil, nil
}
