package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	big   = 99999
	limit = 10000
)

func main() {
	ps, err := readInput("./input/input-6.txt")
	if err != nil {
		panic(err)
	}

	//boundaries
	min_x, min_y := big, big
	max_x, max_y := 0, 0

	for _, p := range ps {
		if p.x > max_x {
			max_x = p.x
		}
		if p.y > max_y {
			max_y = p.y
		}
		if p.x < min_x {
			min_x = p.x
		}
		if p.y < min_y {
			min_y = p.y
		}
	}
	//part 1
	area_size := make(map[point]int)
	infinite := make(map[point]bool)

	for x := min_x; x < max_x+1; x++ {
		for y := min_y; y < max_y+1; y++ {
			min_point := point{x: 0, y: 0}
			min_dist := big
			for _, p := range ps {
				dist := p.distance(point{x: x, y: y})
				if dist < min_dist {
					min_dist = dist
					min_point = p
				} else if dist == min_dist {
					min_point = point{-1, -1}
				}
			}
			if x == max_x || x == min_x || y == max_y || y == min_y {
				infinite[min_point] = true
			}
			area_size[min_point] += 1
		}
	}
	max_area := 0
	for k, v := range area_size {
		if v > max_area && !infinite[k] {
			max_area = v
		}
	}
	fmt.Println(max_area)

	//part2
	count := 0
	for x := min_x; x < max_x+1; x++ {
		for y := min_y; y < max_y+1; y++ {
			total_dist := 0
			for _, p := range ps {
				total_dist += p.distance(point{x: x, y: y})
				if total_dist >= limit {
					break
				}
			}
			if total_dist < limit {
				count++
			}
		}
	}
	fmt.Println(count)

}

type point struct {
	x, y int
}

func (p *point) distance(q point) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readInput(path string) (ps []point, err error) {
	ps = make([]point, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return ps, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var x, y int
		_, err := fmt.Sscanf(scanner.Text(), "%d, %d", &x, &y)
		if err != nil {
			return ps, err
		}
		ps = append(ps, point{x: x, y: y})
	}
	return ps, nil
}
