package tryscannertest

import (
	"bufio"
	"strconv"
	"strings"
)

func MyScanner(in *bufio.Scanner) ([]int, error) {
	in.Scan()
	sliced := strings.Split(in.Text(), " ")
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
