package isNotEmpty

import "testing"

func TestIsNotEmpty(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "Should be empty",
			str:  "",
			want: false,
		},
		{
			name: "Should not be empty",
			str:  "not empty",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.str); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
