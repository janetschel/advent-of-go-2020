# Day 2: Password Philosophy
## Problem Summary ([?](https://adventofcode.com/2020/day/2))

This problem requires us to find valid passwords in a list and count them up.

Suppose following list of passwords:
```
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
```

Each line of the input gives a password policy and a password, which we need to validate.

For the first line `1-3 a: abcde` our policy is as follows:
- For **part 1** `1-3` indicates that the character must be contained in the password at least `1` and at most `3` times
- `a` is the character
- `abcde` is the password

Following these criterias, ex. 1 and 3 are valid.  
The middle password **is not** since `b` is containt 0 times in it.

For **part 2** all the above rules still apply, but `1-3` does not correspond to the number of times the char must appear in the password,
it corresponds to not zero-indexed indices, at which the password must contain the char exactly 1 time.

`1-3 a` means that the password has to contain the char `a` at either position `1` or position `3`, not both (XOR)!

For **part 1** my solution is `447`.  
For **part 1** my solution is `249`.

## Recap
Pretty easy problem, but harder than day 1.  

I hate to say it, but my attempt at solving this was not better than my attempt at solving day 1. :(  

First of all, I had another undiscovered bug in my fetch-utils which I fixed before attempting to solve the problem, which cost me at least 15 minutes.  
After that, I had problems splitting the string on the right indices... I started by splitting on the `:` and not on spaces, which made the problem harder than it needed to be.

After 5 more minutes or so I decided to start splitting my string on the spaces, which then pretty quickly lead me to a solution for part 1 and part 2.

My solve-times were:
- 00:29:14 for part 1 (#5610)
- 00:36:35 for part 2 (#5204)

I think all the bugs should be fixed by now and tomorrow will be my day.
