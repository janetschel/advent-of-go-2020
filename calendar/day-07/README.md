# Day 7: Handy Haversacks
## Problem Summary ([?](https://adventofcode.com/2020/day/7))

This problem is pretty easily summed up in luggage counting.

Take following list:
```
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
```

This list specifies rules of 9 different type of bags, and which bags they contain.

We - according to the problem - have a **shiny golden** bag.

Considering the above rules, here is an example-case:

- A `bright white` bag, which can hold your `shiny gold` bag **directly**.
- A `muted yellow` bag, which can hold your `shiny gold` bag **directly**, plus some other bags.
- A `dark orange` bag, which can hold `bright white` and `muted yellow` bags, either of which could then hold your `shiny gold` bag.
- A `light red` bag, which can hold `bright white` and muted `yellow bags`, either of which could then hold your `shiny gold` bag.

So, in this example, the number of bag colors that can eventually contain **at least one** `shiny gold` bag is `4`.

For **part 1** we need to calculate how many colored bags can at least contain one shiny golden one.  
The important part to not miss is **at least**.

For **part 2** we need to calculate how many total other bags are in the golden shiny one.  
All the above rules still apply.

My solution for **part 1** is `235`.  
My solution for **part 2** is `158493`.

## Recap
The whole problem was a mess for me... I wasn't to forget it asap. 🤐

I started looking at the input trying to figure out how in the hell I want to parse it and consider every edge-case there might be.  
Keep in mind, I'm - at this point - just glancing over the input, so I don't actually know how many edge-cases there are.

I - for what reason in the whole world ever - decided to do this input parsing using a custom regex...  
As you can probably guess, writing up my regex took me 75 minutes. I didn't want to throw the idea of regex-parsing overboard (what I probably should've after 10 minutes of trying), so my solve times are once again horrendous.

Want to see my regex? Of course, here is the beauty:
```regexp
^((?:\w+)? \\w+) bags contain (?:(no other bags.)|((?:(?:(?:\d+) (?:(?:\w+)? \w+) (?:bags|bag), )|(?:(?:\d+) (?:(?:\w+)? \w+) (?:bags|bag)\.))*))$
```

And here is another, because I couldn't manage to fit it in one... *(damn you, capture groups!)*

```regexp
^(\d+) ((?:\w)*(?: |)\w+) (?:bags|bag)$
```

My first regex captures all the possible variaties of containing bags, the second regex parses them.  
I know these regular expressions are not really performant and could probably be shortened a lot, but I won't ever touch those again.

My solve-times for part 1 were subsequently (because of the regex) not very good. I - however - manage to do the second part after that really quickly, since I had a good datastructure already set up.  
My solve-times for this problem are:
- `02:20:22` for part 1 (#9720)
- `02:29:29` for part 2 (#7547)

I hope for no regex parsing tomorrow (or at least for a better approach of me)...  
Eventhough my solve-times were pretty bad, I placed kinda better than I expected from 2 hours of regex-parsing. Maybe this was a harder problem for everyone.