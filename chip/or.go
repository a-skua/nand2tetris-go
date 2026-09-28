package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

// Or
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 1
func Or(a, b builtin.Bit) (out builtin.Bit) {
	out = Not(And(Not(a), Not(b)))
	return
}

// Or16
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 1
func Or16(a, b [16]builtin.Bit) (out [16]builtin.Bit) {
	for i := range a {
		out[i] = Or(a[i], b[i])
	}
	return
}

// Or8Way
//
// a | b | c | d | e | f | g | h | out
// 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0
// 0 | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 1
// 1 | 1 | 1 | 1 | 1 | 1 | 1 | 1 | 1
func Or8Way(a [8]builtin.Bit) (out builtin.Bit) {
	// TODO
	return
}
