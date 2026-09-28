package hdl

import "github.com/a-skua/nand2tetris/hdl/builtin"

// And
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 0
// 1 | 0 | 0
// 1 | 1 | 1
func And(a, b builtin.Bit) (out builtin.Bit) {
	out = Not(builtin.Nand(a, b))
	return
}
