package maxLength

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
		{name: "should return false to max lenght 6", args: args{n: 6, str: "test string"}, want: false},
		{name: "should return true max lenght 6", args: args{n: 6, str: "test"}, want: true},
		{name: "should return false to max lenght 7", args: args{n: 7, str: "test string"}, want: false},
		{name: "should return true max lenght 10", args: args{n: 10, str: "test 2"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxLength(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("MaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
