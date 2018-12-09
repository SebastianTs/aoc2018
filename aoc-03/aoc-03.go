package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 1000

func main() {
	claims, err := readInput("./input/input-3.txt")
	if err != nil {
		panic(err)
	}

	field := make([][]int, size)
	for i := 0; i < len(field); i++ {
		field[i] = make([]int, size)
	}

	for _, c := range claims {
		for i := 0; i < c.width; i++ {
			for j := 0; j < c.heigth; j++ {
				if field[c.x+i][c.y+j] == 0 {
					field[c.x+i][c.y+j] = c.id
				} else {
					field[c.x+i][c.y+j] = -1
				}
			}
		}
	}
	count := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field); j++ {
			if field[i][j] == -1 {
				count++
			}
		}
	}
	fmt.Println(count)

	//part2
	for _, c := range claims {
		match := true

		for i := 0; i < c.width; i++ {
			for j := 0; j < c.heigth; j++ {
				if field[c.x+i][c.y+j] != c.id {
					match = false
					return
				}
			}
		}
		if match {
			fmt.Println(c.id)
			break
		}
	}
}

type claim struct {
	id     int
	x, y   int
	width  int
	heigth int
}

func readInput(path string) (cs []claim, err error) {
	cs = make([]claim, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return cs, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var id, x, y, width, heigth int
		_, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &heigth)
		if err != nil {
			return cs, err
		}
		cs = append(cs, claim{id: id, x: x, y: y, width: width, heigth: heigth})
	}
	return cs, nil
}
