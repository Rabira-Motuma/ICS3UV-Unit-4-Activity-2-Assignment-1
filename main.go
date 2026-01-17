/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2026-01-5
 * Fileoverview: This program prints a pattern of numbers in a right-angled triangle.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	// variables
	var start string
	var end string
	var startNumber int
	var endNumber int
	var inclusiveSum int
	var exclusiveSum int
	var numberString string
	var number int

	reader := bufio.NewReader(os.Stdin)

	// prompts
	fmt.Printf("Enter an integer for the start of the range: ")
	start, _ = reader.ReadString('\n')
	start = strings.TrimSpace(start)
	startNumber, _ = strconv.Atoi(start)

	fmt.Printf("Enter an integer for the end of the range: ")
	end, _ = reader.ReadString('\n')
	end = strings.TrimSpace(end)
	endNumber, _ = strconv.Atoi(end)

	// for statement
	number = 1
	for number != 0{
	fmt.Printf("Enter an integer or zero (0) to end: ") 
	numberString, _ = reader.ReadString('\n')
	numberString = strings.TrimSpace(numberString)
	number, _ = strconv.Atoi(numberString)
			if number >= startNumber && number <= endNumber {
					inclusiveSum += number
				} else if number != 0 {
					exclusiveSum += number
				}
		
	}
	
	fmt.Printf("The sum of the integers inside the range is %d\n", inclusiveSum)
	fmt.Printf("The sum of the integers outside the range is %d\n", exclusiveSum)

	fmt.Println("\nDone.")
}
