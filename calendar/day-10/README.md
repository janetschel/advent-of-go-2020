# Day 10: Adapter Array
## Problem Summary ([?](https://adventofcode.com/2020/day/10))

In this problem, we discover that a battery has no power anymore, and we need to plug it into an outlet.

There's only one problem: the charging outlet near your seat produces the wrong number of **jolts**.

Each of your joltage adapters is rated for a specific output **joltage**.  
Any given adapter can take an input `1`, `2`, or `3` jolts **lower** than its rating and still produce its rated output joltage.  
Additionally, our device has a built-in joltage adapter rated for `3` jolts **higher** than the highest-rated adapter in our bag.

Suppose following joltage ratings:
```
16
10
15
5
1
11
7
19
6
12
4
```

For this list, our device's built-in joltage adapter would be rated for `19 + 3 = 22`.

We can interpret the given list in following ways:
- The charging outlet has an effective rating of `0` jolts, so the only adapters that could connect to it directly would need to have a joltage rating of 1, 2, or 3 jolts. Of these, only one we have is an adapter rated `1 jolt` (difference of 1).
- From the `1-jolt` rated adapter, the only choice is your `4-jolt` rated adapter (difference of 3).
- From the `4-jolt` rated adapter, the adapters rated `5`, `6`, or `7` are valid choices. However, in order to not skip any adapters, we have to pick the adapter rated 5 jolts .
- Similarly, the next choices would need to be the adapter rated `6` and then the adapter rated `7`.
- *And so on...*
- Finally, our device's built-in adapter is always `3 higher` than the highest adapter, so its rating is 22 jolts (**always** a difference of 3).

In this example, when using every adapter, there are `7` **differences of 1 jolt** and `5` **differences of 3 jolts**.

---
For **part 1** we have to multiply the differences of 1 jolt with the differences of 3 jolts.

For **part 2** we have to find the total number of distrinct ways we can arrange the adapters to connect the charging outlet to your device.

My solution for **part 1** is `2368`.  
My solution for **part 2** is `1727094849536`.  

## Recap
All in all a very **chaotic** but productive day for me...  
It all started with waking up 10 minutes late, at `6:30am`, when the puzzle was already unlocked for **30 MINUTES**.

After that rough start in the day, **part 1** was actually very easy to solve. Not really to say here, except that I maybe expected an `Intcode v2.0` problem.

**Part 2** actually did not seem - at the first glance - that hard either, but I didn't cache my results the first time... my execution time was horrible.  
After reading about [Memoization](https://en.wikipedia.org/wiki/Memoization) (which I shouldn't have to read up in the first place) it clicked for me, and I implemented that strategy, which turned out to be a real lifesaver.

As you have probably guessed, the execution time of my program was *waaay* faster (an average of `+5.000000e-005` seconds [or `0.05ms`] for 1,000,000 exections), which helped me get my answer very quick.  
Don't get me wrong, my first implementation of **part 2** already worked fine-ish, but I had to wait a solid minute or two for it to calculate a result, which actually is no wonder when you see the input.

I loved solving this problem and learning something *kinda* new (the first time in `AoC`).  
I was already familiar with the concept of caching results of expensive function calls, but I never really implemented it on some real problems (except the Fibonacci numbers). 

With all circumstances put into consideration, I am actually very happy with my solve-times this time.  
My solve-times for this problem are:
- `00:51:55` (*-30 minutes* of sleeping ;)) for part 1 (#9105)
- `01:25:01` (an additional ~33 minutes) for part 2 (#4521)

Let's hope I wake up on time tomorrow.  
>! And let's hope for some Intcode v2.0
