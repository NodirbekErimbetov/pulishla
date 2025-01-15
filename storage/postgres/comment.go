package postgres

import (
	"context"
	"pulishla/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type commentRepo struct {
	db *pgxpool.Pool
}

func NewCommentRepo(db *pgxpool.Pool) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) CreateComment(ctx context.Context, req *models.CreateComment) (*models.Comments, error) {
	return nil, nil
}
