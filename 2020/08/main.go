package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	mnemonic string
	operand  int
	executed bool
}

type program []*command

type State struct {
	pc        int
	acc       int
	executed  map[int]bool
	changedPC int
}

type runtime struct {
	state   *State
	program program
}

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	program := make(program, 0)

	for scanner.Scan() {
		program = append(program, parseLine(scanner.Text()))
	}

	runtime := createRuntime(program)

	endState := runtime.exec(State{0, 0, make(map[int]bool), -1})

	endState.print()

}

func createRuntime(program program) runtime {
	var state State
	return runtime{&state, program}
}

func (r runtime) exec(state State) State {
	if state.pc < -1 {
		state.print()
		return State{-2, state.acc, state.executed, state.changedPC}
	}
	var nextPC int
	if state.pc >= len(r.program) {
		fmt.Println("exited successfully")
		state.print()
		os.Exit(1)
		return State{-1, state.acc, state.executed, state.changedPC}
	}
	cmd := r.program[state.pc]
	if _, exists := state.executed[state.pc]; exists {
		fmt.Println("LOOP DETECTED: returning current state")
		return State{-2, state.acc, state.executed, state.changedPC}
	} else {
		state.executed[state.pc] = true
	}
	executed := copyMap(state.executed)
	switch cmd.mnemonic {
	case "acc":
		state.acc += cmd.operand
		state.pc++
	case "jmp":
		if state.changedPC < 0 {
			nextPC = state.pc + 1
			r.exec(State{nextPC, state.acc, executed, nextPC})
		}
		state.pc += cmd.operand
	case "nop":
		if state.changedPC < 0 {
			nextPC = state.pc + cmd.operand
			r.exec(State{nextPC, state.acc, executed, nextPC})
		}
		state.pc++
	default:
		log.Fatalf("Unknown command %v\n", cmd.mnemonic)
	}
	cmd.executed = true
	return r.exec(state)
}

func (c *command) flipMnemonic() {
	if c.mnemonic == "jmp" {
		c.mnemonic = "nop"
	} else {
		c.mnemonic = "jmp"
	}
}

func (r runtime) cmdCausesLoop(cmd command) bool {
	var nextpc int
	state := r.state
	if state.pc > len(r.program) {
		return false
	}
	switch cmd.mnemonic {
	case "acc":
		return false
	case "jmp":
		nextpc = state.pc + cmd.operand
	case "nop":
		nextpc = state.pc + 1
	default:
		log.Fatalf("Unknown command %v\n", cmd.mnemonic)
	}
	return r.program[nextpc].executed
}

func (r runtime) halt() {
	r.state.print()
	//os.Exit(0)
}

func (s State) print() {
	fmt.Printf("pc: %v acc: %v changed: %v\n", s.pc, s.acc, s.changedPC)
}

func parseLine(line string) *command {
	var cmd command

	splitCmd := strings.Split(line, " ")

	cmd.mnemonic = splitCmd[0]

	opString := splitCmd[1]
	op, err := strconv.Atoi(opString)
	if err != nil {
		log.Fatal(err)
	}
	cmd.operand = op

	return &cmd
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return r
}

func copyMap(m map[int]bool) map[int]bool {
	cp := make(map[int]bool)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}
