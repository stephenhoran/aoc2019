package main

import (
	"../../commons"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var t time.Time

func main() {
	var original []int

	index := 0
	f := commons.OpenFile("input.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := strings.Split(scanner.Text(), ",")
		for _, n := range i {
			num, _ := strconv.Atoi(n)
			original = append(original, num)
		}
	}

	t = time.Now()

	for v := 0; v <= 99; v++ {
		for n := 0; n <= 99; n++ {
			input := make([]int, len(original))
			copy(input, original)
			intCode(index, input, n, v)
		}
	}
}

func intCode(index int, input []int, noun int, verb int) {
	input[1] = noun
	input[2] = verb

	opCode := input[index]
	pos1 := input[index+1]
	pos2 := input[index+2]
	store := input[index+3]

	if opCode == 99 {
		if input[0] == 19690720 {
			fmt.Printf("Noun: %d - Verb: %d\n", noun, verb)
			fmt.Println(100*noun + verb)
			fmt.Printf("Time to work: %v", time.Since(t))
			os.Exit(1)
		}
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

	intCode(index+4, input, noun, verb)
}
