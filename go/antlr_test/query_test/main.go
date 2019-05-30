package main

import (
	"Sandbox/go/antlr_test/query_test/query_parser"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	// "os"
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
	fmt.Printf("enter every rule: %v\n", ctx.GetText())
}

func (this *TreeShapeListener) EnterQuery(ctx *query_parser.QueryContext) {
	// a, _ := ctx.ToStringTree()
	// fmt.Printf("ToStringTree %v", a)
	fmt.Printf("enter query: %v\n", ctx.GetText())
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

func (v *TreeQueryVisitor) VisitQuery(ctx *query_parser.QueryContext) interface{} {
	n := ctx.GetChildCount()
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		fmt.Printf("child: %v\n", reflect.TypeOf(c))
	}
	fmt.Println("Start to visit query")
	fmt.Println("Start to visit all pairs")
	pairs := ctx.AllPair()
	for _, p := range pairs {
		r := v.VisitPair(p.(*query_parser.PairContext))
		// r := v.VisitPair(*p)
		fmt.Printf("Visit pair return: %v\n", r)
	}
	fmt.Println("Start to visit all singles")
	singles := ctx.AllSingle()
	for _, s := range singles {
		r := v.VisitSingle(s.(*query_parser.SingleContext))
		fmt.Printf("Visit single return: %v\n", r)
	}
	return map[string]string{}
	// s := v.VisitChildren(ctx)
	// fmt.Printf("visit query: %v", s)
}

func (v *TreeQueryVisitor) VisitPair(ctx *query_parser.PairContext) interface{} {
	fmt.Printf("visit pair %v\n", ctx.GetChildren())
	symbol := ctx.SYMBOL()
	fmt.Printf("pair symbol: %v\n", symbol)
	s := ctx.Single()
	r := v.VisitSingle(s.(*query_parser.SingleContext))
	return r
}

func (v *TreeQueryVisitor) VisitSingle(ctx *query_parser.SingleContext) interface{} {
	fmt.Printf("visit single: %v \n", ctx.GetChildren())
	r := map[string]interface{}{}
	r["symbol"] = ctx.SYMBOL()
	r["string"] = ctx.STRING()
	return r
}

func testQueryParserListener() {
	t := "subject:test after:2019-01-02 and test or -from:jinlian and (from:ping or to:ping) $hello subject:xx"
	input := antlr.NewInputStream(t)
	// input, _ := antlr.NewFileStream(os.Args[1])
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := query_parser.NewQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func testQueryParserIterate() {
	// input, _ := antlr.NewFileStream(os.Args[1])
	t := "subject:test after:2019-01-02 and test or -from:jinlian and (from:ping or to:ping) $hello subject:"
	input := antlr.NewInputStream(t)
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

func testQueryParserVistor() {
	// input, _ := antlr.NewFileStream(os.Args[1])
	// t := "subject:test after:2019-01-02 and test or -from:jinlian and (from:ping or to:ping) $hello subject:"
	t := "subject:test after:2019-01-02 and test or -from:jinlian and (from:ping or to:ping) $hello subject:xx"
	input := antlr.NewInputStream(t)
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
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
	r := p.Query().Accept(visitor)
	/*
		qc := p.Query()
		r := visitor.VisitQuery(qc.(*query_parser.QueryContext))
	*/
	fmt.Printf("visitor: %v", r)
}

func main() {
	// testQueryParserListener()
	// testQueryParserIterate()
	testQueryParserVistor()
}
