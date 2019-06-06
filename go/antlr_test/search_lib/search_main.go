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
	fmt.Printf("Enter EveryRule. text[%s]\n", ctx.GetText())
}

func (this *SearchQueryListener) EnterQuery(ctx *query_parser.QueryContext) {
	fmt.Printf("Enter Query. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterAndQueryCriterias(ctx *query_parser.AndQueryCriteriasContext) {
	fmt.Printf("Enter AndQueryCriterias. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterNotQueryCriterias(ctx *query_parser.NotQueryCriteriasContext) {
	fmt.Printf("Enter NotQueryCriterias. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterSingleQueryCriteria(ctx *query_parser.SingleQueryCriteriaContext) {
	fmt.Printf("Enter SingleQueryCriteria. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterBracketQueryCriterias(ctx *query_parser.BracketQueryCriteriasContext) {
	fmt.Printf("Enter BracketQueryCriterias. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterOrQueryCriterias(ctx *query_parser.OrQueryCriteriasContext) {
	fmt.Printf("Enter OrQueryCriterias. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterQueryCriteria(ctx *query_parser.QueryCriteriaContext) {
	fmt.Printf("Enter QueryCriteria. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterKeyCriteria(ctx *query_parser.KeyCriteriaContext) {
	fmt.Printf("Enter KeyCriteria. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterHeaderCriteria(ctx *query_parser.HeaderCriteriaContext) {
	fmt.Printf("Enter HeaderCriteria. text[%s]\n", ctx.GetText())
}

func (s *SearchQueryListener) EnterValueCriteria(ctx *query_parser.ValueCriteriaContext) {
	fmt.Printf("Enter ValueCriteria. text[%s]\n", ctx.GetText())
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

func (v *SearchQueryVisitor) visitChild(c antlr.Tree) interface{} {
	fmt.Printf("Visit child. type[%v]\n", reflect.TypeOf(c))
	switch c.(type){
	case *query_parser.AndQueryCriteriasContext:
		v.VisitAndQueryCriterias(c.(*query_parser.AndQueryCriteriasContext))
	case *query_parser.NotQueryCriteriasContext:
		v.VisitOrQueryCriterias(c.(*query_parser.OrQueryCriteriasContext))
	case *query_parser.SingleQueryCriteriaContext:
		v.VisitSingleQueryCriteria(c.(*query_parser.SingleQueryCriteriaContext))
	case *query_parser.BracketQueryCriteriasContext:
		v.VisitBracketQueryCriterias(c.(*query_parser.BracketQueryCriteriasContext))
	case *query_parser.OrQueryCriteriasContext:
		v.VisitOrQueryCriterias(c.(*query_parser.OrQueryCriteriasContext))
	case *query_parser.QueryCriteriaContext:
		v.VisitQueryCriteria(c.(*query_parser.QueryCriteriaContext))
	case *query_parser.KeyCriteriaContext:
		v.VisitKeyCriteria(c.(*query_parser.KeyCriteriaContext))
	case *query_parser.HeaderCriteriaContext:
		v.VisitHeaderCriteria(c.(*query_parser.HeaderCriteriaContext))
	case *query_parser.ValueCriteriaContext:
		v.VisitValueCriteria(c.(*query_parser.ValueCriteriaContext))
	default:
		fmt.Printf("Not support handle child. type[%v]\n", reflect.TypeOf(c))
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitQuery(ctx *query_parser.QueryContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit Query. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitAndQueryCriterias(ctx *query_parser.AndQueryCriteriasContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit AndQueryCriterias. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitNotQueryCriterias(ctx *query_parser.NotQueryCriteriasContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit NotQueryCriterias. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitSingleQueryCriteria(ctx *query_parser.SingleQueryCriteriaContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit SingleQueryCriteria. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitBracketQueryCriterias(ctx *query_parser.BracketQueryCriteriasContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit BracketQueryCriteria. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitOrQueryCriterias(ctx *query_parser.OrQueryCriteriasContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit OrQueryCriterias. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitQueryCriteria(ctx *query_parser.QueryCriteriaContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit QueryCriteria. count[%d] text[%s]\n", n, ctx.GetText())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitKeyCriteria(ctx *query_parser.KeyCriteriaContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit KeyCriteria. count[%d] text[%s]\n", n, ctx.GetText())
	fmt.Printf("KeyCriteria value. key[%v]", ctx.KEY())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitHeaderCriteria(ctx *query_parser.HeaderCriteriaContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit HeaderCriteria. count[%d] text[%s]\n", n, ctx.GetText())
	fmt.Printf("HeaderCriteria value. header[%v]", ctx.HEADER())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

func (v *SearchQueryVisitor) VisitValueCriteria(ctx *query_parser.ValueCriteriaContext) interface{} {
	n := ctx.GetChildCount()
	fmt.Printf("Start visit ValueCriteria. count[%d] text[%s]\n", n, ctx.GetText())
	fmt.Printf("ValueCriteria value. key[%v] quote[%v] noquote[%v]\n", ctx.KEY(), ctx.QUOTE_STRING(), ctx.NOQUOTE_STRING())
	for i := 0; i < n; i++ {
		c := ctx.GetChild(i)
		v.visitChild(c)
	}
	return map[string]string{}
}

// var t = "subject:test after:2019-01-02 and test or -from:jinlian and (from:ping or to:ping) $hello subject:xx"
// var t = "9_subject:test and after:t2019-01-02 and test test2"
// var t = "test:item key:\" value \" and key2:value2 test test"
// var t = "test"
var t = "key:test"
// var fn = "test"
// var fn = os.Args[1]

func testQueryParserListener() {
	fmt.Printf("\ntestQueryParserListener: %s\n", t)
	input := antlr.NewInputStream(t)
	// input, _ := antlr.NewFileStream(fn)
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := query_parser.NewQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	antlr.ParseTreeWalkerDefault.Walk(NewSearchQueryListener(), tree)
}

func testQueryParserIterate() {
	fmt.Printf("\ntestQueryParserIterate: %s\n", t)
	// input, _ := antlr.NewFileStream(fn)
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
	fmt.Printf("\ntestQueryParserVisitor: %s\n", t)
	// input, _ := antlr.NewFileStream(fn)
	input := antlr.NewInputStream(t)
	lexer := query_parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := query_parser.NewQueryParser(stream)
	p.BuildParseTrees = true
	visitor := NewSearchQueryVisitor()
	p.Query().Accept(visitor)

	/*
	// visitor := SearchQueryVisitor{}
	switch t := interface{}(visitor).(type) {
	case query_parser.QueryVisitor:
		fmt.Printf("query type %v", t)
	default:
		fmt.Printf("default type %v", t)
	}
	r := p.Query().Accept(visitor)
		qc := p.Query()
		r := visitor.VisitQuery(qc.(*query_parser.QueryContext))
	fmt.Printf("visitor: %v\n", r)
	*/
}

func main() {
	testQueryParserIterate()
	// testQueryParserListener()
	// testQueryParserVistor()
}
