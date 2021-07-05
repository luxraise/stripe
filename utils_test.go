package stripe

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func Test_handleResponse(t *testing.T) {
	type testcase struct {
		value  interface{}
		r      io.Reader
		wanted error
	}

	tcs := []testcase{
		{
			value:  &Customer{},
			r:      strings.NewReader(`{"id" : "foobar" }`),
			wanted: nil,
		},
		{
			value:  &Customer{},
			r:      strings.NewReader(""),
			wanted: errors.New("error encountered while attempting to decode response as JSON: EOF"),
		},
	}

	for _, tc := range tcs {
		if err := handleResponse(tc.r, tc.value); !compareErrors(tc.wanted, err) {
			t.Fatalf("invalid error, expected %v and received %v", tc.wanted, err)
		}
	}
}

func Test_handleError(t *testing.T) {
	type testcase struct {
		r      io.Reader
		wanted error
	}

	tcs := []testcase{
		{
			r:      strings.NewReader(`{"error" : { "message" : "foobar" } }`),
			wanted: errors.New("foobar"),
		},
		{
			r:      strings.NewReader(""),
			wanted: errors.New("error encountered while attempting to decode response as JSON: EOF"),
		},
	}

	for _, tc := range tcs {
		if err := handleError(tc.r); !compareErrors(tc.wanted, err) {
			t.Fatalf("invalid error, expected %v and received %v", tc.wanted, err)
		}
	}
}

func compareErrors(a, b error) bool {
	switch {
	case a == nil && b == nil:
		return true
	case a == nil && b != nil:
		return false
	case a != nil && b == nil:
		return false

	default:
		return a.Error() == b.Error()
	}
}
