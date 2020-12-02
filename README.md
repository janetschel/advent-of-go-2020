# advent-of-go-2020
Advent of Code 2020 in golang  

These are by no means good or fast solutions, just my personal repo to keep tracks of this years progress  
Please no Intcode..... *(doesn't look like it yet :))))))))*

## Makefile
Use `make new` in the terminal to automatically create a new folder with the important files for the day.  
These files are copied from the static-template folder. If you want to make any changes to the blueprint, do it there.

## Usage for fetch utils
Keep in mind that these fetch utils **do not** request the input **again** if the input file already exists.  
It will just use the existing file without making any HTTP requests, so we don't spam the API of https://adventofcode.com/


If you **want to make** another request to AOC in order to refresh your input file (for whatever reason you want to do so), you need to delete the already existing `puzzle-input.in` file so the fetch-utils make a new request.


### Setup
In order for the fetch utils to work properly:
- create a folder /secrets in the src/advent-of-go-2020
- place session.go inside this newly created folder
- place your session-key inside session.go. You can find your session key in your session cookie when visiting the website

<br />
Your project structure should look likes this:  

```
advent-of-go-2020/
├── calendar/
│   ├── day-01/
│   │   ├── day01.go
│   │   ├── day01_pt02.go
│   │   └── puzzle-input.in  // <-- this is created by fetch utils
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
<br />

`session.go` file:
```golang
package secrets

const (
	Session = "session=<your-session-key>"
)
```


### Usage

```golang
package main

import "advent-of-go-2020/utils"

func main() {
	input := utils.ReadFile(1, "\n")
	//		        ^    ^
	//	       	        |    |
	//                     day   |
	//                           |
	//                       delimiter
	
	// more code...
}
```
<br />

## []string to []int function
This blueprint also offers a convenient way to convert string slices to int slices on the fly.

Use it as follows:
```golang
package main

import "advent-of-go-2020/utils"

func main() {
	// explicit type declaration to see what's going on
	
	var inputSlice []string = utils.ReadFile(1, "\n")
	var inputConverted []int = utils.ToIntSlice(input)
}
```

**It offers way more util functions now... just check it out :)**
