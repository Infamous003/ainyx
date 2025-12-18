package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Infamous003/ainyx/internal/database/sqlc"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	q sqlc.Querier
}

func NewUser(q sqlc.Querier) *User {
	return &User{q: q}
}

func (r *User) Create(ctx context.Context, name string, dob time.Time) (sqlc.User, error) {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *User) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
	user, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return sqlc.User{}, ErrUserNotFound
		default:
			return sqlc.User{}, err
		}
	}

	return user, nil
}

func (r *User) Update(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error) {
	user, err := r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return sqlc.User{}, ErrUserNotFound
		default:
			return sqlc.User{}, err
		}
	}

	return user, nil
}

func (r *User) Delete(ctx context.Context, id int32) error {
	err := r.q.DeleteUser(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrUserNotFound
		default:
			return err
		}
	}
	return nil
}
