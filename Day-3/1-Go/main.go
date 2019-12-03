package main

import (
	"../../commons"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var currentX, currentY int
var plot map[string]struct{}
var collisions [][]int

func main() {
	f := commons.OpenFile("input.txt")

	var input [][]string

	plot = make(map[string]struct{})

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ","))
	}

	for index, line := range input {
		currentX = 0
		currentY = 0
		for _, move := range line {
			action := string(move[0])
			iterations, _ := strconv.Atoi(string(move[1:]))

			switch action {
			case "R":
				for i := 0; i < iterations; i++ {
					currentX++
					if index == 0 {
						plot[strconv.Itoa(currentX)+":"+strconv.Itoa(currentY)] = struct{}{}
					} else {
						updatePlot(currentX, currentY)
					}
				}
			case "L":
				for i := 0; i < iterations; i++ {
					currentX--
					if index == 0 {
						plot[strconv.Itoa(currentX)+":"+strconv.Itoa(currentY)] = struct{}{}
					} else {
						updatePlot(currentX, currentY)
					}
				}
			case "U":
				for i := 0; i < iterations; i++ {
					currentY++
					if index == 0 {
						plot[strconv.Itoa(currentX)+":"+strconv.Itoa(currentY)] = struct{}{}
					} else {
						updatePlot(currentX, currentY)
					}
				}
			case "D":
				for i := 0; i < iterations; i++ {
					currentY--
					if index == 0 {
						plot[strconv.Itoa(currentX)+":"+strconv.Itoa(currentY)] = struct{}{}
					} else {
						updatePlot(currentX, currentY)
					}
				}
			}
		}
	}

	point := 0

	for _, collision := range collisions {
		p := int(math.Abs(float64(collision[0]))) + int(math.Abs(float64(collision[1])))
		if point == 0 {
			point = p
		} else if p < point {
			point = p
		}
	}

	fmt.Println(point)
}

func updatePlot(x int, y int) {
	if _, ok := plot[strconv.Itoa(x)+":"+strconv.Itoa(y)]; ok {
		collisions = append(collisions, []int{x, y})
	}
}
