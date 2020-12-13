# Day 13: Shuttle Search
## Problem Summary ([?](https://adventofcode.com/2020/day/13))

For this problem, we have a bus plan a have to find connections to the airport.  
For example, lets have a look at some bus connections:  
```
939
7,13,x,x,59,x,31,19
```

`939` is the earliest connection I can take... we have to wait to **at least** this number, and then take the next bus.  
The second line lists the bus IDs that are in service according to the shuttle company; entries that show x must be out of service, so you decide to ignore them.

The earliest bus we could take is bus `ID 59`. It doesn't depart until timestamp `944`, so we would need to wait `944 - 939 = 5` minutes before it departs.  
Multiplying the bus ID by the number of minutes you'd need to wait gives `295`.

For **part 1** we need to find the earliest bus we can take, calculate the time we need to wait for it, and multiply this number with the bus ID.

For **part 2** things start to get a little more complicated:  
We must now find a timestampe `t`, so that all subsequent busses depart at this timestamp plus the index they are on.

Our problem is now as follows:  
Find the earliest timestamp such that the first bus ID departs at that time and each subsequent listed bus ID departs at that subsequent minute.

My solution for **part 1** is `4938`.  
My solution for **part 2** is `230903629977901`.  

# Recap

To start the day of, not the easiest **part 1** for me.  
I actually found it, for whatever reason, quite hard to wrap my head around this problem, which is quite strange indeed, since - in retrospective - this wasn't a hard **part 1** at all.  
I don't know why it was, that **part 1** struck me as this difficult at the time. It was definitely not the parsing...

**Part 2** - on the other hand - was **then** quite "easy" for me.  
I had the right idea right away and chased that path down to success. It took me some time - sure - but apparently I was a little faster than the majority of other people (you can see that in my rankings for **part 2**... they make a huge jump from part 1).  
I think the essence in **part 2** (and maybe my luck) was just to find the right idea in a good time.

There is actually a good mathmetical theorem that *apparently* does the trick. It is called the [Chinese remainder theorem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem).  
I'm actually not too sure if I used that exact theorom in my solution, and I'm honestly not checking, but my solution works just as well :)

I had fun solving this problem and getting better in **part 2** (from **part 1**).  
I'm still hoping for some `Intcode v2.0` action.

My solve-times for this problem are:
- `00:14:27` for part 1 (#3112) 
- `00:40:00` for part 2 (#799) ⭐️ (very good **part 2**)
