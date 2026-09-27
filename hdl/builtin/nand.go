package builtin

// Nand
//
// a | b | out
// 0 | 0 | 1
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 0
func Nand(a, b Bit) (out Bit) {
	out = NewBit(^(a & b))
	return
}
