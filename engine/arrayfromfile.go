package engine

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func init() {
        DeclFunc("ArrayFromFile", ArrayFromFile,
                "Imports a list of numbers from a file. "+
                "Arguments: filename.")
	DeclFunc("ArrayLen", ArrayLen, "Return length of slice")
}

func ArrayFromFile(fname string) []float64 {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Println("Failed to parse number:", err)
			continue
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}

func ArrayLen(x []float64) int {
	return len(x)
}
