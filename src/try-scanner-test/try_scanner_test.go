package tryscannertest

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {

	type pattern struct {
		in      string
		outList []int
	}

	patterns := []pattern{
		{"777", []int{777}},
	}

	for _, p := range patterns {
		result, e := MyScanner(bufio.NewScanner(strings.NewReader(p.in)))
		if e != nil {
			t.Errorf("Unexpected Error")
		}
		i := 0
		for _, expect := range p.outList {
			got := result[i]
			if got != expect {
				t.Errorf("Expected %d, but %d", expect, got)
			}
		}
	}
}
