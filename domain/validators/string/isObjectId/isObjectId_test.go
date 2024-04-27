package isObjectId

import "testing"

func TestIsObjectId(t *testing.T) {
	type args struct {
		objectId string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "should be object id", args: args{objectId: "123456789012345678901234"}, want: true},
		{name: "should be object id", args: args{objectId: "66299f24a183a4f88dfbfdd5"}, want: true},
		{name: "should not be object id", args: args{objectId: "123456789012345678"}, want: false},
		{name: "should not be object id", args: args{objectId: "1234567890123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsObjectId(tt.args.objectId); got != tt.want {
				t.Errorf("IsObjectId() = %v, want %v", got, tt.want)
			}
		})
	}
}
