// Code generated from SearchParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package search_parser // SearchParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 53, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 3, 3, 3, 5, 3, 29, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 35, 10, 3,
	12, 3, 14, 3, 38, 11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 3, 4, 8, 2, 4, 6, 8, 10, 12,
	2, 3, 4, 2, 8, 8, 13, 14, 2, 53, 2, 14, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2,
	6, 42, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10, 47, 3, 2, 2, 2, 12, 50, 3, 2,
	2, 2, 14, 15, 5, 4, 3, 2, 15, 3, 3, 2, 2, 2, 16, 17, 8, 3, 1, 2, 17, 18,
	7, 4, 2, 2, 18, 25, 5, 4, 3, 5, 19, 20, 7, 6, 2, 2, 20, 21, 5, 4, 3, 2,
	21, 22, 7, 7, 2, 2, 22, 25, 3, 2, 2, 2, 23, 25, 5, 6, 4, 2, 24, 16, 3,
	2, 2, 2, 24, 19, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 36, 3, 2, 2, 2, 26,
	28, 12, 7, 2, 2, 27, 29, 7, 5, 2, 2, 28, 27, 3, 2, 2, 2, 28, 29, 3, 2,
	2, 2, 29, 30, 3, 2, 2, 2, 30, 35, 5, 4, 3, 8, 31, 32, 12, 6, 2, 2, 32,
	33, 7, 3, 2, 2, 33, 35, 5, 4, 3, 7, 34, 26, 3, 2, 2, 2, 34, 31, 3, 2, 2,
	2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 5, 3,
	2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 43, 5, 10, 6, 2, 40, 43, 5, 8, 5, 2, 41,
	43, 5, 12, 7, 2, 42, 39, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2,
	2, 2, 43, 7, 3, 2, 2, 2, 44, 45, 7, 11, 2, 2, 45, 46, 5, 12, 7, 2, 46,
	9, 3, 2, 2, 2, 47, 48, 7, 10, 2, 2, 48, 49, 5, 12, 7, 2, 49, 11, 3, 2,
	2, 2, 50, 51, 9, 2, 2, 2, 51, 13, 3, 2, 2, 2, 7, 24, 28, 34, 36, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'('", "')'",
}
var symbolicNames = []string{
	"", "OR", "NOT", "AND", "LP", "RP", "FIELD", "HEADER", "FIELD_START", "HEADER_START",
	"WS", "NOQUOTE_STRING", "QUOTE_STRING",
}

var ruleNames = []string{
	"query", "queryCriterias", "queryCriteria", "headerCriteria", "fieldCriteria",
	"valueCriteria",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SearchParser struct {
	*antlr.BaseParser
}

func NewSearchParser(input antlr.TokenStream) *SearchParser {
	this := new(SearchParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SearchParser.g4"

	return this
}

// SearchParser tokens.
const (
	SearchParserEOF            = antlr.TokenEOF
	SearchParserOR             = 1
	SearchParserNOT            = 2
	SearchParserAND            = 3
	SearchParserLP             = 4
	SearchParserRP             = 5
	SearchParserFIELD          = 6
	SearchParserHEADER         = 7
	SearchParserFIELD_START    = 8
	SearchParserHEADER_START   = 9
	SearchParserWS             = 10
	SearchParserNOQUOTE_STRING = 11
	SearchParserQUOTE_STRING   = 12
)

// SearchParser rules.
const (
	SearchParserRULE_query          = 0
	SearchParserRULE_queryCriterias = 1
	SearchParserRULE_queryCriteria  = 2
	SearchParserRULE_headerCriteria = 3
	SearchParserRULE_fieldCriteria  = 4
	SearchParserRULE_valueCriteria  = 5
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
	p.RuleIndex = SearchParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchParserRULE_query

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
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SearchParserRULE_query)

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
	p.RuleIndex = SearchParserRULE_queryCriterias
	return p
}

func (*QueryCriteriasContext) IsQueryCriteriasContext() {}

func NewQueryCriteriasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryCriteriasContext {
	var p = new(QueryCriteriasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchParserRULE_queryCriterias

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
	return s.GetToken(SearchParserAND, 0)
}

func (s *AndQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterAndQueryCriterias(s)
	}
}

func (s *AndQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitAndQueryCriterias(s)
	}
}

