package logicgate

import (
	"github.com/a-skua/nand2tetris/logicgate/builtin"
	"testing"
)

func TestXor(t *testing.T) {
	tests := []struct {
		name string
		a, b builtin.Bit
		want builtin.Bit
	}{
		{name: "Xor(0,0)", a: 0, b: 0, want: 0},
		{name: "Xor(0,1)", a: 0, b: 1, want: 1},
		{name: "Xor(1,0)", a: 1, b: 0, want: 1},
		{name: "Xor(1,1)", a: 1, b: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Xor(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Xor(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
