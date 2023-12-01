// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(strings.NewReader(input))
	regex := regexp.MustCompile("[0-9]")
	sum := 0

	for scanner.Scan() {
        line := scanner.Text()
		parsedDigits := regex.FindAllString(line, -1)
		firstDigit := stringToInt(parsedDigits[0])
		lastDigit := stringToInt(parsedDigits[len(parsedDigits) -1])
		sum += firstDigit * 10 + lastDigit
    }

	return sum
}

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	regex := regexp.MustCompile(`[0-9]|eightwo|oneight|twone|one|two|three|four|five|six|seven|eight|eight|nine`)
	sum := 0

	for scanner.Scan() {
        line := scanner.Text()
		readLine := regex.FindAllString(line, -1)
		stringDigits := parseLine(readLine)
		parsedDigits := parseDigits(stringDigits)
		firstDigit := parsedDigits[0]
		lastDigit := parsedDigits[len(parsedDigits)-1]
		sum += firstDigit * 10 + lastDigit
    }

	return sum
}

func parseDigits(parsedDigits []string) []int {
	var newArray []int

	for i := 0; i <= len(parsedDigits)-1; i++ {
		var firstDigit int
		var secondDigit int

		parsedInt := stringToInt(parsedDigits[i])
		if parsedInt > 9 {
			firstDigit = parsedInt / 10
			secondDigit = parsedInt % 10

			newArray = append(newArray, firstDigit, secondDigit)
		} else {
			newArray = append(newArray, parsedInt)
		}
	}

	return newArray
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}

func parseLine(input []string) []string {
	for i := 0; i <= len(input)-1; i++ {
		_, error := strconv.Atoi(input[i])
		if error != nil {
			input[i] = wordToDigit(input[i])
			}
		}
	return input
}

func wordToDigit(input string) string  {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "eightwo":
		return "82"
	case "oneight":
		return "18"
	case "twone":
		return "21"
	default:
		panic("An error occured while parsing")
	}
}