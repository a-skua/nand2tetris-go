package hdl

func Not(in int) (out int) {
	out = Nand(in, in)
	return
}
