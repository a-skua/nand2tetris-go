package hdl

import "github.com/a-skua/nand2tetris/hdl/builtin"

// Not
//
// in | out
// 0  | 1
// 1  | 0
func Not(in builtin.Bit) (out builtin.Bit) {
	out = builtin.Nand(in, in)
	return
}
