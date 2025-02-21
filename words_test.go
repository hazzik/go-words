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
	test(t, -128, "minus one hundred twenty eight")
}

func TestFromUint8(t *testing.T) {
	test := func(t *testing.T, input uint8, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			t.Helper()
			assert.Equal(t, FromUint8(input), expected)
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
	test(t, 255, "two hundred fifty five")
}

func TestFromInt16(t *testing.T) {
	test := func(t *testing.T, input int16, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			t.Helper()
			assert.Equal(t, FromInt16(input), expected)
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
	test(t, 32767, "thirty two thousand seven hundred sixty seven")
	test(t, -32767, "minus thirty two thousand seven hundred sixty seven")
	test(t, -32768, "minus thirty two thousand seven hundred sixty eight")
}

func TestFromUint16(t *testing.T) {
	test := func(t *testing.T, input uint16, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			t.Helper()
			assert.Equal(t, FromUint16(input), expected)
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
	test(t, 65535, "sixty five thousand five hundred thirty five")
}

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
	test(t, -2147483648, "minus two billion one hundred forty seven million four hundred eighty three thousand six hundred forty eight")
}

func TestFromUint32(t *testing.T) {
	test := func(t *testing.T, input uint32, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			assert.Equal(t, FromUint32(input), expected)
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
	test(t, 4294967295, "four billion two hundred ninety four million nine hundred sixty seven thousand two hundred ninety five")
}

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
	test(t, -9223372036854775808, "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred eight")
}

func TestFromUint64(t *testing.T) {
	test := func(t *testing.T, input uint64, expected string) {
		t.Helper()
		t.Run(fmt.Sprint(input), func(t *testing.T) {
			t.Helper()
			assert.Equal(t, FromUint64(input), expected)
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
	test(t, 1111, "one thousand one hundred eleven")
	test(t, 11111111111, "eleven billion one hundred eleven million one hundred eleven thousand one hundred eleven")
	test(t, 9223372036854775807, "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
	test(t, 18446744073709551615, "eighteen quintillion four hundred forty six quadrillion seven hundred forty four trillion seventy three billion seven hundred nine million five hundred fifty one thousand six hundred fifteen")
}

func TestFromInteger(t *testing.T) {
	testI(t, 0, "zero")
	testI(t, 1, "one")
	testI(t, 2, "two")
	testI(t, 19, "nineteen")
	testI(t, 20, "twenty")
	testI(t, 55, "fifty five")
	testI(t, 100, "one hundred")
	testI(t, 101, "one hundred one")
	testI(t, 222, "two hundred twenty two")
	testI(t, -10, "minus ten")
	testI(t, 1111, "one thousand one hundred eleven")
	testI(t, 11111111111, "eleven billion one hundred eleven million one hundred eleven thousand one hundred eleven")
	testI(t, 9223372036854775807, "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
	testI(t, -9223372036854775807, "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
	testI(t, -9223372036854775808, "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred eight")
	testI(t, uint64(18446744073709551615), "eighteen quintillion four hundred forty six quadrillion seven hundred forty four trillion seventy three billion seven hundred nine million five hundred fifty one thousand six hundred fifteen")
}

func testI[I interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}](t *testing.T, input I, expected string) {
	t.Helper()
	t.Run(fmt.Sprint(input), func(t *testing.T) {
		assert.Equal(t, FromInteger(input), expected)
	})
}
