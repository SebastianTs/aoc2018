package main

import (
	"bufio"
	"fmt"
	"os"
)

const(
	big = 100000
 	duration = 100000
 	)


func main() {
	vs, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}
	untouched := make([]vector,len(vs))
	copy(untouched, vs)

	distances := make([]int,duration)
	for i:=0; i<duration; i++{
		distances[i] = minxDistance(vs)
		for j:=0; j<len(vs);j++ {
			vs[j].update()
		}
	}

	minx := big
	seconds := 0
	for i, distance := range distances {
		if minx > distance {
			minx = distance
			seconds = i
		}
	}

	for i:=0; i<seconds; i++{
		for j:=0; j<len(untouched); j++ {
			untouched[j].update()
		}
	}
	printSky(untouched)
	fmt.Println("they have needed to wait for that message to appear:", seconds)

}

type xy struct {
	x,y int
}

type vector struct{
	p xy
	right,up int
}

func minxDistance(vs []vector) int{
	minx, _, maxx, _ := getBoundaries(vs)
	return abs(maxx - minx)
}

func abs (n int) int{
	if n < 0 {
		return -n
	}
	return n
}

func  getBoundaries(vs []vector) (minx, miny, maxx, maxy int){
	minx, miny = big, big
	maxx, maxy = -big, -big

	for _, v := range vs {
		if v.p.x > maxx {
			maxx = v.p.x
		}
		if v.p.y > maxy {
			maxy = v.p.y
		}
		if v.p.x < minx {
			minx = v.p.x
		}
		if v.p.y < miny {
			miny = v.p.y
		}
	}
	return minx, miny, maxx, maxy
}

func (v *vector) update(){
	v.p.x += v.right
	v.p.y += v.up
}

func printSky(vs []vector ){

	stars := make(map[xy]bool)
	for _, v := range vs {
		stars[v.p] = true
	}
	minx, miny, maxx, maxy := getBoundaries(vs)
	for y:= miny; y <= maxy; y++{
		for x:=minx; x <= maxx; x++{
			if stars[xy{x,y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func readInput(path string) (vs []vector, err error) {
	vs = make([]vector,0)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return vs,  err
	}
	scanner := bufio.NewScanner(f)
	var x,y,right,up int

	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "position=<%d,%d> velocity=<%d,%d>", &x, &y, &right, &up)
		if err != nil {
			return vs, err
		}
		vs = append(vs,vector{xy{x,y},right, up})
	}
	return vs, nil
}
