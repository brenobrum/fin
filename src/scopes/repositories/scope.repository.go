package userRepositories

import (
	"context"
	createScopesDto "gin-api/src/scopes/dtos/createUserDto"
	"gin-api/src/scopes/models"
)

type ScopeRepository interface {
	//Find(ctx context.Context, id primitive.ObjectID) (models.User, error)
	//FindMany(ctx context.Context, page, limit int64) ([]models.User, error)
	Create(ctx context.Context, scope createScopesDto.CreateScopesDto) (*models.Scope, error)
	//Update(ctx context.Context, userUpdate updateUserDto.UpdateUsersDto, userId primitive.ObjectID) (models.User, error)
	//Delete(ctx context.Context, id primitive.ObjectID) (bool, error)
}
