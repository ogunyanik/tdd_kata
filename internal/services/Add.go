package services

import (
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {

	numbers = strings.Replace(numbers, "\n", ",", -1)

	numberArray := strings.Split(numbers, ",")

	result := 0
	for i := 0; i < len(numberArray); i++ {

		i, _ := strconv.Atoi(numberArray[i])
		result += i
	}
	return result, nil
}
