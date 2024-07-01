package main

func modifyValueByReference(x *int) {
	*x = 0
}

func main() {
	var num int
	num = 2
	// var ptr *int
	// ptr = &num
	ptr := &num
	println("num : ", num)
	println("ptr : ", ptr)
	println("*ptr : ", *ptr)

	value := 5
	modifyValueByReference(&value)
	println("value : ", value)

}
