package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	is, err := readInput("./input/input-8.txt")
	if err != nil {
		panic(err)
	}
	i := 0
	root := parse(is, &i)
	fmt.Println("The sum of all metadata entries is:\t", root.Sum())
	fmt.Println("The value of the root node is:\t\t" ,root.Value())
}

type Node struct {
	metas []int
	nodes []Node
}

func parse(list []int, i *int) Node{
	children := list[*i]
	*i++
	metadata := list[*i]
	*i++
	node := Node{metas: make([]int,metadata),
		         nodes: make([]Node,children),
			}
	for j:=0; j < children; j++ {
		node.nodes[j] = parse(list, i)
	}
	for j:=0;j < metadata; j++{
		node.metas[j] = list[*i]
		*i++
	}
	return node
}

func (n *Node) Sum() int{
	var sumChilds int
	for _, node := range n.nodes{
		sumChilds += node.Sum()
	}
	var sumMetadata int
	for _, k := range n.metas{
		sumMetadata += k
	}
	return sumChilds + sumMetadata
}

func (n *Node) Value() int{
	if len(n.nodes) == 0 {
		var sumMetadata int
		for _, k := range n.metas{
			sumMetadata += k
		}
		return sumMetadata
	}
	var value int
	for _, m := range n.metas{
		if m < len(n.nodes)+1 {
			value += n.nodes[m-1].Value()
		}
	}
	return value
}


func readInput(path string) (is []int, err error) {
	is = make([]int, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return is, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		if err != nil {
			return is, err
		}
		for _, el := range s{
			var number int
			_, err := fmt.Sscanf(el,"%d", &number)
			if err != nil{
				return is, err
			}
			is = append(is,number)
		}
	}
	return is, nil
}
