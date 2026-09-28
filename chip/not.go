package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

// Not
//
// in | out
// 0  | 1
// 1  | 0
func Not(in builtin.Bit) (out builtin.Bit) {
	out = builtin.Nand(in, in)
	return
}

// Not16
//
// in | out
// 0  | 1
// 1  | 0
func Not16(in [16]builtin.Bit) (out [16]builtin.Bit) {
	for i := range in {
		out[i] = Not(in[i])
	}
	return
}
