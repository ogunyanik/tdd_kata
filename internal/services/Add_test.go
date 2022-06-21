package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	result, err := Add("1,2")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, result, 3)

}
