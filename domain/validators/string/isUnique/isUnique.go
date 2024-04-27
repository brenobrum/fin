package isUnique

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// IsUnique checks if the given string is unique in the database
func IsUnique(collection *mongo.Collection, attributeName string, attributeValue string) bool {
	// Define a filter to search for documents with the given attribute value
	filter := bson.M{attributeName: attributeValue}

	// Count the number of documents that match the filter
	count, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return false
	}

	// If count is greater than 0, the attribute value is not unique
	return count == 0
}
