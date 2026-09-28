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

// Xor16
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 0
func Xor16(a, b [16]builtin.Bit) (out [16]builtin.Bit) {
	for i := range a {
		out[i] = Xor(a[i], b[i])
	}
	return
}
