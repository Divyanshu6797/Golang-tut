package main

func main() {
	x := 1

	if x == 0 {
		println("x is 0")
	} else if x == 1 {
		println("x is1")
	} else {
		println("x is neithrt 0 nor 1")
	}

	y := 1

	if (y == 0) && (x == 1) {
		println("y is 0 and x is 1")
	}
}
