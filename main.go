package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		startRepl()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	}

}

func startRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runFile(file string) {
	if !fileExists(file) {
		fmt.Println("File not found")
		return
	}
	if !checkExt(file) {
		fmt.Println("File type not supported")
		return
	}
	dat, err := ioutil.ReadFile(file)
	check(err)

	env := object.NewEnviroment()
	l := lexer.New(string(dat))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
		return
	}

	evaluator.Eval(program, env)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func checkExt(filename string) bool {
	return filepath.Ext(filename) == ".ape"
}

const monkeyFace = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, monkeyFace)
	io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
