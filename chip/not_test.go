package chip

import (
	"github.com/a-skua/nand2tetris/chip/builtin"
	"testing"
)

func TestNot(t *testing.T) {
	tests := []struct {
		name string
		in   builtin.Bit
		want builtin.Bit
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

func TestNot16(t *testing.T) {
	tests := []struct {
		name string
		in   [16]builtin.Bit
		want [16]builtin.Bit
	}{
		{
			name: "Not16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0])",
			in:   [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "Not16([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
			in:   [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "Not16([0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1])",
			in:   [16]builtin.Bit{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			want: [16]builtin.Bit{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Not16(tt.in)
			if got != tt.want {
				t.Errorf("Not16(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
