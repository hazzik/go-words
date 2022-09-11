package words

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestFromInt32(t *testing.T) {
	test := func(t *testing.T, input int32, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			assert.Equal(t, FromInt32(input), expected)
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
	test(t, 1111, "one thousand one hundred eleven")
	test(t, 11111, "eleven thousand one hundred eleven")
	test(t, 111111, "one hundred eleven thousand one hundred eleven")
	test(t, 1111111, "one million one hundred eleven thousand one hundred eleven")
	test(t, 11111111, "eleven million one hundred eleven thousand one hundred eleven")
	test(t, 111111111, "one hundred eleven million one hundred eleven thousand one hundred eleven")
	test(t, 1111111111, "one billion one hundred eleven million one hundred eleven thousand one hundred eleven")
	test(t, 2147483647, "two billion one hundred forty seven million four hundred eighty three thousand six hundred forty seven")
	test(t, -2147483647, "minus two billion one hundred forty seven million four hundred eighty three thousand six hundred forty seven")
}
