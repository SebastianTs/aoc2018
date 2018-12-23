package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	bots, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	sort.Slice(bots, func(i, j int) bool {
		return bots[i].r > bots[j].r
	})
	count := 0
	for _, bot := range bots {
		if bots[0].inRange(bot){
			count++
		}
	}
	fmt.Println(count, "nanobots are in range of the nanobot with the largest signal radius.")
	fmt.Println(findSpotDistance(bots), "is the shortest manhattan distance between any of those points and 0,0,0")
}

func findSpotDistance(bots []bot) int{ //thx to www.reddit.com/user/seligman99
	var minx,maxx,miny,maxy,minz,maxz int
	for _, bot := range bots{
		if bot.pos.x > maxx {maxx = bot.pos.x}
		if bot.pos.x < minx {minx = bot.pos.x}
		if bot.pos.y > maxy {maxy = bot.pos.y}
		if bot.pos.y < miny {miny = bot.pos.y}
		if bot.pos.z > maxz {maxz = bot.pos.z}
		if bot.pos.z < minz {minz = bot.pos.z}
	}
	dist := 1
	for dist < maxx-minx{
		dist *= 2
	}
	for true{
		var candidateCount int
		var bestPos xyz
		var bestDist int
		for x:=minx; x <= maxx; x +=dist{
			for y:=miny; y <= maxy; y +=dist{
				for z:=minz; z <= maxz; z +=dist{
					count := 0
					for _, bot := range bots{
						calc := abs(bot.pos.x - x) + abs(bot.pos.y - y) + abs(bot.pos.z -z)
						if (calc - bot.r) / dist <= 0{
							count++
						}
					}
					if count > candidateCount {
						candidateCount = count
						bestDist = abs(x) + abs(y) + abs(z)
						bestPos = xyz{x,y,z}
					} else if count == candidateCount {
						if abs(x) + abs(y) + abs(z) < bestDist{
							bestDist = abs(x) + abs(y) + abs(z)
							bestPos = xyz{x,y,z}
						}
					}
				}
			}
		}
		if dist == 1 {
			fmt.Println("Best place to be while beaming is", bestPos)
			return bestDist
		} else {
			minx,maxx = bestPos.x-dist, bestPos.x+dist
			miny,maxy = bestPos.y-dist, bestPos.y+dist
			minz,maxz = bestPos.z-dist, bestPos.z+dist
			dist /= 2
		}
	}
	panic("should not be here")
	return 0
}

type xyz struct{
	x,y,z int
}

func (p xyz) String() string{
	return fmt.Sprintf("x:%d y:%d, z:%d", p.x, p.y, p.z)
}

type bot struct{
	pos xyz
	r int
}

func (b *bot) inRange(a bot) bool{
	dist := abs(a.pos.x - b.pos.x)  + abs(a.pos.y - b.pos.y) + abs(a.pos.z - b.pos.z)
	return dist <= b.r
}

func abs(n int) int{
	if n < 0 {
		return -n
	}
	return n
}

func readInput(path string) (bots []bot, err error) {
	bots = make([]bot,0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return  bots, err
	}
	scanner := bufio.NewScanner(f)

	var x,y,z,r int
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d", &x, &y,&z,&r)
		if err != nil {
			return bots, err
		}
		bots = append(bots,bot{xyz{x,y,z},r})

	}
	return bots, nil
}