# Day 4: Passport Processing
## Problem Summary ([?](https://adventofcode.com/2020/day/4))

This problem requires us to find valid passports in a list of given passports.  
A passport *can* have following fields:
- `byr` (Birth Year)
- `iyr` (Issue Year)
- `eyr` (Expiration Year)
- `hgt` (Height)
- `hcl` (Hair Color)
- `ecl` (Eye Color)
- `pid` (Passport ID)
- `cid` (Country ID)

Suppose following passport list: (passports are seperated by an empty newline from another):
```
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
```

For both parts we need to find the amount of valid passports in our input.

For **part 1** a passport is valid if it has all the fields from above (except the field `cid`).  
If the field `cid` is present **or missing** we treat the passport as valid

The first passport in our example list is valid since all eight fields are present. 
The second passport is invalid because it is missing `hgt`.

For **part 2** all ob the above rules still apply, however each field of the passport has its own criterias that need to be fulfilled in order for the whole passport to be valid.  
These criterias are as follows:
- `byr`: four digits; at least 1920 and at most 2002.
- `iyr`: four digits; at least 2010 and at most 2020.
- `eyr`: four digits; at least 2020 and at most 2030.
- `hgt`: a number followed by either cm or in:
    - If cm, the number must be at least 150 and at most 193.
    - If in, the number must be at least 59 and at most 76.
- `hcl`: a # followed by exactly six characters 0-9 or a-f.
- `ecl`: exactly one of: amb, blu, brn, gry, grn, hzl or oth.
- `pid`: a nine-digit number, including leading zeroes.

If all of these criterias apply to the fields in the passport, the passport is valid.

My solution for **part 1** is `182`.  
My solution for **part 2** is `109`.

## Recap
Well.. I have very mixed opinions on this problem... Not because it is a rather strange problem which mostly involves string-manipulation, but because of my apparent ability to read the problem stated.

**Part 1** was very easy and straig-forward for me... just some `String.Contains(...)` calls and that's it.  
**Part 2** actually does not differ from this all that much, but I missed reading a rule in the list and did not check it... it took me a long while to re-read the problem and to find out, that I was missing an if-check (about 50 minutes, sadly). I was desperatly searching for a bug in my `solvePart2()` function, sadly.

My solve-time for **part 1** was very good, but since I was unable to read the second part carefully enough, the solve-time there did not keep up with the first one.  
I did not have any bugs in my fetch-utils and everything (besides **part 2**) worked out perfectly. 

I notice I'm getting better at solving such a kind of problems, **IF** I get them ;)

As stated in an earlier recap: my personal goal is to grasp problems better and understand them... and crucially not missing key components in the problem statements.

My solve-times for this problem:  
(I'm almost positive my solve-time for part 2 could've been around 30 minutes)
- `00:05:43` for part 1 (#379) ⭐️
- `01:09:45` for part 2 (#5072)

I hope for better reading comprehension tomorrow ;)