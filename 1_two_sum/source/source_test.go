package source

import (
	"reflect"
	"testing"
)

func Test_Source_Empty(t *testing.T) {

	nums := []int{1, 4, 5, 6}
	target := 2
	expected := []int{}

	result := TwoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected:", expected, "Actual:", result)
	}

}

func Test_Source_1(t *testing.T) {

	nums := []int{1, 4, 5, 6}
	target := 9
	expected := []int{1, 2}

	result := TwoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected:", expected, "Actual:", result)
	}

}
