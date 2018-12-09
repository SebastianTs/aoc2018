package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	records, err := readInput("./input/input-4.txt")
	if err != nil {
		panic(err)
	}
	year := func(r1, r2 *record) bool { return r1.year < r2.year }
	month := func(r1, r2 *record) bool { return r1.month < r2.month }
	day := func(r1, r2 *record) bool { return r1.day < r2.day }
	hour := func(r1, r2 *record) bool { return r1.hour < r2.hour }
	minute := func(r1, r2 *record) bool { return r1.minute < r2.minute }

	OrderedBy(year, month, day, hour, minute).Sort(records)

	id := 0
	start := time.Now()
	duration := make(map[int]time.Duration)
	minutes := make(map[int][60]int)

	for _, r := range records {
		switch r.event[0] {
		case 'G': //Guard #10 begins shift
			id, err = getGuardID(r.event)
			if err != nil {
				log.Fatal(err)
			}
		case 'f': // falls asleep
			start = time.Date(r.year, time.Month(r.month), r.day, r.hour, r.minute, 0, 0, time.UTC)
		case 'w': //wakes up
			stop := time.Date(r.year, time.Month(r.month), r.day, r.hour, r.minute, 0, 0, time.UTC)
			duration[id] += stop.Sub(start)

			for m := start.Minute(); m < stop.Minute(); m++ {
				a := minutes[id]
				a[m]++
				minutes[id] = a
			}
		}
	}
	//Part1 most minutes asleep
	max := 0
	current := 0
	for guard, dur := range duration {
		if max < int(dur.Minutes()) {
			max = int(dur.Minutes())
			current = guard
		}
	}
	max = 0
	idx := 0
	for i := 0; i < len(minutes[current]); i++ {
		if max < minutes[current][i] {
			max = minutes[current][i]
			idx = i
		}
	}
	fmt.Println("Guard", current, "Minute", idx, "Answer", current*idx)

	//Part 2 most frequently asleep on the same minute
	max = 0
	idx = 0
	current = 0
	for guard, times := range minutes {
		for i := 0; i < len(times); i++ {
			if times[i] > max {
				max = times[i]
				current = guard
				idx = i
			}
		}
	}
	fmt.Println("Guard", current, "Minute", idx, "Answer", current*idx)
}

func getGuardID(event string) (int, error) {
	var text string
	var id int
	_, err := fmt.Sscanf(event, "%s #%d", &text, &id)
	if err != nil {
		return 0, err
	}
	return id, err
}

type record struct {
	year, month, day, hour, minute int
	event                          string
}

type lessFunc func(p1, p2 *record) bool

type multiSorter struct {
	records []record
	less    []lessFunc
}

func (ms *multiSorter) Sort(records []record) {
	ms.records = records
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.records)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.records[i], ms.records[j] = ms.records[j], ms.records[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.records[i], &ms.records[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

func readInput(path string) (rs []record, err error) {
	rs = make([]record, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return rs, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var year, month, day, hour, minute int
		var token0, token1 string
		_, err := fmt.Sscanf(scanner.Text(), "[%d-%d-%d %d:%d] %s %s", &year, &month, &day, &hour, &minute, &token0, &token1)
		if err != nil {
			return rs, err
		}
		event := token0 + " " + token1
		rs = append(rs, record{year, month, day, hour, minute, event})
	}
	return rs, nil
}
