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

func lineChecksum(line string, sep string) (int, error) {
	var min, max, curr int
	var err error

	cells := strings.Split(line, sep)

	for i, cell := range cells {
		curr, err = strconv.Atoi(cell)
		if err != nil {
			return 0, err
		}
		if i == 0 || curr < min {
			min = curr
		}

		if i == 0 || curr > max {
			max = curr
		}
	}

	return max - min, nil
}

func computeChecksum(r *bufio.Reader) (sum int, err error) {
	var line string
	var partialSum int
	var err2 error

	for {
		line, err = r.ReadString('\n')

		partialSum, err2 = lineChecksum(strings.TrimRight(line, "\n"), "\t")
		if err2 != nil {
			return sum, err2
		}

		sum += partialSum

		if err != nil {
			break
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func main() {
	flag.Parse() // print an error if someone passes an option

	if len(os.Args) != 2 {
		log.Fatalf("Usage:\n    %s <problem_filename>", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	checksum, err := computeChecksum(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Checksum: %d\n", checksum)
}
