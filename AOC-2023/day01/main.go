package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func puzzle1() int {
	file, err := os.Open("../data/day01.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close() //Closing file after reading it

	//Create scanner to read file line by line
	scanner := bufio.NewScanner(file)

	var last int
	var first int
	number := 0

	//Iterate over every line of the file
	for scanner.Scan() {
		line := scanner.Text()

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = int(line[i]) - int('0')
				break
			}
		}

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = int(line[i]) - int('0')
				break
			}
		}
		number = number + first*10 + last
	}
	return number
}

func puzzle2() int {
	file, err := os.Open("../data/day01.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close() //Closing file after reading it

	//Create scanner to read file line by line
	scanner := bufio.NewScanner(file)

	number := 0

	//Iterate over every line of the file
	for scanner.Scan() {
		line := scanner.Text()
		regex := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")

		var current []string

		for i := range line {
			found := regex.FindAllString(line[i:], -1)
			// The three dots are used to unpacked elements of a slice
			current = append(current, found...)
		}

		//first, ok := digits[current[0]]
		//if ok
		if first, ok := digits[current[0]]; ok {
			number += first * 10
		}

		if last, ok := digits[current[len(current)-1]]; ok {
			number += last
		}
	}
	return number
}

func main() {
	fmt.Println("First result is: ", puzzle1())
	fmt.Println("Second result is: ", puzzle2())
}
