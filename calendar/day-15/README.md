# Day 15: Rambunctious Recitation
## Problem Summary ([?](https://adventofcode.com/2020/day/15))

Todays problem consists of a number-game we play with the elves.
In this game, the players take turns saying **numbers**. They begin by taking turns reading from a list of starting numbers, which is our puzzle input for this problem.  
We keep track of the **last number spoken**

The next number spoken after that can be derived from following rules :
- If that was the first time the previous number has been spoken, the current player says 0.
- Otherwise, the current player announces how many turns apart the number is from when it was previously spoken.  
If it has been spoken multiple times, the player announces the difference between the last two times.
  
For **part 1** we need to find the **2020**th number spoken.  
For **part 2** we need to find the **30000000**th number spoken.

My solution for **part 1** is `1085`.  
My solution for **part 2** is `10652`.  

# Recap

**Short** recap action this time:

This problem was not very hard, considering it was stated on the 15th.  
Once you had the hang of what the problem was asking for, it was very easy to solv eit (especially **part 2**).

Nothing really to say about the problem. I solved it fairly quickly and enjoyed it.

For **part 2** however, I just adjusted the for-loop exit condition from `i < 2020` to `i < 30000000`, which of course took a while to compute.  
I'm wondering if there is a faster, more elegant solution to **part 2**.

Besides that, here are my solve-times:  
- `00:39:08` for part 1 (#4318)
- `00:39:43` for part 2 (#2683)
