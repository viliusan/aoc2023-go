// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var testInput string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	testInput = strings.TrimRight(testInput, "\n")
	if len(testInput) == 0 {
		panic("empty test.txt file")
	}
}

func main() {
	var part int
	var test bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&test, "test", false, "run with test.txt inputs?")
	flag.Parse()
	fmt.Println("Running part", part, ", test inputs = ", test)

	if test {
		input = testInput
	}

	var ans int
	switch part {
	case 1:
		ans = part1(input)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	var ans int

	numRegex := regexp.MustCompile("[0-9]+")
	splitInput := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	for i := 0; i < len(splitInput); i++ {
		cardNumbersStartIndex := strings.Index(splitInput[i], ":")
		cardNumbersEndIndex := strings.Index(splitInput[i], "|")
		cardNumbers := numRegex.FindAllString(splitInput[i][cardNumbersStartIndex+1:cardNumbersEndIndex], -1)
		ourNumbers := numRegex.FindAllString(splitInput[i][cardNumbersEndIndex+1:], -1)

		var cardSum int
		for i := 0; i < len(cardNumbers); i++ {
			for j := 0; j < len(ourNumbers); j++ {
				if cardNumbers[i] == ourNumbers[j] {
					if cardSum == 0 {
						cardSum += 1
					} else {
						cardSum *= 2
					}
				}
			}
		}

		ans += cardSum
	}

	return ans
}

func part2(input string) int {
	var ans int

	numRegex := regexp.MustCompile("[0-9]+")
	splitInput := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	scratchCardCount := make([]int, len(splitInput))
	for i := range scratchCardCount {
		scratchCardCount[i] = 1
	}

	for i := 0; i < len(splitInput); i++ {
		cardNumbersStartIndex := strings.Index(splitInput[i], ":")
		cardNumbersEndIndex := strings.Index(splitInput[i], "|")
		cardNumbers := numRegex.FindAllString(splitInput[i][cardNumbersStartIndex+1:cardNumbersEndIndex], -1)
		ourNumbers := numRegex.FindAllString(splitInput[i][cardNumbersEndIndex+1:], -1)

		var numberMatch int
		for i := 0; i < len(cardNumbers); i++ {
			for j := 0; j < len(ourNumbers); j++ {
				if cardNumbers[i] == ourNumbers[j] {
					numberMatch++
				}
			}
		}

		for n := 0; n < scratchCardCount[i]; n++ {
			for k := i + 1; k < numberMatch+i+1; k++ {
				scratchCardCount[k]++
			}
		}
	}

	for i := range scratchCardCount {
		ans += scratchCardCount[i]
	}

	return ans
}

func parseInput(input string) (parsedInput []int) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, stringToInt(line))
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}
