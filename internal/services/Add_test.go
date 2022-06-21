package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	result, err := Add("1\n2,3")
	if err != nil {
		t.Fail()
	}

	fmt.Println(result)
	assert.Equal(t, result, 6)

}
