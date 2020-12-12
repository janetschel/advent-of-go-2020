# Day 12: Rain Risk
## Problem Summary ([?](https://adventofcode.com/2020/day/12))

For this problem we have to navigate a ship through water with helps of given instructions.

Instructions are made of an **action**, and an **amount**.  
For example:
```
F10
N3
F7
R90
F11
```

- Action `N` means to move north by the given value.
- Action `S` means to move south by the given value.
- Action `E` means to move east by the given value.
- Action `W` means to move west by the given value.
- Action `L` means to turn left the given number of degrees.
- Action `R` means to turn right the given number of degrees.
- Action `F` means to move forward by the given value in the direction the ship is currently facing.

For **part 1** we need to execute every given instruction and calculate the [Manhattan distance ](https://en.wikipedia.org/wiki/Taxicab_geometry) from the starting point to the point where the ship ends up.

For **part 2** the rules actually change a bit:  
Almost all the actions indicate how to move **a waypoint** which is **relative to the ship's position**:  
- Action `N` means to move **the waypoint** north by the given value.
- Action `S` means to move the waypoint** south by the given value.
- Action `E` means to move the waypoint east by the given value.
- Action `W` means to move the waypoint west by the given value.
- Action `L` means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
- Action `R` means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
- Action `F` means to **move forward TO the waypoint** a number of times equal to the given value.

For **part 2** we again have to calculate the Manhattan distance, but with the changed set of rules.

My solution for **part 1** is `562`.  
My solution for **part 2** is `101860`.  

## Recap

From day to day, I am hoping more for some `Intcode v2.0` problems, but I had very much fun solving this problem instead indeed.

I hade some little problems with **part 1**.  
I ran my code, hoping it would work, and got a solution (the right one, but I didn't know that) which was smaller than `1.000`, which I then commited **WITH A TYPO**.  

Of course the solution was incorrect... So I started debugging my code with absolutely no bugs in it.  
After 30 minutes I was so out of hope finding a bug that I actually resubmitted my previous answer: **AND IT WORKED!** HEUREKA

I sadly lost so many precious minutes while working on **part 1**, not because I couldn't read (and understand the problem) but because I couldn't type (who would imagine that).

The problem was - for day 12 - very easy, though. I have a feeling that the whole `Advent of Code 2020` will be way easier than the previous year.  
**Part 2** was not a challenge at all as I had **part 1** working already. It was a very easy transition for me.

My solve-times for this problem are:
- `00:50:42` for part 1 (#5864)
- `01:10:41 ` for part 2 (#4529)
