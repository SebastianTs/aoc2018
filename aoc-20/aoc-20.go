package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	pos := xy{0,0}
	prev := pos
	dist := make(map[xy]int)
	branches := NewStack()
	dir := map[rune]xy{'N':{0,1}, 'E':{1,0}, 'S':{0,-1}, 'W':{-1,0}}

	for _, c := range string(input[1:len(input)-2]){
		switch c{
		case '(':
			branches.Push(pos)
		case ')':
			pos, err = branches.Pop()
			if err != nil{log.Fatal(err)}
		case '|':
			pos, err = branches.Peek()
			if err != nil{log.Fatal(err)}
		default: //NESW
			cur := dir[c]
			pos.x += cur.x
			pos.y += cur.y

			if dist[pos] != 0{
				min := dist[pos]
				if min > dist[prev]+1{
					min = dist[prev]+1
				}
				dist[pos] = min
			} else {
				dist[pos] = dist[prev]+1
			}
		}
		prev = pos
	}
	max := -1
	count := 0
	for _, v := range dist {
		if max < v{
			max = v
		}
		if v >= 1000{
			count++
		}
	}
	fmt.Println("the largest number of doors you would be required to pass is", max)
	fmt.Println(count, "rooms have a shortest path from your current location that pass through at least 1000 doors")
	//TODO Check example output (code fails on some examples)
}

type xy struct{
	x,y int
}

type stack struct {
	s []xy
}

func NewStack() *stack { return &stack{make([]xy, 0)} }

func (s *stack) Peek() (xy, error) {
	l := len(s.s)
	if l == 0 {
		return xy{}, errors.New("Empty Stack")
	}
	return s.s[l-1], nil
}

func (s *stack) Push(v xy) { s.s = append(s.s, v) }

func (s *stack) Pop() (xy, error) {
	l := len(s.s)
	if l == 0 {
		return xy{}, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}