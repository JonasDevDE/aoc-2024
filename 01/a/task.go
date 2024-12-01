package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInputLines() []string {
	workingDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to get working directory", err)
	}

	input, err := os.ReadFile(workingDirectory + "/01/" + "input.txt")

	if err != nil {
		log.Fatalln("Error while reading input.txt", err)
	}

	return strings.Split(strings.ReplaceAll(string(input), "\r", ""), "\n")
}

func main() {
	input := readInputLines()

	var columnA []int
	var columnB []int

	for i := range input {
		split := strings.Split(input[i], "   ")

		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])

		columnA = append(columnA[:i], first)
		columnB = append(columnB[:i], second)
	}

	sort.Ints(columnA)
	sort.Ints(columnB)

	sum := 0

	for i := range input {
		columnAIndex := float64(columnA[i])
		columnBIndex := float64(columnB[i])

		sum += int(math.Max(columnAIndex, columnBIndex) - math.Min(columnAIndex, columnBIndex))
	}

	fmt.Printf("Calculated sum: %d\n", sum)
}
