// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 53, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10, 3, 12, 3, 14, 3, 36,
	11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 41, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 3, 4, 8, 2, 4, 6, 8, 10, 12, 2,
	3, 4, 2, 9, 9, 11, 12, 2, 52, 2, 14, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6,
	40, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 46, 3, 2, 2, 2, 12, 50, 3, 2, 2,
	2, 14, 15, 5, 4, 3, 2, 15, 3, 3, 2, 2, 2, 16, 17, 8, 3, 1, 2, 17, 18, 7,
	7, 2, 2, 18, 25, 5, 4, 3, 5, 19, 20, 7, 3, 2, 2, 20, 21, 5, 4, 3, 2, 21,
	22, 7, 4, 2, 2, 22, 25, 3, 2, 2, 2, 23, 25, 5, 6, 4, 2, 24, 16, 3, 2, 2,
	2, 24, 19, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 34, 3, 2, 2, 2, 26, 27,
	12, 7, 2, 2, 27, 28, 7, 8, 2, 2, 28, 33, 5, 4, 3, 8, 29, 30, 12, 6, 2,
	2, 30, 31, 7, 6, 2, 2, 31, 33, 5, 4, 3, 7, 32, 26, 3, 2, 2, 2, 32, 29,
	3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 5, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 41, 5, 8, 5, 2, 38, 41, 5, 10,
	6, 2, 39, 41, 5, 12, 7, 2, 40, 37, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40,
	39, 3, 2, 2, 2, 41, 7, 3, 2, 2, 2, 42, 43, 7, 9, 2, 2, 43, 44, 7, 5, 2,
	2, 44, 45, 5, 12, 7, 2, 45, 9, 3, 2, 2, 2, 46, 47, 7, 10, 2, 2, 47, 48,
	7, 5, 2, 2, 48, 49, 5, 12, 7, 2, 49, 11, 3, 2, 2, 2, 50, 51, 9, 2, 2, 2,
	51, 13, 3, 2, 2, 2, 6, 24, 32, 34, 40,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "OR", "NOT", "AND", "KEY", "HEADER", "NOQUOTE_STRING",
	"QUOTE_STRING", "WS",
}

var ruleNames = []string{
	"query", "queryCriterias", "queryCriteria", "keyCriteria", "headerCriteria",
	"valueCriteria",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type QueryParser struct {
	*antlr.BaseParser
}

func NewQueryParser(input antlr.TokenStream) *QueryParser {
	this := new(QueryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Query.g4"

	return this
}

// QueryParser tokens.
const (
	QueryParserEOF            = antlr.TokenEOF
	QueryParserT__0           = 1
	QueryParserT__1           = 2
	QueryParserT__2           = 3
	QueryParserOR             = 4
	QueryParserNOT            = 5
	QueryParserAND            = 6
	QueryParserKEY            = 7
	QueryParserHEADER         = 8
	QueryParserNOQUOTE_STRING = 9
	QueryParserQUOTE_STRING   = 10
	QueryParserWS             = 11
)

// QueryParser rules.
const (
	QueryParserRULE_query          = 0
	QueryParserRULE_queryCriterias = 1
	QueryParserRULE_queryCriteria  = 2
	QueryParserRULE_keyCriteria    = 3
	QueryParserRULE_headerCriteria = 4
	QueryParserRULE_valueCriteria  = 5
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) QueryCriterias() IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.queryCriterias(0)
	}

	return localctx
}

// IQueryCriteriasContext is an interface to support dynamic dispatch.
type IQueryCriteriasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryCriteriasContext differentiates from other interfaces.
	IsQueryCriteriasContext()
}

type QueryCriteriasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryCriteriasContext() *QueryCriteriasContext {
	var p = new(QueryCriteriasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_queryCriterias
	return p
}

func (*QueryCriteriasContext) IsQueryCriteriasContext() {}

func NewQueryCriteriasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryCriteriasContext {
	var p = new(QueryCriteriasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_queryCriterias

	return p
}

func (s *QueryCriteriasContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryCriteriasContext) CopyFrom(ctx *QueryCriteriasContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryCriteriasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryCriteriasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndQueryCriteriasContext struct {
	*QueryCriteriasContext
}

func NewAndQueryCriteriasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndQueryCriteriasContext {
	var p = new(AndQueryCriteriasContext)

	p.QueryCriteriasContext = NewEmptyQueryCriteriasContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryCriteriasContext))

	return p
}

func (s *AndQueryCriteriasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndQueryCriteriasContext) AllQueryCriterias() []IQueryCriteriasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem())
	var tst = make([]IQueryCriteriasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryCriteriasContext)
		}
	}

	return tst
}

func (s *AndQueryCriteriasContext) QueryCriterias(i int) IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *AndQueryCriteriasContext) AND() antlr.TerminalNode {
	return s.GetToken(QueryParserAND, 0)
}

