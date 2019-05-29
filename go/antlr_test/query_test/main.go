package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"Sandbox/go/antlr_test/query_test/query_parser"
	"os"
	"fmt"
)

type TreeShapeListener struct {
	*query_parser.BaseQueryListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// a, _ := ctx.ToStringTree()
	// fmt.Printf("ToStringTree %v", a)
	fmt.Printf("enter: %v\n", ctx.GetText())
}


type TreeQueryVisitor struct {
	*query_parser.BaseQueryVisitor
}

func NewTreeQueryVisitor() *TreeQueryVisitor {
	return &TreeQueryVisitor{
		&query_parser.BaseQueryVisitor{},
	}
	// return new(TreeQueryVisitor)
}

func (v *TreeQueryVisitor) VisitQuery(ctx *query_parser.QueryContext) interface {} {
	fmt.Println("Start to visit query")
	pairs := ctx.AllPair()
	for _, p := range pairs {
		r := v.VisitPair(p.(*query_parser.PairContext))
		// r := v.VisitPair(*p)
		fmt.Printf("Visit pair: %v\n", r)
	}
	singles := ctx.AllSingle()
	for _, s := range singles {
		r := v.VisitSingle(s.(*query_parser.SingleContext))
		fmt.Printf("Visit single: %v\n", r)
	}
	return map[string]string {}
	// s := v.VisitChildren(ctx)
	// fmt.Printf("visit query: %v", s)
}

func (v *TreeQueryVisitor) VisitPair(ctx *query_parser.PairContext) interface{} {
	// fmt.Println("visit pair")
	s := ctx.Single()
	r := v.VisitSingle(s.(*query_parser.SingleContext))
	return r
}

func (v *TreeQueryVisitor) VisitSingle(ctx *query_parser.SingleContext) interface{} {
	// fmt.Println("visit single")
	r := map[string]interface {} {}
	r["symbol"] = ctx.SYMBOL()
	r["string"] = ctx.STRING()
	return r
}

func testQueryParserListener() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := query_parser.NewQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func testQueryParserIterate(){
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := query_parser.NewQueryLexer(input)
	/*
	tokens := lexer.GetAllTokens()
	for _, t := range tokens {
		fmt.Printf("token: %v\n", t)
	}
	*/
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func testQueryParserVistor(){
	// input, _ := antlr.NewFileStream(os.Args[1])
	input, _ := antlr.NewFileStream("test")
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := query_parser.NewQueryParser(stream)
	p.BuildParseTrees = true
	visitor := NewTreeQueryVisitor()
	// visitor := TreeQueryVisitor{}
	/*
	switch t := interface{}(visitor).(type) {
	case query_parser.QueryVisitor:
		fmt.Printf("query type %v", t)
	default:
		fmt.Printf("default type %v", t)
	}
	*/
	// visitor := new(query_parser.BaseQueryVisitor)
	// r := p.Query().Accept(visitor.(query_parser.QueryVisitor))
	qc := p.Query()
	r := visitor.VisitQuery(qc.(*query_parser.QueryContext))
	fmt.Printf("visitor: %v", r)
}

func main() {
	// testQueryParserListener()
	// testQueryParserIterate()
	testQueryParserVistor()
}
