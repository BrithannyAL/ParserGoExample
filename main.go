package main

import (
	"GoExample/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream( /*os.Args[1]*/ "test.txt")
	lexer := parser.NewgoScanner(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgoParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.Program()
	//tree := p.Program()
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
