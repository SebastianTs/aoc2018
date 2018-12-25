package main

import (
	"bufio"
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
	"os"
)

func main() {
	ps, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	g := graph.New(graph.Undirected)
	nodes := make(map[int]graph.Node)
	for i, _ := range ps{
		nodes[i] = g.MakeNode()
	}
	for i, p1 := range ps {
		for j, p2 := range ps{
			if p1.dist(p2) <= 3 {
				g.MakeEdge(nodes[i],nodes[j])
			}
		}
	}
	c := g.StronglyConnectedComponents()
	fmt.Println(len(c))
}

type xyzt struct{
	x,y,z,t int
}

func (p xyzt) String() string{
	return fmt.Sprintf("x:%d y:%d, z:%d t:%d", p.x, p.y, p.z,p.t)
}


func (a *xyzt) dist(b xyzt) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z) + abs(a.t-b.t)
}

func abs(n int) int{
	if n < 0 {
		return -n
	}
	return n
}

func readInput(path string) (ps []xyzt, err error) {
	ps = make([]xyzt,0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return  ps, err
	}
	scanner := bufio.NewScanner(f)

	var x,y,z,t int
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "%d,%d,%d,%d", &x, &y,&z,&t)
		if err != nil {
			return ps, err
		}
		ps = append(ps,xyzt{x,y,z,t})

	}
	return ps, nil
}