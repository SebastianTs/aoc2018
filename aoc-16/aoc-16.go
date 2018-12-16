package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	vs, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}
	fns := []opFunc{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	gCount := 0
	for _, v := range vs{
		if check(v,fns) > 2 {
			gCount++
		}
	}
	fmt.Println(gCount)

}

func check(n note,fn []opFunc) (count int){
	for _, f := range fn{
		cur := f(n.op.inputA, n.op.inputB, n.op.output, n.before)
		if cur == n.after{
				count++
		}
	}
	return count
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

type note struct {
	before register
	op operation
	after register
}

type register [4]int
type operation struct{
	instruction int
	inputA int
	inputB int
	output int
}
func readInput(path string) (vs []note, err error) {
	vs = make([]note,0)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return vs,  err
	}
	scanner := bufio.NewScanner(f)
	var aBefore,bBefore,cBefore,dBefore,aAfter,bAfter,cAfter, dAfter,ins, aInput,bInput,output int
	count := 0
	for scanner.Scan() {
		 switch count%4{
		 case 0:
			_, err = fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &aBefore, &bBefore, &cBefore, &dBefore)
		 case 1:
			 _, err = fmt.Sscanf(scanner.Text(), "%d %d %d %d", &ins, &aInput, &bInput, &output)
		 case 2:
			 _, err = fmt.Sscanf(scanner.Text(), "After: [%d, %d, %d, %d]", &aAfter, &bAfter, &cAfter, &dAfter)
		 case 3:
			 if err != nil {
				 return vs, err
			 }
			 before := register{aBefore,bBefore,cBefore,dBefore}
			 after := register{aAfter,bAfter,cAfter,dAfter}
			 instruction := operation{ins, aInput, bInput, output}
			 vs = append(vs,note{before,instruction,after})
		}
		 count++
		 if count == 3129{
		 	break
		 }
	}
	return vs, nil
}
