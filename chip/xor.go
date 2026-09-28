package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

// Xor
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 0
func Xor(a, b builtin.Bit) (out builtin.Bit) {
	out = And(Or(a, b), Not(And(a, b)))
	return
}
