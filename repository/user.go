package repository

import (
	"context"
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
)

type UserRepositoryDb struct {
	db *sql.DB
}
type UserRepository interface {
	Register(ctx context.Context, user entity.User) (*int, error)
	Login(ctx context.Context, user entity.User) (*entity.User, error)
}

func NewUserRepository(db *sql.DB) *UserRepositoryDb {
	return &UserRepositoryDb{
		db: db,
	}
}

func (r *UserRepositoryDb) Register(ctx context.Context, user entity.User) (*int, error) {
	var id int
	statment := `insert into users(username,email,password,role)values($1,$2,$3,$4) returning id;`
	err := r.db.QueryRowContext(ctx, statment, user.Username, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
func (r *UserRepositoryDb) Login(ctx context.Context, user entity.User) (*entity.User, error) {
	var userData entity.User
	statment := `select id,password,role from users where email=$1;`
	err := r.db.QueryRowContext(ctx, statment, user.Email).Scan(&userData.Id, &userData.Password, &userData.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &userData, nil
}
