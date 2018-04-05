package prime

import (
	"strconv"
	"testing"
)

func TestIs(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{13, true},
		{14, false},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			if got := Is(tc.input); got != tc.want {
				t.Errorf("got %t", got)
			}
		})
	}
}
