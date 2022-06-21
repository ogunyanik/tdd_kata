package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	case1, _ := Add("1,2")
	case2, _ := Add("1\n2,3")
	case3, err_case3 := Add("1,\n")

	assert.Equal(t, 3, case1)
	assert.Equal(t, 6, case2)
	assert.Equal(t, 0, case3)
	assert.Equal(t, "Not Valid", err_case3)

}
