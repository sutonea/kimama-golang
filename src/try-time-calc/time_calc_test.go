package timecalc

import (
	"fmt"
	"testing"
	"time"
)

func getTime(str string) time.Time {
	result, e := time.Parse(time.RFC3339, str)
	if e != nil {
		fmt.Println("ERRR")
		fmt.Println(e)
	}
	return result
}
func TestAddDay(t *testing.T) {
	cases := []struct {
		strDate string
		d       int
		want    time.Time
	}{
		{"2020-03-01T00:00:00+09:00", 3, getTime("2020-03-04T00:00:00+09:00")},
		{"2020-03-30T00:00:00+09:00", 2, getTime("2020-04-01T00:00:00+09:00")},
		{"2020-12-29T00:00:00+09:00", 3, getTime("2021-01-01T00:00:00+09:00")},
	}

	for _, c := range cases {
		result := AddDay(c.strDate, c.d)
		if result != c.want {
			t.Errorf("Expected %q, but %q", c.want, result)
		}
	}
}
