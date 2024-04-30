package userRepositories

import (
	"context"
	"gin-api/domain/dtos/pagination"
	createScopesDto "gin-api/src/scopes/dtos/createScopesDto"
	updateScopesDto "gin-api/src/scopes/dtos/updateScopesDto"
	"gin-api/src/scopes/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScopeRepository interface {
	Find(ctx context.Context, id primitive.ObjectID) (*models.Scope, error)
	FindMany(ctx context.Context, pagination pagination.Pagination) ([]models.Scope, error)
	Create(ctx context.Context, scope createScopesDto.CreateScopesDto) (*models.Scope, error)
	Update(ctx context.Context, scopeUpdate updateScopesDto.UpdateScopesDto, userId primitive.ObjectID) (models.Scope, error)
	Delete(ctx context.Context, id primitive.ObjectID) (bool, error)
}
