/*
 Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

Submit your source code for the program, “slice.go”.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func scanInt(num *int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temp := scanner.Text()
	if temp == "X" {
		os.Exit(0)
	}
	n, err := strconv.Atoi(temp)
	if err == nil {
		*num = n
	} else {
		fmt.Println("Please Enter A Valid Number!!")
		scanInt(num)
	}

}
func main() {
	sortedSlice := make([]int, 3, 3)
	fmt.Println("Enter numbers to add into sorted list & Enter 'X' to Exit")
	num := 0
	for true {
		fmt.Println("\nEnter a Number : ")
		scanInt(&num)
		sortedSlice = append(sortedSlice, num)
		sort.Ints(sortedSlice)
		fmt.Println("\nNew Sorted Slice\n", sortedSlice)
	}
}
