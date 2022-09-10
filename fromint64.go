package words

import (
	"math"
	"strings"
)

var units [20]string = [20]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens [10]string = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var powers [9]string = [9]string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion"}

func collectParts(parts []string, n int64, d int64, s string) (int64, []string) {
	if n < d {
		return n, parts
	}

	parts = collectPartsUnderAThousand(parts, n/d)
	parts = append(parts, s)

	return n % d, parts
}

func collectPartsUnderAThousand(parts []string, n int64) []string {
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

func FromInt64(n int64) string {
	var parts []string = make([]string, 0, 32)
	if n < 0 {
		parts = append(parts, "minus")
		n = -n
	}
	for i := 6; i > 0; i-- {
		n, parts = collectParts(parts, n, int64(math.Pow10(i*3)), powers[i])
	}
	parts = collectPartsUnderAThousand(parts, n)
	return strings.Join(parts, " ")
}
