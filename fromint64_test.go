package words

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestZero(t *testing.T) {
	assert.Equal(t, FromInt64(0), "zero")
}

func TestOne(t *testing.T) {
	assert.Equal(t, FromInt64(1), "one")
}

func TestTwo(t *testing.T) {
	assert.Equal(t, FromInt64(2), "two")
}

func TestNineteen(t *testing.T) {
	assert.Equal(t, FromInt64(19), "nineteen")
}

func TestTwenty(t *testing.T) {
	assert.Equal(t, FromInt64(20), "twenty")
}

func TestFiftyFive(t *testing.T) {
	assert.Equal(t, FromInt64(55), "fifty five")
}

func TestOneHundred(t *testing.T) {
	assert.Equal(t, FromInt64(100), "one hundred")
}

func TestOneHundredAndOne(t *testing.T) {
	assert.Equal(t, FromInt64(101), "one hundred one")
}

func TestTwoHundredsTwentyTwo(t *testing.T) {
	assert.Equal(t, FromInt64(222), "two hundred twenty two")
}

func TestMinusTen(t *testing.T) {
	assert.Equal(t, FromInt64(-10), "minus ten")
}

func Test1111(t *testing.T) {
	assert.Equal(t, FromInt64(1111), "one thousand one hundred eleven")
}

func Test11111111111(t *testing.T) {
	assert.Equal(t, FromInt64(11111111111), "eleven billion one hundred eleven million one hundred eleven thousand one hundred eleven")
}

func TestMax(t *testing.T) {
	assert.Equal(t, FromInt64(9223372036854775807), "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
}

func TestMin(t *testing.T) {
	assert.Equal(t, FromInt64(-9223372036854775807), "minus nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven")
}
