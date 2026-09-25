package hdl

import (
	"testing"
)

func TestNand(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "Nand(0,0)", a: 0, b: 0, want: 1},
		{name: "Nand(0,1)", a: 0, b: 1, want: 1},
		{name: "Nand(1,0)", a: 1, b: 0, want: 1},
		{name: "Nand(1,1)", a: 1, b: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Nand(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Nand(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
