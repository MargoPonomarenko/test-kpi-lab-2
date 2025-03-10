package kpi_lab_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "( 4 - 2 ) * 3 + 5 ", res)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
