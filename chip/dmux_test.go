package chip

import (
	"github.com/a-skua/nand2tetris/chip/builtin"
	"testing"
)

func TestDMUX(t *testing.T) {
	tests := []struct {
		name         string
		in, sel      builtin.Bit
		wantA, wantB builtin.Bit
	}{
		{name: "DMUX(0,0)", in: 0, sel: 0, wantA: 0, wantB: 0},
		{name: "DMUX(0,1)", in: 0, sel: 1, wantA: 0, wantB: 0},
		{name: "DMUX(1,0)", in: 1, sel: 0, wantA: 1, wantB: 0},
		{name: "DMUX(1,1)", in: 1, sel: 1, wantA: 0, wantB: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := DMUX(tt.in, tt.sel)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("DMUX(%d,%d) = (%d,%d); want (%d,%d)", tt.in, tt.sel, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}
