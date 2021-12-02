package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	d int
}

type Direction int

const (
	UP Direction = iota
	DOWN
	FORWARD
)

type Command struct {
	direction Direction
	scalar    int
}

func main() {
	var subPos Position
	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		line := scanner.Text()
		cmd, err := stringToCommand(line)
		if err != nil {
			log.Fatal(err)
		}

		subPos.move(cmd)

	}

	fmt.Printf("The Sub has a horizontal position of %v\n", subPos.hPos())
}

func stringToCommand(str string) (Command, error) {
	var command Command

	commandSlice := strings.Fields(str)

	if commandSlice[0] == "up" {
		command.direction = UP
	} else if commandSlice[0] == "down" {
		command.direction = DOWN
	} else if commandSlice[0] == "forward" {
		command.direction = FORWARD
	} else {
		err := errors.New("invalid command given")
		return Command{0, 0}, err
	}

	scalar, err := strconv.Atoi(commandSlice[1])

	if err != nil {
		return Command{0, 0}, err
	}
	command.scalar = scalar
	return command, nil
}

func (pos *Position) move(cmd Command) {
	if cmd.direction == UP {
		pos.d -= cmd.scalar
	} else if cmd.direction == DOWN {
		pos.d += cmd.scalar
	} else if cmd.direction == FORWARD {
		pos.x += cmd.scalar
	}
}

func (pos *Position) hPos() int {
	return pos.d * pos.x
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
