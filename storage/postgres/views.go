package postgres

import (
	"context"
	"pulishla/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type viewRepo struct {
	db *pgxpool.Pool
}

func NewViewRepo(db *pgxpool.Pool) *viewRepo {
	return &viewRepo{
		db: db,
	}
}

func (v *viewRepo) CreateView(ctx context.Context, req *models.CreateView) (*models.Views, error) {
	return nil, nil
}
