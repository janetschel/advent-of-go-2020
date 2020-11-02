# advent-of-go-2020
Advent of Code 2020 in golang

## Usage for fetch utils
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
