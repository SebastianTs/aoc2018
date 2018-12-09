package main

import (
	"bufio"
	"container/list"
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
	fmt.Println("if the number of the last marble were 100 times larger the winning Elf's score is", playList(players, points*100))
}

func play(players, last_marble int) int {
	scores := make([]int,players)
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

func playList(players, totalPlays int) int {
	marbles := list.New()
	currentMarble := marbles.PushBack(0)
	score := make([]int, players)

	for nextMarble := 1; nextMarble < totalPlays; nextMarble++ {
		if nextMarble%23 == 0 {
			rm := currentMarble
			for i := 0; i < 7; i++ {
				rm = rm.Prev()
				if rm == nil {
					rm = marbles.Back()
				}
			}
			score[nextMarble % players] += nextMarble + rm.Value.(int)
			currentMarble = rm.Next()
			marbles.Remove(rm)
		} else {
			n := currentMarble.Next()
			if n == nil {
				n = marbles.Front()
			}
			currentMarble = marbles.InsertAfter(nextMarble, n)
		}
	}
	max := 0
	for _, v := range score {
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
