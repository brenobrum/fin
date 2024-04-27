package randoString

import "testing"

func TestRandoString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "should have 6 chars", args: args{n: 6}, want: 6},
		{name: "should have 11 chars", args: args{n: 11}, want: 11},
		{name: "should have 14 chars", args: args{n: 15}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandoString(tt.args.n); len(got) != tt.want {
				t.Errorf("RandoString() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
