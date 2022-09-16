package user

import (
	"context"

	"github.com/dellta103/s3-golang-assignment/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	if err := user.Insert(ctx, i.db, boil.Whitelist("name", "email", "password", "role")); err != nil {
		return model.User{}, err
	}
	return user, nil
}
