package strconvuint

import (
	"testing"
)

type parseUintTest struct {
	in  string
	out uint64
}

var parseUintTests = []parseUintTest{
	{"", 0},
	{"0", 0},
	{"1", 1},
	{"12345", 12345},
	{"012345", 12345},
	{"12345x", 0},
	{"98765432100", 98765432100},
	{"18446744073709551615", 1<<64 - 1},
}

func TestAtouint(t *testing.T) {
	for i := range parseUintTests {
		test := &parseUintTests[i]
		out, _ := Atouint(test.in)
		if test.out != uint64(out) {
			t.Errorf("Atouint(%q) = %v want %v",
				test.in, out, test.out)
		}
	}
}
