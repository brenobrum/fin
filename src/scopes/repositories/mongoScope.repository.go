package userRepositories

import (
	"context"
	"fmt"
	"gin-api/domain/utils/copyType"
	createScopesDto "gin-api/src/scopes/dtos/createUserDto"
	"gin-api/src/scopes/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

//
//	createdUser.ID = result.InsertedID.(primitive.ObjectID)
//	return &createdUser, nil
//}
//
//func (r *MongoUserRepository) Update(ctx context.Context, userUpdate updateUserDto.UpdateUsersDto, userId primitive.ObjectID) (models.User, error) {
//	_, err := r.collection.UpdateByID(ctx, userId, bson.M{"$set": userUpdate})
//
//	updatedUser, err := r.Find(ctx, userId)
//
//	return *updatedUser, err
//}
//
//func (r *MongoUserRepository) Find(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
//	var user models.User
//	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//
//func (r *MongoUserRepository) FindMany(ctx context.Context, page, limit int64) ([]models.User, error) {
//	var users []models.User
//
//	// Create a FindOptions instance and set the limit and skip
//	findOptions := options.Find()
//	findOptions.SetLimit(limit)
//	findOptions.SetSkip((page - 1) * limit)
//
//	// Passing the find options to the Find method
//	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(ctx)
//
//	if err = cursor.All(ctx, &users); err != nil {
//		return nil, err
//	}
//	return users, nil
//}
//
//func (r *MongoUserRepository) Delete(ctx context.Context, id primitive.ObjectID) (bool, error) {
//	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
//	if err != nil {
//		return false, err
//	}
//	return result.DeletedCount > 0, nil
//}