func (s *AndQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
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
	return s.GetToken(SearchParserNOT, 0)
}

func (s *NotQueryCriteriasContext) QueryCriterias() IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *NotQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterNotQueryCriterias(s)
	}
}

func (s *NotQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitNotQueryCriterias(s)
	}
}

func (s *NotQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
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
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterSingleQueryCriteria(s)
	}
}

func (s *SingleQueryCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitSingleQueryCriteria(s)
	}
}

func (s *SingleQueryCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
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

func (s *BracketQueryCriteriasContext) LP() antlr.TerminalNode {
	return s.GetToken(SearchParserLP, 0)
}

func (s *BracketQueryCriteriasContext) QueryCriterias() IQueryCriteriasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryCriteriasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryCriteriasContext)
}

func (s *BracketQueryCriteriasContext) RP() antlr.TerminalNode {
	return s.GetToken(SearchParserRP, 0)
}

func (s *BracketQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterBracketQueryCriterias(s)
	}
}

func (s *BracketQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitBracketQueryCriterias(s)
	}
}

func (s *BracketQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
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
	return s.GetToken(SearchParserOR, 0)
}

func (s *OrQueryCriteriasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterOrQueryCriterias(s)
	}
}

func (s *OrQueryCriteriasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitOrQueryCriterias(s)
	}
}

func (s *OrQueryCriteriasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
		return t.VisitOrQueryCriterias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchParser) QueryCriterias() (localctx IQueryCriteriasContext) {
	return p.queryCriterias(0)
}

func (p *SearchParser) queryCriterias(_p int) (localctx IQueryCriteriasContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryCriteriasContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryCriteriasContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, SearchParserRULE_queryCriterias, _p)
	var _la int

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
	case SearchParserNOT:
		localctx = NewNotQueryCriteriasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(15)
			p.Match(SearchParserNOT)
		}
		{
			p.SetState(16)
			p.queryCriterias(3)
		}

	case SearchParserLP:
		localctx = NewBracketQueryCriteriasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(SearchParserLP)
		}
		{
			p.SetState(18)
			p.queryCriterias(0)
		}
		{
			p.SetState(19)
			p.Match(SearchParserRP)
		}

	case SearchParserFIELD, SearchParserFIELD_START, SearchParserHEADER_START, SearchParserNOQUOTE_STRING, SearchParserQUOTE_STRING:
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndQueryCriteriasContext(p, NewQueryCriteriasContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SearchParserRULE_queryCriterias)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(26)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SearchParserAND {
					{
						p.SetState(25)
						p.Match(SearchParserAND)
					}

				}
				{
					p.SetState(28)
					p.queryCriterias(6)
				}

			case 2:
				localctx = NewOrQueryCriteriasContext(p, NewQueryCriteriasContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SearchParserRULE_queryCriterias)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(30)
					p.Match(SearchParserOR)
				}
				{
					p.SetState(31)
					p.queryCriterias(5)
				}

			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.RuleIndex = SearchParserRULE_queryCriteria
	return p
}

func (*QueryCriteriaContext) IsQueryCriteriaContext() {}

func NewQueryCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryCriteriaContext {
	var p = new(QueryCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchParserRULE_queryCriteria

	return p
}

func (s *QueryCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryCriteriaContext) FieldCriteria() IFieldCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldCriteriaContext)
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
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterQueryCriteria(s)
	}
}

func (s *QueryCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitQueryCriteria(s)
	}
}

