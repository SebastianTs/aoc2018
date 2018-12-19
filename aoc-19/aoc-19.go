package main

import (
	"bufio"
	"fmt"
	"os"
)

const myPuzzleRegister = 3
var (
	opCodeMap = map[string]opFunc{
		"addr": addr, "addi": addi, "mulr": mulr, "muli": muli,
		"banr": banr, "bani": bani, "borr": borr, "bori": bori,
		"setr": setr, "seti": seti, "gtir": gtir, "gtri": gtri,
		"gtrr": gtrr, "eqir": eqir, "eqri": eqri, "eqrr": eqrr,
	}
)

func main() {
	program, ip, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	cur := register{}
	for i:=cur[ip]; i < len(program); {
		p := program[i]
		cur[ip] = i
		cur = p.Code(p.A,p.B,p.C,cur)
		i = cur[ip]+1
	}
	fmt.Println("The value left in register 0 is", cur[0])

	cur = register{1,0,0,0,0,0}
	for i:=cur[ip]; i < len(program); {
		p := program[i]
		cur[ip] = i
		cur = p.Code(p.A,p.B,p.C,cur)
		i = cur[ip]+1
		if i == 1{
			n := cur[myPuzzleRegister]
			sum := 0
			for j:=1; j<=n;j++{
				if n % j == 0{
					sum +=j
				}
			}
			fmt.Println("The value left in register 0 is", sum)
			break
		}
	}
}


// opFunc is the signature for a function that can be applied to an register
type opFunc func(a, b, c int, r register) register

// (add register) stores into register C the result of adding register A and register B.
func addr(a, b, c int, r register) register {
	r[c] = r[a] + r[b]
	return r
}

// addi (add immediate) stores into register C the result of adding register A and value B.
func addi(a, b, c int, r register) register {
	r[c] = r[a] + b
	return r
}

// mulr (multiply register) stores into register C the result of multiplying register A and register B.
func mulr(a, b, c int, r register) register {
	r[c] = r[a] * r[b]
	return r
}

// muli (multiply immediate) stores into register C the result of multiplying register A and value B.
func muli(a, b, c int, r register) register {
	r[c] = r[a] * b
	return r
}

// banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
func banr(a, b, c int, r register) register {
	r[c] = r[a] & r[b]
	return r
}

// bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func bani(a, b, c int, r register) register {
	r[c] = r[a] & b
	return r
}

// borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
func borr(a, b, c int, r register) register {
	r[c] = r[a] | r[b]
	return r
}

// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
func bori(a, b, c int, r register) register {
	r[c] = r[a] | b
	return r
}

// setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func setr(a, b, c int, r register) register {
	r[c] = r[a]
	return r
}

// seti (set immediate) stores value A into register C. (Input B is ignored.)
func seti(a, b, c int, r register) register {
	r[c] = a
	return r
}

// gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
func gtir(a, b, c int, r register) register {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

// gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func gtri(a, b, c int, r register) register {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

// gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
func gtrr(a, b, c int, r register) register {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

// eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
func eqir(a, b, c int, r register) register {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

// eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func eqri(a, b, c int, r register) register {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

// eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
func eqrr(a, b, c int, r register) register {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

type register [6]int

type operation struct {
	Code opFunc
	A    int
	B    int
	C    int
}

func readInput(path string) (program []operation, ip int, err error) {
	program = make([]operation, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return  program,0, err
	}
	scanner := bufio.NewScanner(f)

	var ins string
	var aInput, bInput, cInput int
	count := 0
	for scanner.Scan() {
		if count == 0 {
				_, err = fmt.Sscanf(scanner.Text(), "#ip %d", &ip)
				if err != nil {
					return program, 0, err
				}

		} else {
			_, err = fmt.Sscanf(scanner.Text(), "%s %d %d %d", &ins, &aInput, &bInput, &cInput)
			if err != nil {
				return program, ip, err
			}
			program = append(program, operation{opCodeMap[ins], aInput, bInput, cInput})
		}
		count++
	}
	return program, ip, nil
}
