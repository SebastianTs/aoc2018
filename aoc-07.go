package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("The answer to part one is:", partOne())
	fmt.Println("The answer to part two is:", partTwo())
}

func partOne() string {
	is, err := readInput("./input/input-7.txt")
	if err != nil {
		panic(err)
	}
	adjList := make(map[rune][]rune)
	parCount := make(map[rune]int)

	for _, entry := range is {
		key := entry[0]
		val := entry[1]
		adjList[key] = append(adjList[key], val)
		parCount[val]++
	}
	done := make([]rune, 0)
	for k := range adjList {
		if parCount[k] == 0 {
			done = append(done, k)
		}
	}
	result := make([]rune, 0)
	for len(done) > 0 {
		next := make([]rune, len(done))
		copy(next, done)
		sort.Slice(next, func(i, j int) bool {
			return next[i] < next[j]
		})
		cur := next[0]
		for i := 0; i < len(done); i++ {
			if done[i] == cur {
				done = append(done[:i], done[i+1:]...)
			}
		}
		result = append(result, cur)
		for _, node := range adjList[cur] {
			parCount[node]--
			if parCount[node] == 0 {
				done = append(done, node)
			}
		}
	}
	return string(result)
}

func partTwo() string {

	const worker = 5

	is, err := readInput("./input/input-7.txt")
	if err != nil {
		panic(err)
	}
	steps := map[rune]map[rune]struct{}{}
	for _, entry := range is {
		if _, ok := steps[entry[1]]; !ok {
			steps[entry[1]] = map[rune]struct{}{}
		}
		if _, ok := steps[entry[0]]; !ok {
			steps[entry[0]] = map[rune]struct{}{}
		}
		steps[entry[1]][entry[0]] = struct{}{}
	}

	workers := map[rune]int{}
	t := 0
	for len(steps) > 0 {
		avail := make([]rune, 0)
		for todo, succ := range steps {
			if len(succ) == 0 {
				avail = append(avail, todo)
			}
		}
		sort.Slice(avail, func(i, j int) bool {
			return avail[i] < avail[j]
		})

		for _, todo := range avail {
			if _, ok := workers[todo]; !ok && len(workers) < worker {
				workers[todo] = t + int(todo) - int('A') + 61
			}
		}

		mint := make([]int, 0)
		for _, doneTime := range workers {
			mint = append(mint, doneTime)
		}
		sort.Ints(mint)
		t = mint[0]

		for task, doneTime := range workers {
			if doneTime == t {
				delete(workers, task)
				delete(steps, task)
				for todo, succs := range steps {
					for succ := range succs {
						if succ == task {
							delete(steps[todo], succ)
						}
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", t)
}

func readInput(path string) (is [][2]rune, err error) {
	is = make([][2]rune, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return is, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		if err != nil {
			return is, err
		}
		a, b := []rune(s[1]), []rune(s[7])
		is = append(is, [2]rune{a[0], b[0]})
	}
	return is, nil
}
