package main

import (
	"fmt"
	"math"
)

const(
	gridMaxX      = 300
	gridMaxY      = 300
	maxSquareSize = 300
	serial        = 1723
)

func main() {
	sum := make(map[xy]int)
	var bx,by,bs int
	best := math.MinInt32

	for x:= 1; x <= gridMaxY; x++{
		for y:= 1; y <= gridMaxY; y++{
			id := x + 10
			p := id * y + serial
			p = (p*id) / 100 % 10 - 5
			// https://de.wikipedia.org/wiki/Integralbild
			sum[xy{x,y}]	= p + sum[xy{x,y-1}] + sum[xy{x-1,y}] - sum[xy{x-1,y-1}]
		}
	}

	for x := 3; x <= gridMaxY; x++{
		for y:= 3; y <= gridMaxY; y++{
			total := sum[xy{x,y}] - sum[xy{x,y-3}] - sum[xy{x-3,y}] + sum[xy{x-3,y-3}]
			if total > best{
				best = total
				bx = x
				by = y
			}
		}
	}
	fmt.Println("The X,Y coordinate of the top-left fuel cell of the 3x3 square with the largest total power is", bx-2,by-2)
	best = math.MinInt32
	for size := 1; size <= maxSquareSize; size++{
		for x := size; x <= gridMaxX; x++{
			for y := size; y <= gridMaxY; y++{
				total := sum[xy{x,y}] - sum[xy{x,y-size}] - sum[xy{x-size,y}] + sum[xy{x-size,y-size}]
				if total > best{
					best = total
					bx = x
					by = y
					bs = size
				}
			}
		}
	}
	fmt.Println("The X,Y,size identifier of the square with the largest total power is",bx-bs+1,by-bs+1,bs)

}

type xy struct {
	x,y int
}


