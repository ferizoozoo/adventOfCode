package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	var inputs []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		inputs = append(inputs, input)
	}

	gamma := getGammaRate(inputs)
	epsilon := getEpsilonRate(inputs)

	fmt.Printf("gamma: %d\nepsilon: %d\n", gamma, epsilon)
	fmt.Printf("power: %d", gamma*epsilon)
}

func getGammaRate(numbers []string) int64 {
	var resultString string = ""

	for i, _ := range numbers[0] {
		numberOfOnes := 0
		numberOfZeros := 0

		for j, _ := range numbers {
			if numbers[j][i] == '0' {
				numberOfZeros++
			} else if numbers[j][i] == '1' {
				numberOfOnes++
			}
		}

		if numberOfOnes > numberOfZeros {
			resultString += "1"
		} else {
			resultString += "0"
		}
	}

	result, _ := strconv.ParseInt(resultString, 2, 32)
	return result
}

func getEpsilonRate(numbers []string) int64 {
	var resultString string = ""
	lengthOfEachNumber := utf8.RuneCountInString(numbers[0])
	lengthOfNumbers := len(numbers)

	for i := 0; i < lengthOfEachNumber; i++ {
		numberOfOnes := 0
		numberOfZeros := 0

		for j := 0; j < lengthOfNumbers; j++ {
			if numbers[j][i] == '0' {
				numberOfZeros++
			} else if numbers[j][i] == '1' {
				numberOfOnes++
			}
		}

		if numberOfOnes < numberOfZeros {
			resultString += "1"
		} else {
			resultString += "0"
		}
	}

	result, _ := strconv.ParseInt(resultString, 2, 32)
	return result
}
