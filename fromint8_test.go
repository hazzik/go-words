package words

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestFromInt8(t *testing.T) {
	test := func(t *testing.T, input int8, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			t.Helper()
			assert.Equal(t, FromInt8(input), expected)
		})
	}

	test(t, 0, "zero")
	test(t, 1, "one")
	test(t, 2, "two")
	test(t, 19, "nineteen")
	test(t, 20, "twenty")
	test(t, 21, "twenty one")
	test(t, 51, "fifty one")
	test(t, 99, "ninety nine")
	test(t, 100, "one hundred")
	test(t, 101, "one hundred one")
	test(t, 111, "one hundred eleven")
	test(t, 127, "one hundred twenty seven")
	test(t, -127, "minus one hundred twenty seven")
}
