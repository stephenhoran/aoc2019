package main

import (
	"../../commons"
	"bufio"
	"fmt"
	"strconv"
)

var input []int
var count int

func main() {

	f := commons.OpenFile("input.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	for _, m := range input {
		calcFuel(m)
	}

	fmt.Println(count)
}

func calcFuel(mass int) {
	num := (mass / 3) - 2
	if num <= 0 {
		return
	}
	count += num
	calcFuel(num)
}
