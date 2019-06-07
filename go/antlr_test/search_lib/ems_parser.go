package main

import (
	"reflect"
	"log"
	"strings"
	"Sandbox/go/antlr_test/search_lib/search_parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SearchOpType int32
type SearchCriteriaType int32

const (
	OPERATOR_NONE SearchOpType = iota
	OPERATOR_AND
	OPERATOR_OR
	OPERATOR_NOT
)

const (
	CRITERIA_FIELD SearchCriteriaType = iota
	CRITERIA_HEADER
	CRITERIA_VALUE
)

const (
	NonLeafChildCount = 2
	LeafChildCount = 1

)
type SearchCriteriaNode struct {
	OpType SearchOpType
	CritType SearchCriteriaType
	Field string
	Value string
	LeftChild *SearchCriteriaNode
	RightChild *SearchCriteriaNode
}

type EmsSearchParserListener struct {
	*search_parser.BaseSearchParserListener
}

func NewEmsSearchParserListener() *EmsSearchParserListener {
	return new(EmsSearchParserListener)
}

func (s *EmsSearchParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.Printf("Enter every rule: %v\n", ctx.GetText())
}

type EmsSearchParserVisitor struct {
	*search_parser.BaseSearchParserVisitor
	SearchTree *SearchCriteriaNode
}

func NewEmsSearchParserVisitor() *EmsSearchParserVisitor {
	return new(EmsSearchParserVisitor)
}

func (v *EmsSearchParserVisitor) visit(tree antlr.Tree) interface{} {
	log.Printf("Start visit child. Type[%v]\n", reflect.TypeOf(tree))

	if ctx, ok := tree.(antlr.ParserRuleContext); ok {
		return ctx.Accept(v)
	} else if node, ok := tree.(antlr.ErrorNode); ok {
		log.Printf("Error node. text[%s]\n", node.GetText())
	} else if node, ok := tree.(antlr.TerminalNode); ok {
		log.Printf("Terminal node. text[%s]\n", node.GetText())
	} else {
		log.Printf("Unsupport child: %v\n", reflect.TypeOf(tree))
	}
	return nil
}

func (v *EmsSearchParserVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	n := node.GetChildCount()
	log.Printf("Start visit children. type[%v] count[%d] text[%s]\n", reflect.TypeOf(node), n, node.GetText())
	for i := 0; i < n; i++ {
		c := node.GetChild(i)
		v.visit(c)
	}
	return nil
}

func (v *EmsSearchParserVisitor) VisitQuery(ctx *search_parser.QueryContext) interface{} {
	n := ctx.GetChildCount()
	if n != LeafChildCount {
		log.Fatalf("Visit query count error, should be 1. count[%d] text[%s]",
			n, ctx.GetText())
			return nil
	}
	log.Printf("Start visit children. type[%v] count[%d] text[%s]\n", reflect.TypeOf(ctx), n, ctx.GetText())
	child := ctx.GetChild(0)
	return v.visit(child)
}

func (v *EmsSearchParserVisitor) VisitAndQueryCriterias(ctx *search_parser.AndQueryCriteriasContext) interface{} {
	criterias := ctx.AllQueryCriterias()
	count := len(criterias)
	if count != NonLeafChildCount {
		log.Fatalf("queryCriterias count[%d] need to be %d. text[%s]",
			count, NonLeafChildCount, ctx.GetText())
		return nil
	}
	node := &SearchCriteriaNode{
		OpType: OPERATOR_AND,
	}
	leftChild := v.visit(criterias[0])
	if lc, ok := leftChild.(*SearchCriteriaNode); ok {
		node.LeftChild = lc
	}
	rightChild := v.visit(criterias[1])
	if rc, ok := rightChild.(*SearchCriteriaNode); ok {
		node.RightChild = rc
	}
	return node
}

func (v *EmsSearchParserVisitor) VisitNotQueryCriterias(ctx *search_parser.NotQueryCriteriasContext) interface{} {
	criteria := ctx.QueryCriterias()
	if criteria == nil {
		log.Fatalf("visit NotQueryCriterias has no QueryCriterias. text[%s]", ctx.GetText())
		return nil
	}
	res := v.visit(criteria)
	if cn, ok := res.(*SearchCriteriaNode); ok {
		// Multiple not
		if cn.OpType == OPERATOR_NOT {
			cn.OpType = OPERATOR_NONE
			return cn
		}
		node := &SearchCriteriaNode{
			OpType: OPERATOR_NOT,
			LeftChild: cn,
		}
		return node
	}
	return nil
}

func (v *EmsSearchParserVisitor) VisitSingleQueryCriteria(ctx *search_parser.SingleQueryCriteriaContext) interface{} {
	criteria := ctx.QueryCriteria()
	if criteria == nil {
		log.Fatalf("visit SingleQueryCriterias has no QueryCriterias. text[%s]", ctx.GetText())
		return nil
	}
	return v.visit(criteria)
}

func (v *EmsSearchParserVisitor) VisitBracketQueryCriterias(ctx *search_parser.BracketQueryCriteriasContext) interface{} {
	criteria := ctx.QueryCriterias()
	if criteria == nil {
		log.Fatalf("visit BracketQueryCriterias has no QueryCriterias. text[%s]", ctx.GetText())
		return nil
	}
	return v.visit(criteria)
}

func (v *EmsSearchParserVisitor) VisitOrQueryCriterias(ctx *search_parser.OrQueryCriteriasContext) interface{} {
	log.Printf("Start VisitOrQueryCriterias")
	criterias := ctx.AllQueryCriterias()
	count := len(criterias)
	if count != NonLeafChildCount {
		log.Fatalf("queryCriterias count[%d] need to be %d", count, NonLeafChildCount)
		return nil
	}
	node := &SearchCriteriaNode{
		OpType: OPERATOR_OR,
	}
	leftChild := v.visit(criterias[0])
	log.Printf("Left child type[%v] detail[%v]", reflect.TypeOf(leftChild), leftChild)
	if lc, ok := leftChild.(*SearchCriteriaNode); ok {
		node.LeftChild = lc
	}
	rightChild := v.visit(criterias[1])
	log.Printf("Right child type[%v] detail[%v]", reflect.TypeOf(rightChild), rightChild)
	if rc, ok := rightChild.(*SearchCriteriaNode); ok {
		node.RightChild = rc
		log.Printf("Right child set type[%v] detail[%v]", reflect.TypeOf(rc), rc)
	}
	return node
}

func (v *EmsSearchParserVisitor) VisitQueryCriteria(ctx *search_parser.QueryCriteriaContext) interface{} {
	n := ctx.GetChildCount()
	if n != LeafChildCount {
		log.Fatalf("Visit QueryCriteria count error, should be 1. count[%d] text[%s]",
			n, ctx.GetText())
	}
	log.Printf("Start visit QueryCriteria. type[%v] count[%d] text[%s]\n", reflect.TypeOf(ctx), n, ctx.GetText())
	child := ctx.GetChild(0)
	return v.visit(child)
}

func (v *EmsSearchParserVisitor) VisitHeaderCriteria(ctx *search_parser.HeaderCriteriaContext) interface{} {
	hl := strings.Split(ctx.HEADER_START().GetText(), ":")
	value := ctx.ValueCriteria()
	if value == nil {
		return nil
	}
	node := &SearchCriteriaNode{
		OpType: OPERATOR_NONE,
		CritType: CRITERIA_HEADER,
		Field: hl[0],
		Value: value.GetText(),
	}
	return node
}

func (v *EmsSearchParserVisitor) VisitFieldCriteria(ctx *search_parser.FieldCriteriaContext) interface{} {
	hl := strings.Split(ctx.FIELD_START().GetText(), ":")
	value := ctx.ValueCriteria()
	if value == nil {
		return nil
	}
	node := &SearchCriteriaNode{
		OpType: OPERATOR_NONE,
		CritType: CRITERIA_FIELD,
		Field: hl[0],
		Value: value.GetText(),
	}
	return node
}

func (v *EmsSearchParserVisitor) VisitValueCriteria(ctx *search_parser.ValueCriteriaContext) interface{} {
	node := &SearchCriteriaNode{
		OpType: OPERATOR_NONE,
		CritType: CRITERIA_VALUE,
		Value: ctx.GetText(),
	}
	return node
}

// var t = "test:item key:\" value \" and key2:value2 test test"
var t = "subject:test AND header[to]:jinlian AND test (SUMMARY OR user) label:inbox OR (from)"
// var t = "subject:"
func testSearchParserListener() {
	log.Printf("\ntestQueryParserListener: %s\n", t)
	input := antlr.NewInputStream(t)
	// input, _ := antlr.NewFileStream(fn)
	lexer := search_parser.NewSearchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := search_parser.NewSearchParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	antlr.ParseTreeWalkerDefault.Walk(NewEmsSearchParserListener(), tree)
}

func testSearchParserVisitor() {
	log.Printf("\n\ntestVisitor: %s\n", t)
	input := antlr.NewInputStream(t)
	lexer := search_parser.NewSearchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := search_parser.NewSearchParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	visitor := NewEmsSearchParserVisitor()
	t := p.Query().Accept(visitor)
	log.Printf("parse result: %v\n", t)
	visitNode(t.(*SearchCriteriaNode))
}

func visitNode(root *SearchCriteriaNode) {
	log.Printf("node optype[%v] criteriatype[%v] field[%v] value[%v]",
		root.OpType, root.CritType, root.Field, root.Value)
	if root.LeftChild != nil {
		visitNode(root.LeftChild)
	}
	if root.RightChild != nil {
		visitNode(root.RightChild)
	}
}

func main() {
	log.SetFlags(log.Lshortfile|log.LstdFlags)
	// testSearchParserListener()
	testSearchParserVisitor()
}