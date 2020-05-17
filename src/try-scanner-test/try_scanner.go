package tryscannertest

import (
	"strconv"
	"strings"
)

func MyScanner(in string) ([]int, error) {
	sliced := strings.Split(in, " ")
	result := []int{}
	for _, s := range sliced {
		i, e := strconv.Atoi(s)
		if e != nil {
			return result, e
		}
		result = append(result, i)
	}
	return result, nil
}
