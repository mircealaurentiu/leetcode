package source

import "testing"

func Test_Source1(t *testing.T) {

	a := "1"
	b := "1"

	if AddBinary(a, b) != "10" {
		t.Error("Expected 10")
	}
}
