package hdl

import "testing"

func TestOr(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "Or(0,0)", a: 0, b: 0, want: 0},
		{name: "Or(0,1)", a: 0, b: 1, want: 1},
		{name: "Or(1,0)", a: 1, b: 0, want: 1},
		{name: "Or(1,1)", a: 1, b: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Or(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Or(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
