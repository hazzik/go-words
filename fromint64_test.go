package words

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestFromInt64(t *testing.T) {
	test := func(t *testing.T, input int64, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			assert.Equal(t, FromInt64(input), expected)
		})
	}

	test(t, 0, "zero")
	test(t, 1, "one")
	test(t, 2, "two")
	test(t, 19, "nineteen")
	test(t, 20, "twenty")
	test(t, 55, "fifty five")
	test(t, 100, "one hundred")
	test(t, 101, "one hundred one")
	test(t, 222, "two hundred twenty two")
	test(t, -10, "minus ten")
	test(t, 1111, "one thousand one hundred eleven")
	test(t, 11111111111, "eleven billion one hundred eleven million one hundred eleven thousand one hundred eleven")
	test(t, 9223372036854775807, "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
	test(t, -9223372036854775807, "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
}
