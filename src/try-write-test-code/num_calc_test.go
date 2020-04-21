package trywritetest

import "testing"

func TestNumCalc(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{30, 4, 34},
		{5, -3, 2},
	}

	for _, c := range cases {
		got := MyAdd(c.a, c.b)
		if got != c.want {
			t.Errorf("Expected %q, but %q", c.want, got)
		}
	}
}
