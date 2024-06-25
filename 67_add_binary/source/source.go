package source

import (
	"math/big"
)

func AddBinary(a string, b string) string {
	// Create big.Int values for the input binary strings
	res_a := new(big.Int)
	res_b := new(big.Int)

	// Set the values of res_a and res_b from the binary strings
	res_a.SetString(a, 2)
	res_b.SetString(b, 2)

	// Add the two big.Int values
	res_final := new(big.Int).Add(res_a, res_b)

	// Convert the result to a binary string
	ret_string := res_final.Text(2)

	return ret_string
}
