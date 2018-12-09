package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

func main() {
	players, points, err := readInput("./input/input-9.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Println(play(10,1618))
	fmt.Println("the winning Elf's score is", play(players, points))
	fmt.Println("if the number of the last marble were 100 times larger the winning Elf's score is", play(players, points*100))
}

func play(players, last_marble int) int {
	scores := make(map[int]int)
	circle := ring.New(1)
	circle.Value = 0

	for m := 1; m < last_marble+1; m++ {
		if m%23 == 0 {
			circle = circle.Move(-6)
			s := circle.Move(-2).Link(circle)
			scores[m%players] += m + s.Value.(int)
		} else {
			circle = circle.Next()
			circle.Link(func(i int) *ring.Ring {
				r := ring.New(1)
				r.Value = i
				return r
			}(m))
			circle = circle.Next()
		}
	}
	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max
}

func readInput(path string) (players, points int, err error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return 0, 0, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &players, &points)
		if err != nil {
			return 0, 0, err
		}
	}
	return players, points, nil
}
