package chip

import (
	"github.com/a-skua/nand2tetris/chip/builtin"
	"testing"
)

func TestAnd(t *testing.T) {
	tests := []struct {
		name string
		a, b builtin.Bit
		want builtin.Bit
	}{
		{name: "And(0,0)", a: 0, b: 0, want: 0},
		{name: "And(0,1)", a: 0, b: 1, want: 0},
		{name: "And(1,0)", a: 1, b: 0, want: 0},
		{name: "And(1,1)", a: 1, b: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := And(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("And(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
