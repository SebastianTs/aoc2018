package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ids, err := readInputIDs("./input/input-2.txt")
	if err != nil {
		panic(err)
	}
	freq := make(map[int]int)
	for _, id := range ids {
		count := make(map[rune]int)
		for _, c := range id {
			count[c] += 1
		}
		seen := make(map[int]bool)
		for _, v := range count {
			if v > 1 && !seen[v] {
				freq[v] += 1
				seen[v] = true
			}
		}
	}
	res := 1
	for _, v := range freq {
		res *= v
	}
	fmt.Println("checksum is: ", res)

	//part2
	solution := make([]byte, len(ids[0])-1)
	idx := 0
	for _, x := range ids {
		for _, y := range ids {
			diff := 0
			for i := range x {
				if x[i] != y[i] {
					diff += 1
				}
			}
			if diff == 1 {
				for i := range x {
					if x[i] == y[i] {
						if idx < len(solution) {
							solution[idx] = x[i]
							idx++
						} else {
							break
						}
					}
				}
			}
		}
	}
	fmt.Print(string(solution))

}

func readInputIDs(filename string) (ids []string, err error) {
	ids = make([]string, 0)
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return ids, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ids = append(ids, line)
	}
	return ids, nil
}
