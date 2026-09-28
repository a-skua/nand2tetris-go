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

func TestAnd16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Bit
		want [16]builtin.Bit
	}{
		{
			name: "And16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
			a:    [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "And16([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], [1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0])",
			a:    [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Bit{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			want: [16]builtin.Bit{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := And16(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("And16(%v,%v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
