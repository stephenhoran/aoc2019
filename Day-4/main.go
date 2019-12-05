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

		if firstDigit > secondDigit || secondDigit > thirdDigit || thirdDigit > forthDigit || forthDigit > fifthDigit || fifthDigit > lastDigit {
			continue
		}

		if firstDigit == secondDigit || secondDigit == thirdDigit || thirdDigit == forthDigit || forthDigit == fifthDigit || fifthDigit == lastDigit {
			double = true
		}

		if double {
			count++
		}

	}

	fmt.Println(count)
	fmt.Println(time.Since(t))
}
