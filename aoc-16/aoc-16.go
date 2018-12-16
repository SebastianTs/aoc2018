package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	numberOfOperations = 16
	splitLine = 3129
)

func main() {
	vs, ops, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}
	fns := []opFunc{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	mapping := make(map[int]map[int]struct{})
	// setup
	gCount := 0
	for i:=0; i < numberOfOperations; i++{
		mapping[i] = make(map[int]struct{})
	}

	// looking for functions
	for _, v := range vs{
		index := matchRegistersToOperations(v,fns)
		if len(index) >= 3 {
			gCount++
		}
		for j:=0; j < len(index); j++{
			mapping[v.op.opCode][index[j]] = struct{}{}
		}
	}
	fmt.Println( gCount, "behave like three or more opcodes")

	ready := false
	opCodes := make(map[int]opFunc)
	for !ready{
		for k,v := range mapping{
			if len(v) == 1 {
				for code := range v{
					opCodes[k] = fns[code]
					for _, candidates := range mapping{
						delete(candidates,code)
					}
				}
				if len(opCodes) == numberOfOperations {
					ready = true
				}
			}
		}
	}
	g := register{}
	for _, op := range ops{
		g = opCodes[op.opCode](op.A,op.B,op.O,g)
	}

	fmt.Println(g[0], "is contained in register 0 after executing the test program")

}

func matchRegistersToOperations(n sample,fn []opFunc) (index []int){
	for i, f := range fn{
		cur := f(n.op.A, n.op.B, n.op.O, n.before)
		if cur == n.after{
				index = append(index,i)
		}
	}
	return index
}

type opFunc func (a,b,c int, r register) register

// (add register) stores into register C the result of adding register A and register B.
func addr(a,b,c int, r register) register{
	r[c] = r[a] + r[b]
	return r
}

func addi(a,b,c int, r register) register{
	r[c] = r[a] + b
	return r
}

func mulr(a,b,c int, r register) register{
	r[c] = r[a]*r[b]
	return r
}

func muli(a,b,c int, r register) register{
	r[c] = r[a] * b
	return r
}

func banr(a,b,c int, r register) register{
	r[c] = r[a] & r[b]
	return r
}

func bani(a,b,c int, r register) register{
	r[c] = r[a] & b
	return r
}

func borr(a,b,c int, r register) register{
	r[c] = r[a] | r[b]
	return r
}

func bori(a,b,c int, r register) register{
	r[c] = r[a] | b
	return r
}

func setr(a,b,c int, r register) register{
	r[c] = r[a]
	return r
}

func seti(a,b,c int, r register) register{
	r[c] = a
	return r
}

func gtir(a,b,c int, r register) register{
	if a > r[b] {
		r[c] = 1
		} else {
			r[c] = 0
	}
	return r
}

func gtri(a,b,c int, r register) register{
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtrr(a,b,c int, r register) register{
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqir(a,b,c int, r register) register{
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqri(a,b,c int, r register) register{
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqrr(a,b,c int, r register) register{
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

type sample struct {
	before register
	op operation
	after register
}

type register [4]int
type operation struct{
	opCode int
	A      int
	B      int
	O      int
}

func readInput(path string) (vs []sample, testprogram []operation, err error) {
	vs = make([]sample, 0)
	testprogram = make([]operation, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return vs, testprogram, err
	}
	scanner := bufio.NewScanner(f)
	var aBefore, bBefore, cBefore, dBefore, aAfter, bAfter, cAfter, dAfter, ins, aInput, bInput, output int
	count := 0
	for scanner.Scan() {
		if count > splitLine {
			_, err = fmt.Sscanf(scanner.Text(), "%d %d %d %d", &ins, &aInput, &bInput, &output)
			if err != nil {
				return vs, testprogram, err
			}
			testprogram = append(testprogram, operation{ins, aInput, bInput, output})

		} else {
			switch count % 4 {
			case 0:
				_, err = fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &aBefore, &bBefore, &cBefore, &dBefore)
			case 1:
				_, err = fmt.Sscanf(scanner.Text(), "%d %d %d %d", &ins, &aInput, &bInput, &output)
			case 2:
				_, err = fmt.Sscanf(scanner.Text(), "After: [%d, %d, %d, %d]", &aAfter, &bAfter, &cAfter, &dAfter)
			case 3:
				if err != nil {
					return vs, testprogram, err
				}
				before := register{aBefore, bBefore, cBefore, dBefore}
				after := register{aAfter, bAfter, cAfter, dAfter}
				instruction := operation{ins, aInput, bInput, output}
				vs = append(vs, sample{before, instruction, after})
			}
		}
		count++

	}
	return vs, testprogram, nil
}
