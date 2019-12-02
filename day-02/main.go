package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runProgram(program []int, index int) []int {
	opcode := program[index]
	indexOffset := 4

	switch opcode {
	case 1:
		program[program[index+3]] = program[program[index+1]] + program[program[index+2]]

		program = runProgram(program, index+indexOffset)
		break
	case 2:
		program[program[index+3]] = program[program[index+1]] * program[program[index+2]]

		program = runProgram(program, index+indexOffset)
		break
	case 99:
		break
	default:
		fmt.Println("Something went wrong here")
		break
	}

	return program
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		program := make([]int, 0)
		 for _, value := range strings.Split(scanner.Text(), ",") {
		 	iValue, _ := strconv.Atoi(value)
		 	program = append(program, iValue)
		}

		solvedProgram := runProgram(program, 0)

		fmt.Println(solvedProgram)
	}

	if scanner.Err() != nil {
		fmt.Println("RIP")
	}
}
