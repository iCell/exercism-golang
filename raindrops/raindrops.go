package raindrops

import (
	"fmt"
)

func Convert(num int) string {
	var result string
	if divide3(num) {
		result += "Pling"
	}
	if divide5(num) {
		result += "Plang"
	}
	if divide7(num) {
		result += "Plong"
	}
	if len(result) == 0 {
		result += fmt.Sprintf("%d", num)
	}
	return result
}

func divide3(num int) bool {
	return num % 3 == 0
}

func divide5(num int) bool {
	return num % 5 == 0
}

func divide7(num int) bool {
	return num % 7 == 0
}