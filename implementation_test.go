package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfixEasyTwoOperands(t *testing.T) {
	res, err := PrefixToInfix("+ 254 4")
	if assert.Nil(t, err) {
		assert.Equal(t, "254 + 4", res)
	}
}

func TestPrefixToInfixEasyThreeOperands(t *testing.T) {
	res, err := PrefixToInfix("+ 254 * 4 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "(254 + (4 * 3))", res)
	}
}

func TestPrefixToInfixHardSevenOperands(t *testing.T) {
	res, err := PrefixToInfix("- 7 + 6 / 5 * 4 - 3 + 1 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(7 - (6 + (5 / (4 * (3 - (1 + 2))))))", res)
	}
}

func TestPrefixToInfixHardEightOperands(t *testing.T) {
	res, err := PrefixToInfix("+ 5 * 3 - 4 * 2 + 2 - 3 + 4 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "(5 + (3 * (4 - (2 * (2 + (3 - (4 + 2)))))))", res)
	}
}

func TestPrefixToInfixNineOperands(t *testing.T) {
	res, err := PrefixToInfix("* 5 + 44 + 22 - 4 + 8 * 4 + 2 ^ 2 5")
	if assert.Nil(t, err) {
		assert.Equal(t, "(5 * (44 + (22 + (4 - (8 + (4 * (2 + (2 ^ 5))))))))", res)
	}
}

func TestPrefixToInfixEmptyInput(t *testing.T) {
	_, err := PrefixToInfix("")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}

func TestPrefixToInfixForbiddenCharacter(t *testing.T) {
	_, err := PrefixToInfix("$")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 * 3 2")
	fmt.Println(res)

	// Output:
	// (2 + (3 * 2))
}
