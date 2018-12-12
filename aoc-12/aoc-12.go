package main

import (
	"bufio"
	"os"
	"strings"
)

const(
 	generations = 20
 	)


func main() {
	state, rules, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}

	for i:=0; i<generations;i++{
		applie(state, rules)
	}

}

func applie(state string,rules map[string]rune) {
	
}

func readInput(path string) (state string, rules map[string]rune, err error) {
	rules = make(map[string]rune)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return state, rules,  err
	}
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		if count == 0 {
			state = scanner.Text()[15:]
		} else if count > 1 {
		token := strings.Fields(scanner.Text())
			rules[token[0]] = rune(token[2][0])
		}
		count++
	}
	return state, rules, nil
}
