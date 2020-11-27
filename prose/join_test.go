package prose

import (
	"errors"
	"fmt"
	"testing"
)

type testCase struct {
	given    []string
	function func(...string) (string, error)
	want     string
}

func errorString(want, got string) string {
	return fmt.Sprintf(`Want "%s", got "%s"`, want, got)
}

func TestZeroAndOneElements(t *testing.T) {
	testCases := []testCase{
		{[]string{}, JoinWithCommas, ""},
		{[]string{""}, JoinWithCommas, ""},
	}
	for _, tc := range testCases {
		res, error := tc.function(tc.given...)
		if error == nil {
			t.Error(errorString(fmt.Sprintf("%s", errors.New("need 2 words or more")), res))
		}
	}
}
func TestTwoElements(t *testing.T) {
	testCases := []testCase{
		{[]string{"hello", "world"}, JoinWithCommas, "hello and world"},
	}
	for _, tc := range testCases {
		res, _ := tc.function(tc.given...)
		if res != tc.want {
			t.Error(errorString(tc.want, res))
		}
	}
}
func TestThreeElements(t *testing.T) {
	testCases := []testCase{
		{[]string{"hello", "hi", "world"}, JoinWithCommas, "hello, hi, and world"},
	}
	for _, tc := range testCases {
		res, _ := tc.function(tc.given...)
		if res != tc.want {
			t.Error(errorString(tc.want, res))
		}
	}
}
