package main

import (
	"fmt"
	"time"
)

const startingRange int = 234208
const endingRange int = 765869

func main() {
	t := time.Now()
	count := 0
	for i := startingRange; i < endingRange; i++ {
		double := false
		firstDigit := (i / 100000) % 10
		secondDigit := (i / 10000) % 10
		thirdDigit := (i / 1000) % 10
		forthDigit := (i / 100) % 10
		fifthDigit := (i / 10) % 10
		lastDigit := i % 10

		// Are all the of numbers in order
		if firstDigit > secondDigit || secondDigit > thirdDigit || thirdDigit > forthDigit || forthDigit > fifthDigit || fifthDigit > lastDigit {
			continue
		}

		// Does the first 2 numbers match without matching the third
		if (firstDigit == secondDigit) && (secondDigit != thirdDigit) {
			double = true
		}

		// Does 2-3 match without matching 1-4
		if (secondDigit == thirdDigit) && (secondDigit != firstDigit) && (thirdDigit != forthDigit) {
			double = true
		}

		// Does 3-4 match without matching 2 or 5
		if (thirdDigit == forthDigit) && (thirdDigit != secondDigit) && (forthDigit != fifthDigit) {
			double = true
		}

		// Does 4-5 match without matching 3-6
		if (forthDigit == fifthDigit) && (forthDigit != thirdDigit) && (fifthDigit != lastDigit) {
			double = true
		}

		// Does 5-6 match
		if (fifthDigit == lastDigit) && (fifthDigit != forthDigit) {
			double = true
		}

		if double {
			count++
		}
	}

	fmt.Println(count)
	fmt.Println(time.Since(t))
}
