package main

func pow(num, exp int) int {
	if exp == 0 {
		return 1
	}

	for i:= 1; i < exp; i++ {
		num *= num
	}

	return num
}