func (s *QueryCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
		return t.VisitQueryCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchParser) QueryCriteria() (localctx IQueryCriteriaContext) {
	localctx = NewQueryCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SearchParserRULE_queryCriteria)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SearchParserFIELD_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.FieldCriteria()
		}

	case SearchParserHEADER_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.HeaderCriteria()
		}

	case SearchParserFIELD, SearchParserNOQUOTE_STRING, SearchParserQUOTE_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.ValueCriteria()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = SearchParserRULE_headerCriteria
	return p
}

func (*HeaderCriteriaContext) IsHeaderCriteriaContext() {}

func NewHeaderCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderCriteriaContext {
	var p = new(HeaderCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchParserRULE_headerCriteria

	return p
}

func (s *HeaderCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderCriteriaContext) HEADER_START() antlr.TerminalNode {
	return s.GetToken(SearchParserHEADER_START, 0)
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
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterHeaderCriteria(s)
	}
}

func (s *HeaderCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitHeaderCriteria(s)
	}
}

func (s *HeaderCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
		return t.VisitHeaderCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchParser) HeaderCriteria() (localctx IHeaderCriteriaContext) {
	localctx = NewHeaderCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SearchParserRULE_headerCriteria)

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
		p.SetState(42)
		p.Match(SearchParserHEADER_START)
	}
	{
		p.SetState(43)
		p.ValueCriteria()
	}

	return localctx
}

// IFieldCriteriaContext is an interface to support dynamic dispatch.
type IFieldCriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldCriteriaContext differentiates from other interfaces.
	IsFieldCriteriaContext()
}

type FieldCriteriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldCriteriaContext() *FieldCriteriaContext {
	var p = new(FieldCriteriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SearchParserRULE_fieldCriteria
	return p
}

func (*FieldCriteriaContext) IsFieldCriteriaContext() {}

func NewFieldCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldCriteriaContext {
	var p = new(FieldCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchParserRULE_fieldCriteria

	return p
}

func (s *FieldCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldCriteriaContext) FIELD_START() antlr.TerminalNode {
	return s.GetToken(SearchParserFIELD_START, 0)
}

func (s *FieldCriteriaContext) ValueCriteria() IValueCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueCriteriaContext)
}

func (s *FieldCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterFieldCriteria(s)
	}
}

func (s *FieldCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitFieldCriteria(s)
	}
}

func (s *FieldCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
		return t.VisitFieldCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchParser) FieldCriteria() (localctx IFieldCriteriaContext) {
	localctx = NewFieldCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SearchParserRULE_fieldCriteria)

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
		p.SetState(45)
		p.Match(SearchParserFIELD_START)
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
	p.RuleIndex = SearchParserRULE_valueCriteria
	return p
}

func (*ValueCriteriaContext) IsValueCriteriaContext() {}

func NewValueCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueCriteriaContext {
	var p = new(ValueCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SearchParserRULE_valueCriteria

	return p
}

func (s *ValueCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueCriteriaContext) QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(SearchParserQUOTE_STRING, 0)
}

func (s *ValueCriteriaContext) NOQUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(SearchParserNOQUOTE_STRING, 0)
}

func (s *ValueCriteriaContext) FIELD() antlr.TerminalNode {
	return s.GetToken(SearchParserFIELD, 0)
}

func (s *ValueCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.EnterValueCriteria(s)
	}
}

func (s *ValueCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SearchParserListener); ok {
		listenerT.ExitValueCriteria(s)
	}
}

func (s *ValueCriteriaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SearchParserVisitor:
		return t.VisitValueCriteria(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SearchParser) ValueCriteria() (localctx IValueCriteriaContext) {
	localctx = NewValueCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SearchParserRULE_valueCriteria)
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

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SearchParserFIELD)|(1<<SearchParserNOQUOTE_STRING)|(1<<SearchParserQUOTE_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SearchParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *SearchParser) QueryCriterias_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
