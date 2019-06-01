package main

import (
	"Sandbox/go/antlr_test/search_lib/query_parser"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	// "os"
	"fmt"
)


type SearchQueryListener struct {
	*query_parser.BaseQueryListener
}

func NewSearchQueryListener() *SearchQueryListener {
	return new(SearchQueryListener)
}

func (this *SearchQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// a, _ := ctx.ToStringTree()
	// fmt.Printf("ToStringTree %v", a)
	fmt.Printf("enter every rule: %v\n", ctx.GetText())
}

func (this *SearchQueryListener) EnterQuery(ctx *query_parser.QueryContext) {
	// a, _ := ctx.ToStringTree()
	// fmt.Printf("ToStringTree %v", a)
	fmt.Printf("enter query: %v\n", ctx.GetText())
}

type SearchQueryVisitor struct {
	*query_parser.BaseQueryVisitor
}

func NewSearchQueryVisitor() *SearchQueryVisitor {
	return &SearchQueryVisitor{
		&query_parser.BaseQueryVisitor{},
	}
	// return new(SearchQueryVisitor)
}

func (v *SearchQueryVisitor) VisitQuery(ctx *query_parser.QueryContext) interface{} {
	n := ctx.GetChildCount()
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		fmt.Printf("child: %v\n", reflect.TypeOf(c))
	}
	/*
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
	*/
	return map[string]string{}
	// s := v.VisitChildren(ctx)
	// fmt.Printf("visit query: %v", s)
}


var t = "subject:test after:2019-01-02 and test or -from:jinlian and (from:ping or to:ping) $hello subject:xx"
var fn = "test"

func testQueryParserListener() {
	input := antlr.NewInputStream(t)
	// input, _ := antlr.NewFileStream(os.Args[1])
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := query_parser.NewQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	antlr.ParseTreeWalkerDefault.Walk(NewSearchQueryListener(), tree)
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
	visitor := NewSearchQueryVisitor()

	// visitor := SearchQueryVisitor{}
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
	testQueryParserListener()
	// testQueryParserIterate()
	// testQueryParserVistor()
}
