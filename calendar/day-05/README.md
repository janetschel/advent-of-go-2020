# Day 5: Binary Boarding
## Problem Summary ([?](https://adventofcode.com/2020/day/5))

In this problem we need to seat people on an airplane using binary search patterns.

More specifically, a seat can is specified by a row of instructions on how to search for the right seat.  
For example, a seat-patter could be `FBFBBFFRLR`, where `F` means front, `B` means back, `L` means left, and `R` means right.

The first 7 characters will either be `F` or `B`, indexing our row. The last three characters index the column of the seat.

This pattern `FBFBBFFRLR` can be interpreted in following way:
- We start by consider the whole row range from `0` to `127`
- `F` means to take the lower half, keeping rows `0` through `63`.
- `B` means to take the upper half, keeping rows `32` through `63`.
- and so on...
- The final `F` keeps the lower of the two, `row 44`.

The same goes for the colum, except the range changes to `0` - `7`.  
`R` means we keep the upper half, for `L` we keep the lower one.

Once we found the row and column our seat is in, we can calculate the seat ID of said seat.  
Calculating the seat ID for a known row and column is rather easy: multiply the row by 8, then add the column.  
In our example the seat ID is: `44 * 8 + 5 = 357`

For **part 1** we have to find the highest seat ID in our input.  

For **part 2** we have to find the seat ID, which is not in the input (it is only one), so we can have a seat as well.  
Apparently, our seat is not at the very front, or the very back of the plane. We have to keep this in mind when searching for the missing seat ID.

My solution for the given boarding passes for **part 1** is `908`.  
My solution for **part 2** `619`.

## Recap
All in all, this was an okay problem.  
It honestly just consisted of some basic binary-searching and splitting ranges.

My inability to read the problem carefully enough persisted a little.  
I first thought the range for the column is `0` to `8` and not `0` to `7`, which lead me to some wrong answers (but also to some correct ones). This was the reason, why it took me so long to debug :(  
After 20 minutes of trying I decided to re-read the problem and found my error (finally).

Also, I got a little cought up in the second part, since I did not immediatly understand what the problem was asking for. I needed to re-read this several times.

Besides that, I had no problems at all with this problem and really enjoyed it.  
It also seems that the problems don't get as hard as quickly as they did in 2019.

My solve-times aren't (because of the stated problems I had with this problem) as good also.  
My solve-times are:
- `00:32:20` for part 1 (#5779)
- `00:56:26` for part 2 (#6807)



