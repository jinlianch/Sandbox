package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"Sandbox/antlr_test/json_parser"
	"os"
	"fmt"
)

type TreeShapeListener struct {
	*json_parser.BaseJSONListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
	fmt.Println("xxxx")
}

func testJsonParser() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := json_parser.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := json_parser.NewJSONParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func main() {
	testJsonParser()
}