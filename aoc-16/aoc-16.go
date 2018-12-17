package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	numberOfOpCodes = 16
	splitLine       = 3129
)

func main() {
	samples, testroutine, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}
	fns := []opFunc{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	OpCodeCandidatesSet := make(map[int]map[int]struct{}) // 01 -> [08, 12, ... ]
	candidateCount := 0
	for i := 0; i < numberOfOpCodes; i++ {
		OpCodeCandidatesSet[i] = make(map[int]struct{})
	}

	for _, sample := range samples {
		candidateIndex := matchSampleToOperationFunctions(sample, fns)
		if len(candidateIndex) >= 3 {
			candidateCount++
		}
		for j := 0; j < len(candidateIndex); j++ {
			OpCodeCandidatesSet[sample.op.Code][candidateIndex[j]] = struct{}{}
		}
	}
	fmt.Println(candidateCount, "behave like three or more opcodes")

	ready := false
	if len(samples) == 0 {
		ready = true
	} //avoid infinite loop when input file was empty
	opCodeMap := make(map[int]opFunc)
	for !ready {
		for code, set := range OpCodeCandidatesSet {
			if len(set) == 1 {
				for idx := range set {
					opCodeMap[code] = fns[idx]
					for _, candidates := range OpCodeCandidatesSet {
						delete(candidates, idx)
					}
				}
				if len(opCodeMap) == numberOfOpCodes {
					ready = true
				}
			}
		}
	}
	testRegister := register{}
	for _, op := range testroutine {
		testRegister = opCodeMap[op.Code](op.A, op.B, op.C, testRegister)
	}
	fmt.Println(testRegister[0], "is contained in register 0 after executing the test program")
}

// matchSampleToOperationFunctions applies a list of functions to the given sample
// and returns the index of functions with valid registers.
func matchSampleToOperationFunctions(s sample, fn []opFunc) (index []int) {
	for i, f := range fn {
		cur := f(s.op.A, s.op.B, s.op.C, s.before)
		if cur == s.after {
			index = append(index, i)
		}
	}
	return index
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

// sample shows the effect of the instruction on the registers
type sample struct {
	before register
	op     operation
	after  register
}

// the device has four registers (numbered 0 through 3)
type register [4]int

// Every instruction consists of four values:
// an opcode, two inputs (named A and B), and an output (named C), in that order.
// The opcode specifies the behavior of the instruction and how the inputs are interpreted.
// The output, C, is always treated as a register.
type operation struct {
	Code int
	A    int
	B    int
	C    int
}

func readInput(path string) (samples []sample, testprogram []operation, err error) {
	samples = make([]sample, 0)
	testprogram = make([]operation, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return samples, testprogram, err
	}
	scanner := bufio.NewScanner(f)
	var rBefore, rAfter [4]int
	var ins, aInput, bInput, cInput int
	count := 0
	for scanner.Scan() {
		if count <= splitLine { // These are the samples for part 1
			switch count % 4 {
			case 0:
				_, err = fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &rBefore[0], &rBefore[1], &rBefore[2], &rBefore[3])
			case 1:
				_, err = fmt.Sscanf(scanner.Text(), "%d %d %d %d", &ins, &aInput, &bInput, &cInput)
			case 2:
				_, err = fmt.Sscanf(scanner.Text(), "After: [%d, %d, %d, %d]", &rAfter[0], &rAfter[1], &rAfter[2], &rAfter[3])
			case 3:
				if err != nil {
					return samples, testprogram, err
				}
				before := register(rBefore)
				after := register(rAfter)
				instruction := operation{ins, aInput, bInput, cInput}
				samples = append(samples, sample{before, instruction, after})
			}
		} else { // This is the testroutine for part 2
			_, err = fmt.Sscanf(scanner.Text(), "%d %d %d %d", &ins, &aInput, &bInput, &cInput)
			if err != nil {
				return samples, testprogram, err
			}
			testprogram = append(testprogram, operation{ins, aInput, bInput, cInput})
		}
		count++
	}
	return samples, testprogram, nil
}
