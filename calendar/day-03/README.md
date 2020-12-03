# Day 3: Toboggan Trajectory
## Problem Summary ([?](https://adventofcode.com/2020/day/3))

This problem requires us to find the number trees (represented by #) on a path from a startpoint to an endpoint.

Suppose the following array given:
```
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
```

We start on the top left `(0,0)` and move to the right and down with specified amounts, eg. `(3,1)` (we would move 3 to the right and 1 down).

Consider the array repeating to the right (the lines are not over when they are, they wrap around):
```
..##.........##.........##.........##.......
#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.....
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#
```

Moving `(3,1)` and repeating this until we hit the last line in our list and counting the trees we encounter on the way will result in the solution for this puzzle's **part 1**.

The answer for **part 2** will be calculated as follwing:
- Move following amounts and count the trees encountered seperatly:
    - `(1,1)`
    - `(3,1)` *(answer for part 1)*
    - `(5,1)`
    - `(7,1)`
    - `(1,2)`
- Multiply the solutions together to get the answer.

My solution for **part 1** is `237`.  
My solution for **part 2** is `2106818610`.

## Recap
**FINALLY** no bugs in my fetch-utils, but I sadly woke up 5 minutes late. 
My solve-time isn't that great either this time...

I really liked the problem. I needed to read over it two or three times to grasp the idea, but the solution was very straight forward for my likings.  
Personal goal: get better at grasping problems.

My solve-times for this problem:
- `00:14:00` *(- 5 min ;))* for part 1 (#3491)
- `00:16:51` for part 2 (#2434)
