package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodingChallengeService_GetPairNumbers(t *testing.T) {
	inputData := "1,9,5,0,20,-4,12,16,7 12"
	t.Run("should get pair numbers when it is call", func(t *testing.T) {
		expectedResponse := "12,0\n16,-4\n7,5\n"
		service := NewCodingChallengeService()
		response := service.GetPairNumbers(inputData)

		assert.Equal(t, expectedResponse, response)
	})
}

func TestCodingChallengeService_GetPairNumbersAlternative(t *testing.T) {
	inputData := "1,9,5,0,20,-4,12,16,7,11,3 12"
	t.Run("should get pair numbers when it is call", func(t *testing.T) {
		expectedResponse := "1,11\n9,3\n5,7\n0,12\n-4,16\n"
		service := NewCodingChallengeService()
		response := service.GetPairNumbersAlternative(inputData)

		assert.Equal(t, expectedResponse, response)
	})
}
