package hdl

func Nand(a, b int) (out int) {
	out = ^(a & b) & 1
	return
}
