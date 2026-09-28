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
