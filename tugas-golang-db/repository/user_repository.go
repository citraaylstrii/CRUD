package repository

import (
	"context"
	"golang-database-user/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error)
	UpdateUser(ctx context.Context, user model.MstUser, userId string) (model.MstUser, error)
	ReadUser(ctx context.Context) ([]model.MstUser, error)
	DeleteUser(ctx context.Context, userId string) error}
