package main

import (
	"bufio"
	"fmt"
	"os"
)

const(
 	minutes = 10
 	minutesPart2 = 1000000000
 	open rune = '.'
 	tree rune = '|'
 	lumberyard rune = '#'
 	)


func main() {
	area, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}
	score := make(map[int]int)
	for i:=1; i <= minutesPart2; i++ {
		next := apply(area)
		area = next
		if i == minutes{
			fmt.Println("the total resource value of the lumber collection acres after 10 minutes is", resourceScore(area))
		}
		cur := resourceScore(area)
		if _, ok := score[cur]; !ok {
			score[cur] = i
		} else {
			period := i - score[cur]
			if i % period == minutesPart2 % period{
				break
			}
		}
	}
	fmt.Println("the total resource value of the lumber collection acres after all is", resourceScore(area))
}

func print(s area){
	for y := 0; y <= s.maxDimensions.y; y++ {
		for x := 0; x <= s.maxDimensions.x; x++ {
			fmt.Print(string(s.acres[xy{x, y}]))
		}
		fmt.Println()
	}
}

func resourceScore(s area) int{
	if s.resourceCount != nil{
		return s.resourceCount[lumberyard] * s.resourceCount[tree]
	}
	count := make(map[rune]int)
	for y := 0; y <= s.maxDimensions.y; y++ {
		for x := 0; x <= s.maxDimensions.x; x++ {
			cur := s.acres[xy{x, y}]
			count[cur]++
		}
	}
	return count[lumberyard] * count[tree]
}

func apply(s area) area {
	out := area{ acres:  make(map[xy]rune),
				maxDimensions: s.maxDimensions,
				resourceCount: make(map[rune]int),
    }
	for x := 0; x <= s.maxDimensions.x; x++ {
		for y := 0; y <= s.maxDimensions.y; y++ {
			adjCount := make(map[rune]int)
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if !(i == 0 && j == 0) {
						cur := s.acres[xy{x + i, y + j}]
						adjCount[cur]++
					}
				}
			}
			pos := xy{x, y}
			switch s.acres[pos] {
			case open:
				if adjCount[tree] >= 3 {
					out.acres[pos] = tree
					out.resourceCount[tree]++
				} else {
					out.acres[pos] = open
				}
			case tree:
				if adjCount[lumberyard] >= 3 {
					out.acres[pos] = lumberyard
					out.resourceCount[lumberyard]++
				} else {
					out.acres[pos] = tree
					out.resourceCount[tree]++
				}
			case lumberyard:
				if adjCount[lumberyard] >= 1 && adjCount[tree] >= 1 {
					out.acres[pos] = lumberyard
					out.resourceCount[lumberyard]++
				} else {
					out.acres[pos] = open
				}
			}
		}
	}
	return out
}

type xy struct{
	x,y int
}

type area struct{
	acres         map[xy]rune
	maxDimensions xy
	resourceCount map[rune]int
}

func readInput(path string) (out area, err error) {
	landscape := make(map[xy]rune)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return area{} ,  err
	}
	scanner := bufio.NewScanner(f)
	y,maxx := 0,0
	for scanner.Scan() {
		for x, acre :=  range scanner.Text(){
			if y == 0 {
				if maxx < x{
					maxx = x
				}
			}
			landscape[xy{x,y}] = acre
		}
		y++
	}
	return area{acres: landscape, maxDimensions: xy{maxx,y}},nil
}
