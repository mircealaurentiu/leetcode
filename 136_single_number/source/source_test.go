package source

import "testing"

func Test_Source1(t *testing.T) {

	nums := [...]int{2, 2, 4, 2}
	if SingleNumber(nums[:]) != 4 {
		t.Error("Expected 4")
	}
}
