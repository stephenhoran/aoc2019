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
			input[1] = n
			input[2] = v
			intCode(index, input)

			if input[0] == 19690720 {
				fmt.Printf("Noun: %d - Verb: %d\n", n, v)
				fmt.Println(100*n + v)
				fmt.Println(time.Since(t))
				os.Exit(0)
			}
		}
	}
}

func intCode(index int, input []int) {
	opCode := input[index]
	pos1 := input[index+1]
	pos2 := input[index+2]
	store := input[index+3]

	if opCode == 99 {
		return
	}

	if opCode == 1 {
		input[store] = input[pos1] + input[pos2]
	}

	if opCode == 2 {
		input[store] = input[pos1] * input[pos2]
	}

	intCode(index+4, input)
}
