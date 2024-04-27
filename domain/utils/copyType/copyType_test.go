package copyType

import (
	"reflect"
	"testing"
)

type NameID struct {
	Name string
	ID   string
}

type Name struct {
	Name string
}

func TestCopyType(t *testing.T) {
	name1 := Name{Name: "name"}
	nameId1 := NameID{Name: "", ID: "ID"}
	wantNameID1 := NameID{Name: "name", ID: ""}

	type args struct {
		source interface{}
		target interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
	}{
		{
			name: "Name should match NameID",
			args: args{
				source: &name1,
				target: &nameId1,
			},
			wantValue: wantNameID1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyType(tt.args.source, tt.args.target)

			// Iterate over fields and compare values
			sourceValue := reflect.ValueOf(tt.args.source)
			targetValue := reflect.ValueOf(tt.args.target)
			for i := 0; i < sourceValue.Elem().NumField(); i++ {
				sourceFieldValue := sourceValue.Elem().Field(i)
				targetFieldValue := targetValue.Elem().Field(i)

				if !reflect.DeepEqual(sourceFieldValue.Interface(), targetFieldValue.Interface()) {
					t.Errorf("Field %s: got %v, want %v", sourceValue.Elem().Type().Field(i).Name, targetFieldValue.Interface(), sourceFieldValue.Interface())
				}
			}

		})
	}
}
