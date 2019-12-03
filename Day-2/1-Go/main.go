package main

import (
	"../../commons"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var input []int

func main() {
	index := 0
	f := commons.OpenFile("input.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := strings.Split(scanner.Text(), ",")
		for _, n := range i {
			num, _ := strconv.Atoi(n)
			input = append(input, num)
		}
	}

	t := time.Now()

	intCode(index)

	fmt.Printf("Time to work: %v", time.Since(t))
}

func intCode(index int) {
	opCode := input[index]
	pos1 := input[index+1]
	pos2 := input[index+2]
	store := input[index+3]

	if opCode == 99 {
		fmt.Println(input[0])
		return
	}

	if opCode == 1 {
		value := input[pos1] + input[pos2]
		input[store] = value
	}

	if opCode == 2 {
		value := input[pos1] * input[pos2]
		input[store] = value
	}

	intCode(index + 4)
}
