package main

import (
	"bufio"
	"fmt"
	"mach-eight-coding-challenge/services"
	"os"
)

func main() {
	service := services.NewCodingChallengeService()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(service.GetPairNumbers(scanner.Text()))
	}
}
