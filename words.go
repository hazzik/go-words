package words

import (
	"math"
	"strings"
)

var units [20]string = [20]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens [10]string = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var powers [9]string = [9]string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion"}

func collectParts(parts []string, n uint64, d uint64, s string) (uint64, []string) {
	if n < d {
		return n, parts
	}

	parts = collectPartsUnderAThousand(parts, n/d)
	parts = append(parts, s)

	return n % d, parts
}

func collectPartsUnderAThousand(parts []string, n uint64) []string {
	if n >= 100 {
		parts = append(parts, units[n/100])
		parts = append(parts, "hundred")
		n = n % 100
	}
	if n >= 20 {
		parts = append(parts, tens[n/10])
		n = n % 10
	}
	if n > 0 {
		parts = append(parts, units[n])
	}
	return parts
}

func FromUint8(n uint8) string {
	return FromInteger(n)
}

func FromInt8(n int8) string {
	return FromInteger(n)
}

func FromInt16(n int16) string {
	return FromInteger(n)
}

func FromUint16(n uint16) string {
	return FromInteger(n)
}

func FromInt32(n int32) string {
	return FromInteger(n)
}

func FromUint32(n uint32) string {
	return FromInteger(n)
}

func FromInt64(n int64) string {
	if n < 0 {
		return "minus " + FromUint64(uint64(-n))
	}

	return FromUint64(uint64(n))
}

func FromUint64(n uint64) string {
	if n == 0 {
		return "zero"
	}

	var parts []string = make([]string, 0, 32)
	for i := 6; i > 0; i-- {
		n, parts = collectParts(parts, n, uint64(math.Pow10(i*3)), powers[i])
	}

	parts = collectPartsUnderAThousand(parts, n)
	return strings.Join(parts, " ")
}

func FromInt(n int) string {
	return FromInteger(n)
}

func FromUint(n uint) string {
	return FromInteger(n)
}

func FromInteger[I interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}](n I) string {
	if n < 0 {
		return "minus " + FromUint64(uint64(-n))
	}

	return FromUint64(uint64(n))
}
