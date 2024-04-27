package minLength

import (
	"testing"
)

func TestMaxLength(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "should return true to max lenght 6", args: args{n: 6, str: "test string"}, want: true},
		{name: "should return false max lenght 6", args: args{n: 6, str: "test"}, want: false},
		{name: "should return true to max lenght 7", args: args{n: 7, str: "test string"}, want: true},
		{name: "should return false max lenght 10", args: args{n: 10, str: "test 2"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinLength(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("MinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
