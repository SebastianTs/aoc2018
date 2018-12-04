package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	freqs, err := readInput("./input/input-1a.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, freq := range freqs {
		sum += freq
	}
	fmt.Println("sum:", sum)

	sum = 0
	seen := make(map[int]bool)
	for i := 0; ; i = (i + 1) % len(freqs) {
		sum += freqs[i]
		if seen[sum] {
			fmt.Print("first:", sum)
			break
		}
		seen[sum] = true
	}
}

func readInput(filename string) (ns []int, err error) {
	ns = make([]int, 0)
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return ns, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		freq, err := strconv.Atoi(line)
		if err != nil {
			return ns, err
		}
		ns = append(ns, freq)
	}
	return ns, nil
}
