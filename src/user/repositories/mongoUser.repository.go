package scopeRepositories

import (
	"context"
	"fmt"
	"gin-api/domain/utils/copyType"
	"gin-api/src/user/dtos/createUserDto"
	"gin-api/src/user/dtos/updateUserDto"
	"gin-api/src/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{
		collection: collection,
	}
}

func (r *MongoUserRepository) Create(ctx context.Context, user createUserDto.CreateUsersDto) (*models.User, error) {

	// higher the cost, higher the time to process the request
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 9)
	if err != nil {
		fmt.Println("cant genarate hash")
	}
	user.Password = string(hashBytes)

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser := models.User{}
	err = copyType.CopyType(&user, &createdUser)
	if err != nil {
		fmt.Println(err)

		return &createdUser, err
	}

	createdUser.ID = result.InsertedID.(primitive.ObjectID)
	return &createdUser, nil
}

func (r *MongoUserRepository) Update(ctx context.Context, userUpdate updateUserDto.UpdateUsersDto, userId primitive.ObjectID) (models.User, error) {

	if userUpdate.Password != nil {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(*userUpdate.Password), 14)
		if err != nil {
			fmt.Println("cant genarate hash")
		}

		password := string(hashBytes)
		userUpdate.Password = &password
	}

	_, err := r.collection.UpdateByID(ctx, userId, bson.M{"$set": userUpdate})

	updatedUser, err := r.Find(ctx, userId)

	return *updatedUser, err
}

func (r *MongoUserRepository) Find(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindMany(ctx context.Context, page, limit int64) ([]models.User, error) {
	var users []models.User

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

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoUserRepository) Delete(ctx context.Context, id primitive.ObjectID) (bool, error) {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}
	return result.DeletedCount > 0, nil
}
