package userRepositories

import (
	"context"
	"fmt"
	"gin-api/domain/dtos/pagination"
	"gin-api/domain/utils/copyType"
	createScopesDto "gin-api/src/scopes/dtos/createScopesDto"
	updateScopesDto "gin-api/src/scopes/dtos/updateScopesDto"
	"gin-api/src/scopes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoScopeRepository struct {
	collection *mongo.Collection
}

func NewMongoScopeRepository(collection *mongo.Collection) *MongoScopeRepository {
	return &MongoScopeRepository{
		collection: collection,
	}
}

func (r *MongoScopeRepository) Create(ctx context.Context, scope createScopesDto.CreateScopesDto) (*models.Scope, error) {
	result, err := r.collection.InsertOne(ctx, scope)
	if err != nil {
		return nil, err
	}

	var createScope models.Scope
	err = copyType.CopyType(&scope, &createScope)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	createScope.ID = result.InsertedID.(primitive.ObjectID)

	return &createScope, nil
}

func (r *MongoScopeRepository) Update(ctx context.Context, scopeUpdate updateScopesDto.UpdateScopesDto, scopeId primitive.ObjectID) (models.Scope, error) {
	_, err := r.collection.UpdateByID(ctx, scopeId, bson.M{"$set": scopeUpdate})

	var scope models.Scope
	found := r.collection.FindOne(ctx, bson.M{"_id": scopeId})
	err = found.Decode(&scope)

	return scope, err
}

func (r *MongoScopeRepository) Find(ctx context.Context, id primitive.ObjectID) (*models.Scope, error) {
	var scope models.Scope
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&scope)
	if err != nil {
		return nil, err
	}
	return &scope, nil
}

func (r *MongoScopeRepository) FindMany(ctx context.Context, pagination pagination.Pagination) ([]models.Scope, error) {
	page := pagination.Page
	limit := pagination.Limit

	var scopes []models.Scope

	// Create a FindOptions instance and set the limit and skip
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)

	// Passing the find options to the Find method
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &scopes); err != nil {
		return nil, err
	}
	return scopes, nil
}

func (r *MongoScopeRepository) Delete(ctx context.Context, id primitive.ObjectID) (bool, error) {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}
	return result.DeletedCount > 0, nil
}
