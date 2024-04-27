package userRepositories

import (
	"context"
	"gin-api/src/user/dtos/createUserDto"
	"gin-api/src/user/dtos/updateUserDto"
	"gin-api/src/user/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Find(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	FindMany(ctx context.Context, page, limit int64) ([]models.User, error)
	Create(ctx context.Context, user createUserDto.CreateUsersDto) (*models.User, error)
	Update(ctx context.Context, userUpdate updateUserDto.UpdateUsersDto, userId primitive.ObjectID) (models.User, error)
	Delete(ctx context.Context, id primitive.ObjectID) (bool, error)
}
