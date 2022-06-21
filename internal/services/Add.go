package services

import (
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	numberArray := strings.Split(numbers, ",")

	result := 0
	for i := 0; i < len(numberArray); i++ {

		i, _ := strconv.Atoi(numberArray[i])
		result += i
	}
	return result, nil
}
