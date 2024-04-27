package isEmail

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "should be a email",
			email: "example@example.com",
			want:  true,
		},
		{
			name:  "should be a email",
			email: "john@doe.com.br",
			want:  true,
		},
		{
			name:  "should be a email",
			email: "breno.brum2001@gmail.com",
			want:  true,
		},
		{
			name:  "should not be a email",
			email: "example@example..email.com",
			want:  false,
		},
		{
			name:  "should not be a email",
			email: "example is not an email @example.com",
			want:  false,
		},
		{
			name:  "should not be a email",
			email: "example@example not email .com",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.email); got != tt.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
