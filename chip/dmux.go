package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

// DMux
//
// sel | in | a | b
//
//	0  | 0  | 0 | 0
//	0  | 1  | 1 | 0
//	1  | 0  | 0 | 0
//	1  | 1  | 0 | 1
func DMux(in, sel builtin.Bit) (a, b builtin.Bit) {
	a = And(Not(sel), in)
	b = And(sel, in)
	return
}

func DMux4Way(in builtin.Bit, sel [2]builtin.Bit) (a, b, c, d builtin.Bit) {
	// TODO
	return
}

func DMux8Way(in builtin.Bit, sel [3]builtin.Bit) (a, b, c, d, e, f, g, h builtin.Bit) {
	// TODO
	return
}
