package postgres

import (
	"context"
	"pulishla/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type linkRepo struct {
	db *pgxpool.Pool
}

func NewLinksRepo(db *pgxpool.Pool) *linkRepo {
	return &linkRepo{
		db: db}
}

func (l *linkRepo) CreateLink(ctx context.Context, req *models.CreateLinks)(*models.Links,error){
	return nil,nil
}
