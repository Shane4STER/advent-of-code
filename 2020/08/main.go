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

type state struct {
	pc  int
	acc int
}

type runtime struct {
	state   *state
	program program
}

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	program := make(program, 0)

	for scanner.Scan() {
		program = append(program, parseLine(scanner.Text()))
	}

	runtime := createRuntime(program)

	for {
		runtime.exec(true)
	}
}

func createRuntime(program program) runtime {
	var state state
	return runtime{&state, program}
}

func (r runtime) exec(haltOnLoop bool) {
	state := r.state
	if state.pc >= len(r.program) {
		halt()
	}
	cmd := r.program[state.pc]
	if haltOnLoop && cmd.executed {
		r.halt()
	}
	switch cmd.mnemonic {
	case "acc":
		state.acc += cmd.operand
		state.pc++
	case "jmp":
		state.pc += cmd.operand
	case "nop":
		state.pc++
	default:
		log.Fatalf("Unknown command %v\n", cmd.mnemonic)
	}
	cmd.executed = true
}

func (r runtime) halt() {
	r.state.print()
	os.Exit(0)
}

func (s state) print() {
	fmt.Printf("pc: %v acc: %v", s.pc, s.acc)
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
