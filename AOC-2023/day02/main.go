package main

import (
	"bufio"
	"fmt"
	"os"
)

func puzzle1() int {
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
		for i := range line {
			number = number + i
		}
	}
	return number
}

// func puzzle2() int {
// 	file, err := os.Open("../data/day01.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file: ", err)
// 	}
// 	defer file.Close() //Closing file after reading it
//
// 	//Create scanner to read file line by line
// 	scanner := bufio.NewScanner(file)
//
// 	//Iterate over every line of the file
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		}
// }

func main() {
	fmt.Println("First result is: ", puzzle1())
	// fmt.Println("Second result is: ", puzzle2())
}
