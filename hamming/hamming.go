package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	bytesA := []rune(a)
	bytesB := []rune(b)
	lenA, lenB := len(bytesA), len(bytesB)
	if lenA != lenB {
		return -1, fmt.Errorf("length valid")
	}
	count := 0
	for index, value := range bytesA {
		if bytesB[index] != value {
			count += 1
		}
	}
	return count, nil
}
