package main

import (
	"strconv"
	"testing"
)

func TestIsPalindrom(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{1221, true},
		{1234, false},
		{12321, true},
		{123321, true},
		{132321, false},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			if got := isPalindrom(tc.input); got != tc.want {
				t.Errorf("got %t", got)
			}
		})
	}
}
