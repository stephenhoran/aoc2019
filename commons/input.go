package commons

import "os"

func OpenFile(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	return f
}
