package isObjectId

import "go.mongodb.org/mongo-driver/bson/primitive"

// IsObjectId validates if a string is empty
func IsObjectId(objectId string) bool {
	_, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return false
	}

	return true
}
