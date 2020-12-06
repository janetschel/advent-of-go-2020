# Day 6: Custom Customs
## Problem Summary ([?](https://adventofcode.com/2020/day/6))

This problem asks us to find valid answers in a custom declaration form.  
The form contains 26 "yes-or-no" questions, that anyone in a group can answer with "yes" (or "no").

Suppose following answers:
```
abcx
abcy
abcz
```

In this group, there are 6 questions to which anyone answered "yes":  
`a`, `b`, `c`, `x`, `y`, and `z`. (Duplicate answers to the same question don't count extra)

Another set of examples is:
```
abc

a
b
c

ab
ac

a
a
a
a

b
```

The way to interpret these answers is as follows:
- The first group contains one person who answered "yes" to 3 questions: `a`, `b`, and `c`.
- The second group contains three people; combined, they answered "yes" to 3 questions: `a`, `b`, and `c`.
- The third group contains two people; combined, they answered "yes" to 3 questions: `a`, `b`, and `c`.
- The fourth group contains four people; combined, they answered "yes" to only 1 question, `a`.
- The last group contains one person who answered "yes" to only 1 question, `b`.

For **part 1** we need - for each group - to count the questions, to which **anyone** answered "yes" and calculate the sum.  
In the previous example, the sum of these counts is `3 + 3 + 3 + 1 + 1 = 11`.

For **part 2** we basically need to do the same thing as in **part 1**, but we need to sum up the questions to which **everyone** said "yes"!

My solution for **part 1** is `6521`.  
My solution for **part 2** is `3305`.

## Recap
Pretty easy problem for day 6. Everything worked out perfectly, and I enjoyed solving the problem.

I had - contrary to the previous days - no problem at understanding the statements in the problem's parts.  
I took some time to read the problems carefully enough to make no mistakes during my attempt to solve it and it worked out! :)  
It seems like I can fulfill my personal goal of grasping the problems better (let's hope I can live up to it the next days).

Nothing else really to add to this recap. Very straight forward solutions.

My solve-times are okay-ish... I guess everyone was having an easy time as well:
- `00:08:39` for part 1 (#3028)
- `00:16:36` for part 2 (#2899)
