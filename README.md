# advent-of-go

Advent of Code in golang  

A fork with some additional templating for multiple years. These solutions are not guaranteed to be performant or quality.

## Makefile

Use `make new` in the terminal to automatically create a new folder with the important files for the day.  
To generate a specific day or year, use the `make new day=1 year=2020` command, otherwise the current date and year will be used.
To generate an entire year of solution boilerplate, use `make all year=2021`.
These files are copied from the static-template folder. If you want to make any changes to the blueprint, do it there.

## Usage for fetch utils

Keep in mind that these fetch utils **do not** request the input **again** if the input file already exists.  
It will just use the existing file without making any HTTP requests, so we don't spam the API of https://adventofcode.com/

If you **want to make** another request to AOC in order to refresh your input file (for whatever reason you want to do so), you need to delete the already existing `puzzle-input.in` file so the fetch-utils make a new request.

### Setup

In order for the fetch utils to work properly:

- Set environment variable `ADVENT_OF_CODE_SESSION_TOKEN` to an Advent of Code session token. You can find your token in your session cookie when visiting the website
- Or, if using dev containers, create or modify the `.devcontainers/devcontainer.env` file to include this variable:

```
ADVENT_OF_CODE_SESSION_TOKEN=<your session token>
```

> :red_circle: If migrating from a previous version with the session key in `session.go`, make sure to remove the key before pushing to prevent committing the session key. :red_circle: 

Your project structure should look likes this:  

```shell
advent-of-go/
├── calendar/
│   ├── 2020/
│   │   ├── day-01/
│   │   │   ├── day01.go
│   │   │   ├── day01_pt02.go
│   │   │   └── puzzle-input.in  // <-- this is created by fetch utils
│   │   └── ...
│   └── ...
├── utils/
│   ├── conv/
│   ├── files/
│   ├── req/
│   ├── slices/
│   └── str/
├── secrets/
│   └── session.go
├── Makefile 
└── template    // <-- change this if you wish to modify your blueprint
```

### Usage

```golang
package main

import "advent-of-go/utils"

func main() {
  input := utils.ReadFile(1, 2020, "\n")
  //                      ^    ^    ^
  //                      |    |    |
  //                     day   |    |
  //                           year |
  //                               delimiter
  
  // more code...
}
```

## Running Solutions

Unless commented out, both Part 1 and Part 2 will be run for a given day. To run a specific day, run the terminal command `go run calendar/<year>/day-<day>/<day>.go`. If properly configured and available, the input file will be fetched automatically when the solution is run.
