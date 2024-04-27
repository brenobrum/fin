package copyType

import (
	"fmt"
	"reflect"
)

func CopyType(source interface{}, target interface{}) error {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	// Check if source and target are pointers
	if sourceValue.Kind() != reflect.Ptr || targetValue.Kind() != reflect.Ptr {
		return fmt.Errorf("both source and target must be pointers")
	}

	// Get the underlying values
	sourceValue = sourceValue.Elem()
	targetValue = targetValue.Elem()

	// Check if the source and target types are structs
	if sourceValue.Kind() != reflect.Struct || targetValue.Kind() != reflect.Struct {
		return fmt.Errorf("both source and target must pointers to structs")
	}

	// Iterate over each field in the source struct
	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceValue.Type().Field(i)
		fieldName := field.Name

		// Find the corresponding field in the target struct
		targetField := targetValue.FieldByName(fieldName)
		if !targetField.IsValid() {
			// does not try to set the unexistent value
			continue
		}

		// Check if the target field is addressable
		if !targetField.CanSet() {
			return fmt.Errorf("target field %s is not addressable", fieldName)
		}

		// Set the value from source to target
		targetField.Set(sourceValue.Field(i))
	}

	return nil
}
