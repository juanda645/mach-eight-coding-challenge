package services

import (
	"fmt"
	"strconv"
	"strings"
)

type CodingChallengeService struct{}

func NewCodingChallengeService() *CodingChallengeService {
	return &CodingChallengeService{}

}

func (service *CodingChallengeService) GetPairNumbers(inputData string) string {
	input := strings.Split(inputData, " ")
	numbers := strings.Split(input[0], ",")
	expectedSum, _ := strconv.Atoi(input[1])
	response := ""
	operationResults := ""

	for _, value := range numbers {
		currentNumber, _ := strconv.Atoi(value)
		result := expectedSum - currentNumber
		operationResults += fmt.Sprintf(":%d:", result)

		if strings.Contains(operationResults, fmt.Sprintf(":%s:", value)) {
			response += fmt.Sprintf("%d,%d\n", currentNumber, result)
		}
	}

	return response
}

func (service *CodingChallengeService) GetPairNumbersAlternative(inputData string) string {
	input := strings.Split(inputData, " ")
	numbers := strings.Split(input[0], ",")
	expectedSum, _ := strconv.Atoi(input[1])
	response := ""

	for i, value := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			currentNumber, _ := strconv.Atoi(value)
			compareNumber, _ := strconv.Atoi(numbers[j])

			if currentNumber+compareNumber == expectedSum {
				response += fmt.Sprintf("%d,%d\n", currentNumber, compareNumber)
			}
		}
	}

	return response
}
