package main

import (
	"../../commons"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var currentX, currentY int
var plot map[string]int
var steps []int
var lineCount int

func main() {
	f := commons.OpenFile("input.txt")

	var input [][]string

	plot = make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ","))
	}

	t := time.Now()

	for index, line := range input {
		currentX = 0
		currentY = 0
		lineCount = 0
		for _, move := range line {

			action := string(move[0])
			iterations, _ := strconv.Atoi(string(move[1:]))

			switch action {
			case "R":
				for i := 0; i < iterations; i++ {
					currentX++
					if index == 0 {
						updatePlot(currentX, currentY)
					} else {
						checkPlot(currentX, currentY)
					}
				}
			case "L":
				for i := 0; i < iterations; i++ {
					currentX--
					if index == 0 {
						updatePlot(currentX, currentY)
					} else {
						checkPlot(currentX, currentY)
					}
				}
			case "U":
				for i := 0; i < iterations; i++ {
					currentY++
					if index == 0 {
						updatePlot(currentX, currentY)
					} else {
						checkPlot(currentX, currentY)
					}
				}
			case "D":
				for i := 0; i < iterations; i++ {
					currentY--
					if index == 0 {
						updatePlot(currentX, currentY)
					} else {
						checkPlot(currentX, currentY)
					}
				}
			}
		}
	}

	score := 0
	for _, i := range steps {
		if score == 0 {
			score = i
		}

		if i < score {
			score = i
		}
	}

	fmt.Println(score)
	fmt.Println(time.Since(t))
}

func updatePlot(x int, y int) {
	lineCount++
	plot[strconv.Itoa(x)+":"+strconv.Itoa(y)] = lineCount
}

func checkPlot(x int, y int) {
	lineCount++
	if _, ok := plot[strconv.Itoa(x)+":"+strconv.Itoa(y)]; ok {
		steps = append(steps, plot[strconv.Itoa(x)+":"+strconv.Itoa(y)]+lineCount)
	}
}
