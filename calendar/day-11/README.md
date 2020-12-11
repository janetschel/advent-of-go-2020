# Day 11: Seating System
## Problem Summary ([?](https://adventofcode.com/2020/day/11))

Todays problem requires us to program our own version of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) to help people seat on an airplane.

Suppose we have following list:
```
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
```

The seat layout fits neatly on a grid. 
Each position is either floor (`.`), an empty seat (`L`), or an occupied seat (`#`).

We operate on every cell **simultaneously** by looking at the **8 adjacents seats** and apllying following rules:
- If a seat is empty (`L`) and there are no occupied seats adjacent to it, the seat becomes occupied.
- If a seat is occupied (`#`) and four or more seats adjacent to it are also occupied, the seat becomes empty.
- Otherwise, the seat's state does not change.

After some number of iterations of this pattern, the seat-occuption does not change anymore and settles down. 

**Part 1** asks us to find the number of occupied seat once the pattern does not change anymore.

**Part 2** changes one rule for us. We do not care to the next adjacent four seat anymores, we care about the next seat we can see on the grid.  
Also, if a seat is occupied, it now needs five or more occupied seats adjacent to it in order to become empty.

The question for **part 2** however stays the same. How many seats are occupied when the pattern stops changing?

My solution for **part 1** is `2334`.  
My solution for **part 2** is `2100`.  

## Recap
As I said in the problem statement, this day required us to basically program [Conway's Game of ](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) ~~Life~~ Seats.

Pretty cool problem, and I really enjoyed solving it.  
I actually had some problems to not index my array out of bounds in part 1, but came up with a solution pretty quick (I [framed](https://github.com/janetschel/advent-of-go-2020/blob/main/utils/slices/slices.go#L53) my array... pretty neat).  
The requirements of **part 2** were very straight-forward for me and easy to implement.

I had no problems of understanding the problem stated and read everything right the first time :)  
I'm really getting better at grasping those types of problems.

Today I was visiting family, which means I wasn't up at exaclty 6am (maybe 6:30 or 6:40am [but I won't use this as an excuse this time around]).  
For that, my solve times were actually alright.

My solve-times for this problem are:
- `01:01:55` for part 1 (#5276)
- `01:13:54` for part 2 (#3680)
