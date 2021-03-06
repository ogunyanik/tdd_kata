package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	case1, _ := Add("1,2")
	case2, _ := Add("1\n2,3")
	case3, err_case3 := Add("1,\n")
	case4, _ := Add("//;\n1;2")
	_, err5 := Add("-1,2")
	_, err5_v2 := Add("-1,-2")
	case6, _ := Add("1001,2")
	case7, _ := Add("")

	assert.Equal(t, 3, case1)
	assert.Equal(t, 6, case2)
	assert.Equal(t, 0, case3)
	assert.Equal(t, "not valid", err_case3.Error())
	assert.Equal(t, 3, case4)
	assert.Equal(t, "negatives not allowed", err5.Error())
	assert.Equal(t, "negatives not allowed!! negative numbers: (-1,-2)", err5_v2.Error())
	assert.Equal(t, 2, case6)
	assert.Equal(t, 0, case7)

}
