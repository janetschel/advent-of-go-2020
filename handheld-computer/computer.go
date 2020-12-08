package handheld_computer

import (
	"advent-of-go-2020/utils/conv"
	"strings"
)

type InstructionCommand int
const (
	NOP  = iota
	ACC
	JMP
)

func (instruction InstructionCommand) Value() string {
	return [...]string{
		"nop", // NOP
		"acc", // ACC
		"jmp", // JMP
	}[instruction]
}

type Instruction struct {
	instructionCommand InstructionCommand
	value              int
}

func (instruction *Instruction) init(command string, value int) {
	instruction.instructionCommand = getCommandRepresentation(command)
	instruction.value = value
}

func getCommandRepresentation(command string) InstructionCommand {
	if command == "nop" {
		return NOP
	} else if command == "acc" {
		return ACC
	} else if command == "jmp" {
		return JMP
	} else {
		panic("Instruction not valid")
	}
}

type Computer struct {
	pointer int
	instructions []Instruction
}

func CreateComputer(input []string) Computer {
	computer := new(Computer)
	computer.pointer = 0

	for _, instructionString := range input {
		parts := strings.Split(instructionString, " ")

		instruction := new(Instruction)
		instruction.init(parts[0], conv.ToInt(parts[1]))

		computer.instructions = append(computer.instructions, *instruction)
	}

	return *computer
}

func (computer *Computer) Execute() int {
	accumulator := 0

	for ; computer.pointer < len(computer.instructions); computer.pointer++ {
		deltaAcc, deltaPointer := runInstruction(computer.instructions[computer.pointer])
		accumulator += deltaAcc
		computer.pointer += deltaPointer
	}

	return accumulator
}

func runInstruction(instruction Instruction) (int, int) {
	command := instruction.instructionCommand
	value := instruction.value

	if command == NOP {
		return 0,0 // no changes in execution
	} else if command == JMP {
		return 0, value - 1 // -1 since we increment after for-loop in execute
	} else if command == ACC {
		return value, 0
	}

	panic("Not a valid instruction.. panic")
}
