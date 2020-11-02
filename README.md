# advent-of-go-2020
Advent of Code 2020 in golang  

These are by no means good or fast solutions, just my personal repo to keep tracks of this years progress  
Please no Intcode.....

## Usage for fetch utils
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
│   ├── day-01
│   └── ...
├── utils/
│   ├── fetch.go
│   └── readfile.go
└── secrets/
    └── session.go
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
	input := utils.ReadInputFile(1, "\n")
	//			     ^    ^
	//			     |    |
	//                          day   |
	//                                |
	//                            delimiter
}
```
