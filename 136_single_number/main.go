package main

import (
	"fmt"
	"singlenumber/source"
)

func main() {

	nums := [...]int{2, 2, 3, 2}
	fmt.Println(source.SingleNumber(nums[:]))
}
