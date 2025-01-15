package postgres

import (
	"context"
	"pulishla/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) CreateCategory(ctx context.Context, req *models.CreateCategory) (*models.Category, error) {

	return nil, nil
}

func (c *categoryRepo) UpdateCategory(ctx context.Context, req *models.UptadeCategory) (*models.Category, error) {
    return nil, nil
}

func (c *categoryRepo) DeleteCategory(ctx context.Context,req *models.DeleteCategory) error {
    return nil
}
