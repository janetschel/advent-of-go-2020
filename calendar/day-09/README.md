# Day 9: Encoding Error
## Problem Summary ([?](https://adventofcode.com/2020/day/9))

For this problem we need to break some rudimentary encryption... it really does seem like [we're the bad guys](https://www.reddit.com/r/adventofcode/comments/k9481t/2020_day_8_are_we_the_bad_guys/).

For this problem, *preamble* amount of numbers (in our case `25`) are transmitted at the start.  
After that, each number you receive should be the sum of any two of the `25` immediately previous numbers. The two numbers will have different values, and there might be more than one such pair.

Suppose following input list with a *preamble* of `5`:
```
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
```

We start by transmitting the first *preamble* (in this case `5`) amount of numbers to begin with.  
After that, each new number must be a sum of any two pairs of the `5` numbers before. The next number must be any pair of the previous 5 numbers, etc.

In the above example, after the `5-number preamble`, almost every number is the sum of two of the previous 5 numbers; the only number that does not follow this rule is `127`.

For **part 1** we need to find the **first** number, with a `preamble of 25` which is not the sum of a pair of numbers *preamble*-steps before.

For **part 2** we first need to find a contiguous set of numbers, which when added together, produce the solution of **part 1**. The set can be as large as needed.  
For the answer to **part 2** we must add together the `min and max value` of this range we just found.

My solution for **part 1** is `776203571`.  
My solution for **part 2** is `104800569`.

## Recap
I am starting to see a [pattern](https://i.redd.it/hfqowwoz3q361.png)...

Either way, pretty easy day today and very straight-forward problem.  
Really nothing to say except I had fun solving this problem.

**Part 1** was very straight-forward.  

While reading **part 2** I missed the key fact, that I need to add together the `min and max` value of the contiguous set of numbers.  
I just thought I had to add up the first and last number in this set. This cost me about 6 minutes or so to find out (after re-reading the problem).  
Not as deadly as searching for the error for an hour, but my solve-time still suffered a little.

My solve-times for this problem are:
- `00:14:51` for part 1 (#3735)
- `00:24:39` for part 2 (#3164)
