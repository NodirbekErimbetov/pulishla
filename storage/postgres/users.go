package postgres

import (
	"context"
	"pulishla/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}

}

func (u *userRepo) UserCreate(ctx context.Context, req *models.UserCreate) (*models.User, error) {
	return nil, nil
}

func (u *userRepo) GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
	return nil, nil
}

func (u *userRepo) UserUpdate(ctx context.Context, req *models.UserUpdate) (*models.User, error) {
	return nil, nil
}
func (u *userRepo) GetListUser(ctx context.Context) (*models.UserGetListResponse, error) {
	return nil, nil
}
func (u *userRepo) UserDelete(ctx context.Context, req *models.UserPrimaryKey) error {
	return nil
}
