### Git workflow

1. make the necessary changes to the codebase
2. stage the changes: `git add .`
3. commit the staged changes: `git commit -m "<your message>"`
4. push current branch to remote: `git push -u origin main` (first time, then: `git push`)

to get the changes from the remote: `git pull`
### About go language

to run go program: `go run <path to main.go file>`. For the current program: `go run src/main.go`

### Go concepts

**value**: specific unit of data. example: `"Max"`, `54`, `true`

**simple value type**: category of data. example: `int`, `bool`, `[5]int`, `*int`

**complex value type**: category of data. example: `string`, `any`, `[]int`

**variable**: named storage for a value. syntax: `name := "Alex"`

**function**: reusable block of code to perform a task. syntax: `func Name() {}`

**main function**: entry point where program execution begins. syntax: `func main() {}`

**package**: folder containing go files to organize code. syntax: `package name`

