package postgres

import (
	"context"
	"fmt"
	"log"
	"pulishla/config"
	"pulishla/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	category storage.CategoryRepoI
	comment  storage.CommentRepoI
	links    storage.LinksRepoI
	post     storage.PostRepoI
	users    storage.UserRepoI
	view     storage.ViewsRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("dbname = %s password = %s port = %s host = %s user = %s sslmode=disable",
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
			cfg.PostgresHost,
			cfg.PostgresUser,
		),
	)
	if err != nil {
		return nil, err
	}
	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Error while connecting to database !")
	}

	return &Store{
		db: pgxpool,
	}, nil

}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}
	return s.category
}

func (s *Store) Comment() storage.CommentRepoI {
	if s.comment == nil {
		s.comment = NewCommentRepo(s.db)
	}
	return s.comment
}

func (s *Store) Links() storage.LinksRepoI {
	if s.links == nil {
		s.links = NewLinksRepo(s.db)
	}
	return s.links
}

func (s *Store) Post() storage.PostRepoI {
	if s.post == nil {
		s.post = NewPostRepo(s.db)
	}
	return s.post
}

func (s *Store) User() storage.UserRepoI {
	if s.users == nil {
		s.users = NewUsersRepo(s.db)
	}
	return s.users
}

func (s *Store) Views() storage.ViewsRepoI {
	if s.view == nil {
		s.view = NewViewRepo(s.db)
	}
	return s.view
}
