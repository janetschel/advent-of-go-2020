# Day 14: Docking Data
## Problem Summary ([?](https://adventofcode.com/2020/day/14))

For this problem we need to save some docking data in memory addresses.

Consider following input:
```
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
```

This program starts by specifying a bitmask (`mask = ....`)  
For **part 1** this mask specifies which bits to overwrite in every written value: the `2s bit` is overwritten with `0`, and the `64s` bit is overwritten with `1`.

The program then attempts to write the value `11` to memory address `8`.  
Because of the mask though, the value `73` is written to memory address `8` instead.  
Then in the last iteration `64` is written to address `8`, **overwriting** the value that was there previously.

For **part 1** we need to find the sum of all values that are in our memory after running the program with our input.

For **part 2** the mask no longer changes the number, but the memory address in following way:  
- If the bitmask bit is `0`, the corresponding memory address bit is unchanged.
- If the bitmask bit is `1`, the corresponding memory address bit is overwritten with `1`.
- If the bitmask bit is `X`, the corresponding memory address bit is floating.

**Floating** does - this time - not mean to just ignore it, floating means that this `X` can basically be every value, so we need to look at every permutation of the `X`'s.

For the answer for **part 2** we once again need to find the sum left in memory after we executed the program..

My solution for **part 1** is `11884151942312`.  
My solution for **part 2** is `2625449018811`.

# Recap

In short: **part 1** was okay, **part 2** was horrible.  
But let's start at the beginning.

I actually really liked today. I do notice the problems get harder each day, but I really like the increasing difficulty of mathmatics we need to apply.

**Part 1** was very straight forward for me, I had no problem understanding the problem stated and solved it fairly quickly. Nothing much to say.

On the other hand, I hated **part 2** today. Not only because I didn't immediately catch on to the `"0 does not change anything"` action going on all of a sudden (yup - I missed that), but because of the permutations we needed to calculate for the memory addresses.  
**Go** sadly does not offer any good libraries for generating permutations, so I needed to not only write my own one, but to also debug it... which cost precious time (almost an hour, since I wanted to make it as generic as possible [which is the next bad idea in Go]).

After I had my permutations-generation up and running, the solution was actually also very straight-forward.  
I feel like I wasted my time a little too much there... I really easily get lost in refactoring useless helper-methods instead of tackling the actual problem.  
I think I really need to work on that next. Not wasting my time doing useless stuff. 😉

My solve-times, despite my problems listed, are actually not that bad... I mean they're bad, but not **THAT** bad - I honestly expected worse.  
**Part 2** was of course much worse than **part 1**, I lost quite a few places on the "leaderboard".  
Here they are:
- `00:26:22` for part 1 (#2377)
- `01:37:52` for part 2 (#4037)
