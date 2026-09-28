package logicgate

import "github.com/a-skua/nand2tetris/logicgate/builtin"

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
