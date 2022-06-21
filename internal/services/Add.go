package services

import (
	"errors"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {

	numbers = strings.Replace(numbers, "\n", ",", -1)

	numberArray := strings.Split(numbers, ",")

	result := 0
	for i := 0; i < len(numberArray); i++ {

		i, err := strconv.Atoi(numberArray[i])
		if err != nil {
			return 0, errors.New("not valid")

		}
		result += i
	}
	return result, nil
}
