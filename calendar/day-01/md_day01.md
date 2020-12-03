# Day 1: Report Repair
## Problem Summary ([?](https://adventofcode.com/2020/day/1))

For **part 1** we need to find a pair of *two* numbers, which when added together will equal the sum `2020`, and then multiply them together.

For **part 2** we need to find a pair of *three* numbers, which fulfill the requirements already stated.

Suppose the following list:
```
1721
979
366
299
675
1456
```

The two entries in the list which sum two `2020` are `1721` and `299`, so the solution would be `1721 * 299 = 514579`.

The three entries in the list which sum two `2020` are `979`, `366` and `675`, so the solution would be `979 * 366 * 675 = 241861950`.

For **part 1** my solution with the given input is `1019371`.  
For **part 2** my solution with the given input is `278064990`.

## Recap
This problem was - just like the past years - a good warumup problem to begin with.  

Sadly I had a little bug in my fetch utils, which needed to be fixed before tackling the problem.  
This is why my horrendous solve-times for this problem are:
- 00:24:45 for part 1 (#3789)
- 00:26:13 for part 2 (#3326)

I am sure I would've been able to do it in less than 5 minutes... but oh well.

Let's hope for better luck next time
