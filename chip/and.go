package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

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

// And16
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 0
// 1 | 0 | 0
// 1 | 1 | 1
func And16(a, b [16]builtin.Bit) (out [16]builtin.Bit) {
	for i := range a {
		out[i] = And(a[i], b[i])
	}
	return
}