func (s *AndQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterAndQueryCriterias(s)
	}
}

func (s *AndQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitAndQueryCriterias(s)
	}
}

func (s *AndQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitAndQueryCriterias(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotQueryCriteriasContext struct {
	*QueryCriteriasContext
}

func NewNotQueryCriteriasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotQueryCriteriasContext {
	var p = new(NotQueryCriteriasContext)

	p.QueryCriteriasContext = NewEmptyQueryCriteriasContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryCriteriasContext))

	return p
}

func (s *NotQueryCriteriasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotQueryCriteriasContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *NotQueryCriteriasContext) QueryCriterias() IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *NotQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterNotQueryCriterias(s)
	}
}

func (s *NotQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitNotQueryCriterias(s)
	}
}

func (s *NotQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNotQueryCriterias(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleQueryCriteriaContext struct {
	*QueryCriteriasContext
}

func NewSingleQueryCriteriaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleQueryCriteriaContext {
	var p = new(SingleQueryCriteriaContext)

	p.QueryCriteriasContext = NewEmptyQueryCriteriasContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryCriteriasContext))

	return p
}

func (s *SingleQueryCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleQueryCriteriaContext) QueryCriteria() IQueryCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriaContext)
}

func (s *SingleQueryCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterSingleQueryCriteria(s)
	}
}

func (s *SingleQueryCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitSingleQueryCriteria(s)
	}
}

func (s *SingleQueryCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitSingleQueryCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketQueryCriteriasContext struct {
	*QueryCriteriasContext
}

func NewBracketQueryCriteriasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketQueryCriteriasContext {
	var p = new(BracketQueryCriteriasContext)

	p.QueryCriteriasContext = NewEmptyQueryCriteriasContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryCriteriasContext))

	return p
}

func (s *BracketQueryCriteriasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketQueryCriteriasContext) QueryCriterias() IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *BracketQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterBracketQueryCriterias(s)
	}
}

func (s *BracketQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitBracketQueryCriterias(s)
	}
}

func (s *BracketQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitBracketQueryCriterias(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrQueryCriteriasContext struct {
	*QueryCriteriasContext
}

func NewOrQueryCriteriasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrQueryCriteriasContext {
	var p = new(OrQueryCriteriasContext)

	p.QueryCriteriasContext = NewEmptyQueryCriteriasContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryCriteriasContext))

	return p
}

func (s *OrQueryCriteriasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrQueryCriteriasContext) AllQueryCriterias() []IQueryCriteriasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem())
	var tst = make([]IQueryCriteriasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryCriteriasContext)
		}
	}

	return tst
}

func (s *OrQueryCriteriasContext) QueryCriterias(i int) IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *OrQueryCriteriasContext) OR() antlr.TerminalNode {
	return s.GetToken(QueryParserOR, 0)
}

func (s *OrQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOrQueryCriterias(s)
	}
}

func (s *OrQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOrQueryCriterias(s)
	}
}

func (s *OrQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitOrQueryCriterias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) QueryCriterias() (localctx IQueryCriteriasContext) {
	return p.queryCriterias(0)
}

func (p *QueryParser) queryCriterias(_p int) (localctx IQueryCriteriasContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryCriteriasContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryCriteriasContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, QueryParserRULE_queryCriterias, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserNOT:
		localctx = NewNotQueryCriteriasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(15)
			p.Match(QueryParserNOT)
		}
		{
			p.SetState(16)
			p.queryCriterias(3)
		}

	case QueryParserT__0:
		localctx = NewBracketQueryCriteriasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(18)
			p.queryCriterias(0)
		}
		{
			p.SetState(19)
			p.Match(QueryParserT__1)
		}

	case QueryParserKEY, QueryParserHEADER, QueryParserNOQUOTE_STRING, QueryParserQUOTE_STRING:
		localctx = NewSingleQueryCriteriaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.QueryCriteria()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndQueryCriteriasContext(p, NewQueryCriteriasContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_queryCriterias)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(25)
					p.Match(QueryParserAND)
				}
				{
					p.SetState(26)
					p.queryCriterias(6)
				}

			case 2:
				localctx = NewOrQueryCriteriasContext(p, NewQueryCriteriasContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_queryCriterias)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(28)
					p.Match(QueryParserOR)
				}
				{
					p.SetState(29)
					p.queryCriterias(5)
				}

			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IQueryCriteriaContext is an interface to support dynamic dispatch.
type IQueryCriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryCriteriaContext differentiates from other interfaces.
	IsQueryCriteriaContext()
}

type QueryCriteriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryCriteriaContext() *QueryCriteriaContext {
	var p = new(QueryCriteriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_queryCriteria
	return p
}

func (*QueryCriteriaContext) IsQueryCriteriaContext() {}

func NewQueryCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryCriteriaContext {
	var p = new(QueryCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_queryCriteria

	return p
}

func (s *QueryCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryCriteriaContext) KeyCriteria() IKeyCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyCriteriaContext)
}

