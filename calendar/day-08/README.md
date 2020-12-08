# Day 8: Handheld Halting
## Problem Summary ([?](https://adventofcode.com/2020/day/8))

This problem requires to write a small computer which executes instructions and returns an answer-value.

Instructions are structured like this:  
`instruction value`

The computer consists of following instructions:
- `acc` increases or decreases a single global value called the `accumulator` by the value given in the argument.
- `jmp` jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction.
- `nop` stands for No operation - it does nothing.

Suppose following set of instructions:
```
acc +10
acc -10
nop +100
```

Stepping through each instruction, the accumulator at the end will be 0, since we add and subtract 10 from it, and `nop` does nothing.

Suppose a more complex set of instructions:
```
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
```

The instructions are visited in following order:
```
nop +0  | 1
acc +1  | 2, 8(!)
jmp +4  | 3
acc +3  | 6
jmp -3  | 7
acc -99 |
acc +1  | 4
jmp -4  | 5
acc +6  |
```

Once an instruction is visited twice, we know the computer will never halt, since we enter an infinite loop of repeating instructions.  
The problem here is, that we visit the second line `acc +1` twice, so the computer will **never halt**.

For **part 1** we need to find the value in the accumulator **right before** an instruction is executed the second time.

For **part 2** we need to find an instruction - when changed - will result in the computer terminating, and give back the result at the end as an answer.  
Clarification:
- The set of instructions in our input will run forever if we make no changes to it.
- If we change **ONE** `nop` instruction to `jmp` **OR** `jmp` to `nop` somewhere in the code, our computer will halt successfully.
- We need to find the value in the accumulator if the program has halted. For that we first need to find the instruction to change, so the computer WILL halt.

My solution for **part 1** is `2025`.  
My solution for **part 2** is `2001`.

## Recap
Pretty easy problem compared to yesterday.  
I had to do virtually no parsing at all, and it was very fun to tackle the problem.  

I - however - have **VERY STRONG** feelings, that this will be the `new Intcode` of 2020.

Since last year I never really refactored my Intcode, I had many bugs in it.  
This year, to counter this problem, I have not reworked my original solutions and just left them be, I however have written a complete new `computer.go` file (go have a look at it) which I can reuse to run the computer, debug it, and make easy changes to it.

I hope this work pays off this year 😉

Since this was a very easy problem, my solve-times were okay. I had some problems with part 2, but worked them out rather quickly.  
Here are my solve-times:
- `00:10:15` for part 1 (#2365)
- `00:26:30` for part 2 (#2590)
