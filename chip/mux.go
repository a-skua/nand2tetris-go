package chip

import "github.com/a-skua/nand2tetris/chip/builtin"

// Mux
//
// a | b | sel | out
// 0 | 0 |  0  | 0
// 0 | 1 |  0  | 0
// 1 | 0 |  0  | 1
// 1 | 1 |  0  | 1
// 0 | 0 |  1  | 0
// 0 | 1 |  1  | 1
// 1 | 0 |  1  | 0
// 1 | 1 |  1  | 1
func Mux(a, b, sel builtin.Bit) (out builtin.Bit) {
	out = Or(And(Not(sel), a), And(sel, b))
	return
}

// Mux16
//
// a | b | sel | out
// 0 | 0 |  0  | 0
// 0 | 1 |  0  | 0
// 1 | 0 |  0  | 1
// 1 | 1 |  0  | 1
// 0 | 0 |  1  | 0
// 0 | 1 |  1  | 1
// 1 | 0 |  1  | 0
// 1 | 1 |  1  | 1
func Mux16(a, b [16]builtin.Bit, sel builtin.Bit) (out [16]builtin.Bit) {
	for i := range a {
		out[i] = Mux(a[i], b[i], sel)
	}
	return
}

func Mux4Way16(a, b, c, d [16]builtin.Bit, sel [2]builtin.Bit) (out [16]builtin.Bit) {
	// TODO
	return
}

func Mux8Way16(a, b, c, d, e, f, g, h [16]builtin.Bit, sel [3]builtin.Bit) (out [16]builtin.Bit) {
	// TODO
	return
}