func (s *QueryCriteriaContext) HeaderCriteria() IHeaderCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderCriteriaContext)
}

func (s *QueryCriteriaContext) ValueCriteria() IValueCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueCriteriaContext)
}

func (s *QueryCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterQueryCriteria(s)
	}
}

func (s *QueryCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitQueryCriteria(s)
	}
}

func (s *QueryCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitQueryCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) QueryCriteria() (localctx IQueryCriteriaContext) {
	localctx = NewQueryCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_queryCriteria)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.KeyCriteria()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.HeaderCriteria()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.ValueCriteria()
		}

	}

	return localctx
}

// IKeyCriteriaContext is an interface to support dynamic dispatch.
type IKeyCriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyCriteriaContext differentiates from other interfaces.
	IsKeyCriteriaContext()
}

type KeyCriteriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyCriteriaContext() *KeyCriteriaContext {
	var p = new(KeyCriteriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_keyCriteria
	return p
}

func (*KeyCriteriaContext) IsKeyCriteriaContext() {}

func NewKeyCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyCriteriaContext {
	var p = new(KeyCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_keyCriteria

	return p
}

func (s *KeyCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyCriteriaContext) KEY() antlr.TerminalNode {
	return s.GetToken(QueryParserKEY, 0)
}

func (s *KeyCriteriaContext) ValueCriteria() IValueCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueCriteriaContext)
}

func (s *KeyCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterKeyCriteria(s)
	}
}

func (s *KeyCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitKeyCriteria(s)
	}
}

func (s *KeyCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitKeyCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) KeyCriteria() (localctx IKeyCriteriaContext) {
	localctx = NewKeyCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryParserRULE_keyCriteria)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(QueryParserKEY)
	}
	{
		p.SetState(41)
		p.Match(QueryParserT__2)
	}
	{
		p.SetState(42)
		p.ValueCriteria()
	}

	return localctx
}

// IHeaderCriteriaContext is an interface to support dynamic dispatch.
type IHeaderCriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderCriteriaContext differentiates from other interfaces.
	IsHeaderCriteriaContext()
}

type HeaderCriteriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderCriteriaContext() *HeaderCriteriaContext {
	var p = new(HeaderCriteriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_headerCriteria
	return p
}

func (*HeaderCriteriaContext) IsHeaderCriteriaContext() {}

func NewHeaderCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderCriteriaContext {
	var p = new(HeaderCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_headerCriteria

	return p
}

func (s *HeaderCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderCriteriaContext) HEADER() antlr.TerminalNode {
	return s.GetToken(QueryParserHEADER, 0)
}

func (s *HeaderCriteriaContext) ValueCriteria() IValueCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueCriteriaContext)
}

func (s *HeaderCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterHeaderCriteria(s)
	}
}

func (s *HeaderCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitHeaderCriteria(s)
	}
}

func (s *HeaderCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitHeaderCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) HeaderCriteria() (localctx IHeaderCriteriaContext) {
	localctx = NewHeaderCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryParserRULE_headerCriteria)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(QueryParserHEADER)
	}
	{
		p.SetState(45)
		p.Match(QueryParserT__2)
	}
	{
		p.SetState(46)
		p.ValueCriteria()
	}

	return localctx
}

// IValueCriteriaContext is an interface to support dynamic dispatch.
type IValueCriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueCriteriaContext differentiates from other interfaces.
	IsValueCriteriaContext()
}

type ValueCriteriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueCriteriaContext() *ValueCriteriaContext {
	var p = new(ValueCriteriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_valueCriteria
	return p
}

func (*ValueCriteriaContext) IsValueCriteriaContext() {}

func NewValueCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueCriteriaContext {
	var p = new(ValueCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_valueCriteria

	return p
}

func (s *ValueCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueCriteriaContext) KEY() antlr.TerminalNode {
	return s.GetToken(QueryParserKEY, 0)
}

func (s *ValueCriteriaContext) QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserQUOTE_STRING, 0)
}

func (s *ValueCriteriaContext) NOQUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserNOQUOTE_STRING, 0)
}

func (s *ValueCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterValueCriteria(s)
	}
}

func (s *ValueCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitValueCriteria(s)
	}
}

func (s *ValueCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitValueCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) ValueCriteria() (localctx IValueCriteriaContext) {
	localctx = NewValueCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryParserRULE_valueCriteria)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserKEY)|(1<<QueryParserNOQUOTE_STRING)|(1<<QueryParserQUOTE_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *QueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *QueryCriteriasContext = nil
		if localctx != nil {
			t = localctx.(*QueryCriteriasContext)
		}
		return p.QueryCriterias_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QueryParser) QueryCriterias_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
