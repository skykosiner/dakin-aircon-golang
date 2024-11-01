package utils

import "testing"

func TestPowerToBool(t *testing.T) {
	testCases := []struct {
		power    string
		shouldBe bool
	}{
		{"On", true},
		{"Off", false},
	}

	for _, test := range testCases {
		if powerToBool := PowerToBool(test.power); powerToBool != test.shouldBe {
			t.Logf("Powre state for bool did not work, it's joever. %+v %t", test, powerToBool)
			t.Fail()
		}
	}
}
