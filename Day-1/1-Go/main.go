package main

import (
	"../../commons"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	var count int
	var input []int

	f := commons.OpenFile("input.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	for _, i := range input {
		count += (i / 3) - 2
	}

	fmt.Println(count)
}
