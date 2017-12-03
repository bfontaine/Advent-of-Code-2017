package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseSpreadsheetLine(line string, sep string) ([]int, error) {
	var err error

	cells := strings.Split(strings.TrimRight(line, "\n"), sep)

	values := make([]int, len(cells))
	for i, cell := range cells {
		values[i], err = strconv.Atoi(cell)
		if err != nil {
			break
		}
	}
	return values, err
}

func minMaxLineChecksum(values []int) (int, error) {
	var min, max int

	for i, curr := range values {
		if i == 0 || curr < min {
			min = curr
		}

		if i == 0 || curr > max {
			max = curr
		}
	}

	return max - min, nil
}

func evenlyDivisibleLineChecksum(values []int) (int, error) {
	return 0, nil
}

type LineChecksum func([]int) (int, error)

func computeChecksum(r *bufio.Reader, lineFn LineChecksum) (int, error) {
	var line string
	var sum, partialSum int
	var values []int

	var readError, parseError, checksumError error

	for {
		line, readError = r.ReadString('\n')
		values, parseError = parseSpreadsheetLine(line, "\t")

		if parseError != nil {
			return sum, parseError
		}

		if readError != nil && readError != io.EOF {
			return sum, readError
		}

		partialSum, checksumError = lineFn(values)
		if checksumError != nil {
			return sum, checksumError
		}

		sum += partialSum

		if readError == io.EOF {
			readError = nil
			break
		}
	}
	return sum, readError
}

func main() {
	var checksumFuncId int

	flag.IntVar(&checksumFuncId, "f", 1, "Checksum function id (1 or 2)")
	flag.Parse()

	if flag.NArg() != 1 || checksumFuncId < 1 || checksumFuncId > 2 {
		log.Fatalf("Usage:\n    %s [options] <problem_filename>", os.Args[0])
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var checksumFunc LineChecksum

	if checksumFuncId == 1 {
		checksumFunc = minMaxLineChecksum
	} else {
		checksumFunc = evenlyDivisibleLineChecksum
	}

	reader := bufio.NewReader(file)
	checksum, err := computeChecksum(reader, checksumFunc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Checksum: %d\n", checksum)
}
