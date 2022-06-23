package services

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {

	//parse string via regex
	regexPattern := regexp.MustCompile("^//(.)\n(.*)$")

	macthedArray := regexPattern.FindStringSubmatch(numbers)

	var delimiter string
	if len(macthedArray) > 0 {
		delimiter = macthedArray[1]
		numbers = macthedArray[2]
	}

	numbers = strings.Replace(numbers, "\n", ",", -1)
	var numberArray []string
	if delimiter != "" {
		numberArray = strings.Split(numbers, delimiter)

	} else {

		numberArray = strings.Split(numbers, ",")
	}

	result := 0
	negativeNumbers := []string{}
	for i := 0; i < len(numberArray); i++ {

		number, err := strconv.Atoi(numberArray[i])
		if err != nil {
			return 0, errors.New("not valid")

		}

		if i > 0 {
			result += number
		}

		negativeNumbers = append(negativeNumbers, numberArray[i])

	}

	negativeNumberCount := len(negativeNumbers)
	if negativeNumberCount == 1 {
		return 0, errors.New("negatives not allowed")
	}
	if negativeNumberCount > 1 {
		return 0, errors.New("negatives not allowed!!" + " negative numbers: (" + strings.Join(negativeNumbers, ",") + ")")
	}

	return result, nil
}
