package chip

import (
	"github.com/a-skua/nand2tetris/chip/builtin"
	"testing"
)

func TestOr(t *testing.T) {
	tests := []struct {
		name string
		a, b builtin.Bit
		want builtin.Bit
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

func TestOr16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Bit
		want [16]builtin.Bit
	}{
		{
			name: "Or16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
			a:    [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "Or16([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0])",
			a:    [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "Or16([0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1], [1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0])",
			a:    [16]builtin.Bit{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			b:    [16]builtin.Bit{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			want: [16]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Or16(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Or16(%v,%v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestOr8Way(t *testing.T) {
	tests := []struct {
		name string
		a    [8]builtin.Bit
		want builtin.Bit
	}{
		{
			name: "Or8Way([0,0,0,0,0,0,0,0])",
			a:    [8]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name: "Or8Way([1,1,1,1,1,1,1,1])",
			a:    [8]builtin.Bit{1, 1, 1, 1, 1, 1, 1, 1},
			want: 1,
		},
		{
			name: "Or8Way([0,0,0,0,0,0,0,1])",
			a:    [8]builtin.Bit{0, 0, 0, 0, 0, 0, 0, 1},
			want: 1,
		},
		{
			name: "Or8Way([1,0,0,0,0,0,0,0])",
			a:    [8]builtin.Bit{1, 0, 0, 0, 0, 0, 0, 0},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Or8Way(tt.a)
			if got != tt.want {
				t.Errorf("Or8Way(%v) = %v; want %v", tt.a, got, tt.want)
			}
		})
	}
}
