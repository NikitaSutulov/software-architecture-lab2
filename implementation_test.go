package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfix(t *testing.T) {
	res, err := PrefixToInfix("- 7 + 6 / 5 * 4 - 3 + 1 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(7 - (6 + (5 / (4 * (3 - (1 + 2))))))", res)
	}
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 * 3 2")
	fmt.Println(res)

	// Output:
	// (2 + (3 * 2))
}
