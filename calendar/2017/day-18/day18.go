package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type operation string
const (
	sound operation = "snd"
	send = "snd"
	set = "set"
	add = "add"
	multiply = "mul"
	modulo = "mod"
	recover = "rcv"
	receive = "rcv"
	jumpGreaterThanZero = "jgz"
)

var waitGroup sync.WaitGroup

func main() {
	input := files.ReadFile(18, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	registers := map[string]int{}
	lastFrequency := 0
	increment := 1
	for i := 0; i < len(input); i += increment {
		increment = 1
		parts := strings.Fields(input[i])
		register := parts[1]

		y := 0
		if len(parts) > 2 {
			var isRegister bool
			y, isRegister = registers[parts[2]]
			if !isRegister {
				y, _ = strconv.Atoi(parts[2])
			}
		}
		switch operation(parts[0]) {
		case sound:
			value, isRegister := registers[register]
			if !isRegister {
				value, _ = strconv.Atoi(register)
			}
			lastFrequency = value
		case set:
			registers[register] = y
		case add:
			registers[register] += y
		case multiply:
			registers[register] *= y
		case modulo:
			registers[register] %= y
		case recover:
			value, isRegister := registers[register]
			if !isRegister {
				value, _ = strconv.Atoi(register)
			}
			if value != 0 {
				return lastFrequency
			}
		case jumpGreaterThanZero:
			value, isRegister := registers[register]
			if !isRegister {
				value, _ = strconv.Atoi(register)
			}
			if value > 0 {
				increment = y
			}
		}
	}

	return result
}

func solvePart2(input []string) int {
	queue0, queue1 := make(chan int, 100), make(chan int, 100)
	registers0, registers1 := map[string]int{ "p": 0 }, map[string]int{ "p": 1 }
	
	waitGroup.Add(1)
	go runProgram(input, registers0, queue0, queue1)
	waitGroup.Add(1)
	go runProgram(input, registers1, queue1, queue0)
	waitGroup.Wait()
	
	return registers1["sent"]
}

func runProgram(instructions []string, registers map[string]int, sendQueue chan int, receiveQueue chan int) {
	increment := 1
	for i := 0; i < len(instructions); i += increment {
		increment = 1
		parts := strings.Fields(instructions[i])
		register := parts[1]

		y := 0
		if len(parts) > 2 {
			var isRegister bool
			y, isRegister = registers[parts[2]]
			if !isRegister {
				y, _ = strconv.Atoi(parts[2])
			}
		}
		switch operation(parts[0]) {
		case send:
			value, isRegister := registers[register]
			if !isRegister {
				value, _ = strconv.Atoi(register)
			}
			sendQueue <- value
			registers["sent"]++
		case set:
			registers[register] = y
		case add:
			registers[register] += y
		case multiply:
			registers[register] *= y
		case modulo:
			registers[register] %= y
		case receive:
			received, timeout := receiveFromChannel(receiveQueue)
			if timeout {
				waitGroup.Done()
			}
			registers[register] = received
		case jumpGreaterThanZero:
			value, isRegister := registers[register]
			if !isRegister {
				value, _ = strconv.Atoi(register)
			}
			if value > 0 {
				increment = y
			}
		}
	}
	waitGroup.Done()
}

func receiveFromChannel(channel chan int) (int, bool) {
	var receive int
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	select {
	case receive = <-channel:
		return receive, false
	case <-timeout:
		fmt.Println("Timeout")

		return receive, true
	}
}
