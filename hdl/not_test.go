package hdl

import (
	"testing"
)

func TestNot(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{name: "Not(0)", in: 0, want: 1},
		{name: "Not(1)", in: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Not(tt.in)
			if got != tt.want {
				t.Errorf("Not(%d) = %d; want %d", tt.in, got, tt.want)
			}
		})
	}
}
