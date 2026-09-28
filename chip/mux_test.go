package chip

import (
	"github.com/a-skua/nand2tetris/chip/builtin"
	"testing"
)

func TestMux(t *testing.T) {
	tests := []struct {
		name      string
		a, b, sel builtin.Bit
		want      builtin.Bit
	}{
		{name: "Mux(0,0,0)", a: 0, b: 0, sel: 0, want: 0},
		{name: "Mux(0,1,0)", a: 0, b: 1, sel: 0, want: 0},
		{name: "Mux(1,0,0)", a: 1, b: 0, sel: 0, want: 1},
		{name: "Mux(1,1,0)", a: 1, b: 1, sel: 0, want: 1},
		{name: "Mux(0,0,1)", a: 0, b: 0, sel: 1, want: 0},
		{name: "Mux(0,1,1)", a: 0, b: 1, sel: 1, want: 1},
		{name: "Mux(1,0,1)", a: 1, b: 0, sel: 1, want: 0},
		{name: "Mux(1,1,1)", a: 1, b: 1, sel: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mux(tt.a, tt.b, tt.sel)
			if got != tt.want {
				t.Errorf("Mux(%d,%d,%d) = %d; want %d", tt.a, tt.b, tt.sel, got, tt.want)
			}
		})
	}
}

func TestMux16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Bit
		sel  builtin.Bit
		want [16]builtin.Bit
	}{
		{
			name: "Mux16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 0)",
			a:    [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			sel:  0,
			want: [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "Mux16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 1)",
			a:    [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			sel:  1,
			want: [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mux16(tt.a, tt.b, tt.sel)
			if got != tt.want {
				t.Errorf("Mux16(%v,%v,%d) = %v; want %v", tt.a, tt.b, tt.sel, got, tt.want)
			}
		})
	}
}
