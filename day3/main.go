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
	specialCharRegex := regexp.MustCompile("[$&+,:;=?@#|'<>/^*()%!-]")

	splitInput := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	for i := 0; i < len(splitInput); i++ {
		parsedNums := numRegex.FindAllString(splitInput[i], -1)
		numsIndexes := numRegex.FindAllStringIndex(splitInput[i], -1)
		for pos, number := range parsedNums {
			startIndex := numsIndexes[pos][0]
			endIndex := numsIndexes[pos][1]
			
			var nextRowMatch bool
			var previousRowMatch bool
			var sidesMatch bool

			if startIndex == 0 {
				sidesMatch = specialCharRegex.Match([]byte(splitInput[i][startIndex:endIndex+1]))
			} else if endIndex == len(splitInput[i]) {
				sidesMatch = specialCharRegex.Match([]byte(splitInput[i][startIndex-1:endIndex]))
			} else {
				sidesMatch = specialCharRegex.Match([]byte(splitInput[i][startIndex-1:endIndex+1]))
			}

			if (i == 0) {
				previousRowMatch = false
			} else {
				if startIndex == 0 {
					previousRowMatch = specialCharRegex.Match([]byte(splitInput[i-1][startIndex:endIndex+1]))

				} else if endIndex == len(splitInput[i]) {
					previousRowMatch = specialCharRegex.Match([]byte(splitInput[i-1][startIndex-1:endIndex]))

				} else {
					previousRowMatch = specialCharRegex.Match([]byte(splitInput[i-1][startIndex-1:endIndex+1]))

				}
			}
			if (i == len(splitInput)-1) {
				nextRowMatch = false
			} else {
				if startIndex == 0 {
					nextRowMatch = specialCharRegex.Match([]byte(splitInput[i+1][startIndex:endIndex+1]))

				} else if endIndex == len(splitInput[i]) {
					nextRowMatch = specialCharRegex.Match([]byte(splitInput[i+1][startIndex-1:endIndex]))

				} else {
					nextRowMatch = specialCharRegex.Match([]byte(splitInput[i+1][startIndex-1:endIndex+1]))

				}
			}

			if nextRowMatch || previousRowMatch || sidesMatch {
				ans += stringToInt(number)
			}

		}
	}

	return ans
}

func part2(input string) int {
	return 0
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