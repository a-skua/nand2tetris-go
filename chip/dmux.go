package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

// DMUX
//
// sel | in | a | b
//
//	0  | 0  | 0 | 0
//	0  | 1  | 1 | 0
//	1  | 0  | 0 | 0
//	1  | 1  | 0 | 1
func DMUX(in, sel builtin.Bit) (a, b builtin.Bit) {
	a = AND(NOT(sel), in)
	b = AND(sel, in)
	return
}

func DMUX4WAY(in builtin.Bit, sel [2]builtin.Bit) (a, b, c, d builtin.Bit) {
	// TODO
	return
}

func DMUX8WAY(in builtin.Bit, sel [3]builtin.Bit) (a, b, c, d, e, f, g, h builtin.Bit) {
	// TODO
	return
}
