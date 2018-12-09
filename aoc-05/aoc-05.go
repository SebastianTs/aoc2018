package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	poly, err := readInput("./input/input-5.txt")
	if err != nil {
		panic(err)
	}
	//Part 1
	fmt.Println("after fully reacting the number of polymer scanned is", react(poly))

	//Part 2
	min := len(poly) + 1
	for i := 0; i < 26; i++ {
		improved := ""
		improved = strings.Replace(poly, string(int('a')+i), "", -1)
		improved = strings.Replace(improved, string(int('A')+i), "", -1)
		cur := react(improved)
		if cur < min {
			min = cur
		}
	}
	fmt.Println("the length of the shortest polymer is", min)

}

func react(poly string) int {
	result := NewStack()

	for _, c := range poly {
		if result.IsEmpty() {
			result.Push(c)
		} else {
			last, _ := result.Peek()
			if last^c == 32 {
				result.Pop()
			} else {
				result.Push(c)
			}
		}
	}
	return result.Size()
}

type stack struct {
	s []rune
}

func NewStack() *stack { return &stack{make([]rune, 0)} }

func (s *stack) IsEmpty() bool { return len(s.s) == 0 }

func (s *stack) Size() int { return len(s.s) }

func (s *stack) Peek() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}
	return s.s[l-1], nil
}

func (s *stack) Push(v rune) { s.s = append(s.s, v) }

func (s *stack) Pop() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func readInput(path string) (str string, err error) {

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return str, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str = scanner.Text()
	}
	return str, nil
}
