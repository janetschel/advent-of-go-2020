package handheld_computer

import (
	"advent-of-go-2020/utils/conv"
	"strings"
)

// Since I think we will need this handheld computer more often in AoC 2020, I refactored it to allow easy
// modifications and extensions of the features of this computer. Intcode v2?

// Run computer.execute() to start the execution of the computer
// Make changes to it if the problem asks for it

// Known instructions to the handheld computer
type InstructionCommand int
const (
	NOP  = iota
	ACC
	JMP
)

// This function returns the string representation of the enum values
func (instruction InstructionCommand) value() string {
	return [...]string{
		"nop", // NOP
		"acc", // ACC
		"jmp", // JMP
	}[instruction]
}

// Struct of actual Instruction to be executed for the handheld computer
type Instruction struct {
	instructionCommand InstructionCommand
	value int
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

func createComputer(input []string) Computer {
	computer := new(Computer)
	computer.pointer = 0 // we start at the beginning

	for _, instructionString := range input {
		parts := strings.Split(instructionString, " ")

		instruction := new(Instruction)
		instruction.init(parts[0], conv.ToInt(parts[1]))

		computer.instructions = append(computer.instructions, *instruction)
	}

	return *computer
}

// Returns accumulator after execution of every instruction
func (computer *Computer) execute() int {
	accumulator := 0

	for ; computer.pointer < len(computer.instructions); computer.pointer++ {
		deltaAcc, deltaPointer := runComputer(computer.instructions[computer.pointer])
		accumulator += deltaAcc
		computer.pointer += deltaPointer
	}

	return accumulator
}

// This function returns (deltaAcc, deltaPointer)
func runComputer(instruction Instruction) (int, int) {
	command := instruction.instructionCommand
	value := instruction.value

	if command == NOP {
		return 0,0 // no changes in execution
	} else if command == JMP {
		return 0, value - 1 // -1 since we increment after for-loop in execute
	} else if command == ACC {
		return value, 0
	}

	// Should never execute, since we stop invalid instructions from entering the computer alltogether
	panic("Not a valid instruction.. panice")
}
