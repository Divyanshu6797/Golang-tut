package main

func main() {
	day := 4

	switch day {
	case 1, 2, 3, 4, 5:
		println("Monday")
	case 6:
		println("Tuesday")
	case 7:
		println("Wednesday")
	default:
		println("Invalid day")
	}

	temperature := 45

	switch {
	case temperature < 0:
		println("very cold")
	case temperature == 0:
		println("cold")
	case temperature > 0:
		println("normal")
	default:
		println("normal")
	}

}
