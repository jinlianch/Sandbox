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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 41, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 19, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6, 2, 2, 2, 43,
	2, 8, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 9, 5, 4, 3,
	2, 9, 3, 3, 2, 2, 2, 10, 11, 8, 3, 1, 2, 11, 12, 7, 8, 2, 2, 12, 19, 5,
	4, 3, 5, 13, 14, 7, 3, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 7, 4, 2, 2, 16,
	19, 3, 2, 2, 2, 17, 19, 5, 6, 4, 2, 18, 10, 3, 2, 2, 2, 18, 13, 3, 2, 2,
	2, 18, 17, 3, 2, 2, 2, 19, 28, 3, 2, 2, 2, 20, 21, 12, 7, 2, 2, 21, 22,
	7, 6, 2, 2, 22, 27, 5, 4, 3, 8, 23, 24, 12, 6, 2, 2, 24, 25, 7, 7, 2, 2,
	25, 27, 5, 4, 3, 7, 26, 20, 3, 2, 2, 2, 26, 23, 3, 2, 2, 2, 27, 30, 3,
	2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30,
	28, 3, 2, 2, 2, 31, 32, 7, 9, 2, 2, 32, 33, 7, 5, 2, 2, 33, 39, 7, 10,
	2, 2, 34, 35, 7, 11, 2, 2, 35, 36, 7, 5, 2, 2, 36, 39, 7, 10, 2, 2, 37,
	39, 7, 10, 2, 2, 38, 31, 3, 2, 2, 2, 38, 34, 3, 2, 2, 2, 38, 37, 3, 2,
	2, 2, 39, 7, 3, 2, 2, 2, 6, 18, 26, 28, 38,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "AND", "OR", "NOT", "ID", "STRING", "HEADER", "WS",
}

var ruleNames = []string{
	"query", "query_item", "single_query_item",
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
	QueryParserEOF    = antlr.TokenEOF
	QueryParserT__0   = 1
	QueryParserT__1   = 2
	QueryParserT__2   = 3
	QueryParserAND    = 4
	QueryParserOR     = 5
	QueryParserNOT    = 6
	QueryParserID     = 7
	QueryParserSTRING = 8
	QueryParserHEADER = 9
	QueryParserWS     = 10
)

// QueryParser rules.
const (
	QueryParserRULE_query             = 0
	QueryParserRULE_query_item        = 1
	QueryParserRULE_single_query_item = 2
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

func (s *QueryContext) Query_item() IQuery_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuery_itemContext)
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
		p.SetState(6)
		p.query_item(0)
	}

	return localctx
}

// IQuery_itemContext is an interface to support dynamic dispatch.
type IQuery_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuery_itemContext differentiates from other interfaces.
	IsQuery_itemContext()
}

type Query_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuery_itemContext() *Query_itemContext {
	var p = new(Query_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_query_item
	return p
}

func (*Query_itemContext) IsQuery_itemContext() {}

func NewQuery_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Query_itemContext {
	var p = new(Query_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_query_item

	return p
}

func (s *Query_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Query_itemContext) CopyFrom(ctx *Query_itemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Query_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Query_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BracketQueryItemContext struct {
	*Query_itemContext
}

func NewBracketQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketQueryItemContext {
	var p = new(BracketQueryItemContext)

	p.Query_itemContext = NewEmptyQuery_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Query_itemContext))

	return p
}

func (s *BracketQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketQueryItemContext) Query_item() IQuery_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuery_itemContext)
}

func (s *BracketQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterBracketQueryItem(s)
	}
}

func (s *BracketQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitBracketQueryItem(s)
	}
}

func (s *BracketQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitBracketQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleQueryItemContext struct {
	*Query_itemContext
}

func NewSingleQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleQueryItemContext {
	var p = new(SingleQueryItemContext)

	p.Query_itemContext = NewEmptyQuery_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Query_itemContext))

	return p
}

func (s *SingleQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleQueryItemContext) Single_query_item() ISingle_query_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_query_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingle_query_itemContext)
}

func (s *SingleQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterSingleQueryItem(s)
	}
}

func (s *SingleQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitSingleQueryItem(s)
	}
}

func (s *SingleQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitSingleQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotQueryItemContext struct {
	*Query_itemContext
}

func NewNotQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotQueryItemContext {
	var p = new(NotQueryItemContext)

	p.Query_itemContext = NewEmptyQuery_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Query_itemContext))

	return p
}

func (s *NotQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotQueryItemContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *NotQueryItemContext) Query_item() IQuery_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuery_itemContext)
}

func (s *NotQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterNotQueryItem(s)
	}
}

func (s *NotQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitNotQueryItem(s)
	}
}

func (s *NotQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNotQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrQueryItemContext struct {
	*Query_itemContext
}

func NewOrQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrQueryItemContext {
	var p = new(OrQueryItemContext)

	p.Query_itemContext = NewEmptyQuery_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Query_itemContext))

	return p
}

func (s *OrQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrQueryItemContext) AllQuery_item() []IQuery_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem())
	var tst = make([]IQuery_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuery_itemContext)
		}
	}

	return tst
}

func (s *OrQueryItemContext) Query_item(i int) IQuery_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuery_itemContext)
}

func (s *OrQueryItemContext) OR() antlr.TerminalNode {
	return s.GetToken(QueryParserOR, 0)
}

func (s *OrQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOrQueryItem(s)
	}
}

func (s *OrQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOrQueryItem(s)
	}
}

func (s *OrQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitOrQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndQueryItemContext struct {
	*Query_itemContext
}

func NewAndQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndQueryItemContext {
	var p = new(AndQueryItemContext)

	p.Query_itemContext = NewEmptyQuery_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Query_itemContext))

	return p
}

func (s *AndQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndQueryItemContext) AllQuery_item() []IQuery_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem())
	var tst = make([]IQuery_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuery_itemContext)
		}
	}

	return tst
}

func (s *AndQueryItemContext) Query_item(i int) IQuery_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuery_itemContext)
}

func (s *AndQueryItemContext) AND() antlr.TerminalNode {
	return s.GetToken(QueryParserAND, 0)
}

func (s *AndQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterAndQueryItem(s)
	}
}

func (s *AndQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitAndQueryItem(s)
	}
}

func (s *AndQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitAndQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Query_item() (localctx IQuery_itemContext) {
	return p.query_item(0)
}

func (p *QueryParser) query_item(_p int) (localctx IQuery_itemContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQuery_itemContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQuery_itemContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, QueryParserRULE_query_item, _p)

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
	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserNOT:
		localctx = NewNotQueryItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(9)
			p.Match(QueryParserNOT)
		}
		{
			p.SetState(10)
			p.query_item(3)
		}

	case QueryParserT__0:
		localctx = NewBracketQueryItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(12)
			p.query_item(0)
		}
		{
			p.SetState(13)
			p.Match(QueryParserT__1)
		}

	case QueryParserID, QueryParserSTRING, QueryParserHEADER:
		localctx = NewSingleQueryItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Single_query_item()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndQueryItemContext(p, NewQuery_itemContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_query_item)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(19)
					p.Match(QueryParserAND)
				}
				{
					p.SetState(20)
					p.query_item(6)
				}

			case 2:
				localctx = NewOrQueryItemContext(p, NewQuery_itemContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_query_item)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(22)
					p.Match(QueryParserOR)
				}
				{
					p.SetState(23)
					p.query_item(5)
				}

			}

		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ISingle_query_itemContext is an interface to support dynamic dispatch.
type ISingle_query_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingle_query_itemContext differentiates from other interfaces.
	IsSingle_query_itemContext()
}

type Single_query_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_query_itemContext() *Single_query_itemContext {
	var p = new(Single_query_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_single_query_item
	return p
}

func (*Single_query_itemContext) IsSingle_query_itemContext() {}

func NewSingle_query_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_query_itemContext {
	var p = new(Single_query_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_single_query_item

	return p
}

func (s *Single_query_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_query_itemContext) CopyFrom(ctx *Single_query_itemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Single_query_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_query_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HeaderQueryItemContext struct {
	*Single_query_itemContext
}

func NewHeaderQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HeaderQueryItemContext {
	var p = new(HeaderQueryItemContext)

	p.Single_query_itemContext = NewEmptySingle_query_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Single_query_itemContext))

	return p
}

func (s *HeaderQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderQueryItemContext) ID() antlr.TerminalNode {
	return s.GetToken(QueryParserID, 0)
}

func (s *HeaderQueryItemContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *HeaderQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterHeaderQueryItem(s)
	}
}

func (s *HeaderQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitHeaderQueryItem(s)
	}
}

func (s *HeaderQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitHeaderQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueQueryItemContext struct {
	*Single_query_itemContext
}

func NewValueQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueQueryItemContext {
	var p = new(ValueQueryItemContext)

	p.Single_query_itemContext = NewEmptySingle_query_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Single_query_itemContext))

	return p
}

func (s *ValueQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueQueryItemContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *ValueQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterValueQueryItem(s)
	}
}

func (s *ValueQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitValueQueryItem(s)
	}
}

func (s *ValueQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitValueQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type KeyQueryItemContext struct {
	*Single_query_itemContext
}

func NewKeyQueryItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeyQueryItemContext {
	var p = new(KeyQueryItemContext)

	p.Single_query_itemContext = NewEmptySingle_query_itemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Single_query_itemContext))

	return p
}

func (s *KeyQueryItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyQueryItemContext) HEADER() antlr.TerminalNode {
	return s.GetToken(QueryParserHEADER, 0)
}

func (s *KeyQueryItemContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *KeyQueryItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterKeyQueryItem(s)
	}
}

func (s *KeyQueryItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitKeyQueryItem(s)
	}
}

func (s *KeyQueryItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitKeyQueryItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Single_query_item() (localctx ISingle_query_itemContext) {
	localctx = NewSingle_query_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_single_query_item)

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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserID:
		localctx = NewHeaderQueryItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(QueryParserID)
		}
		{
			p.SetState(30)
			p.Match(QueryParserT__2)
		}
		{
			p.SetState(31)
			p.Match(QueryParserSTRING)
		}

	case QueryParserHEADER:
		localctx = NewKeyQueryItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(QueryParserHEADER)
		}
		{
			p.SetState(33)
			p.Match(QueryParserT__2)
		}
		{
			p.SetState(34)
			p.Match(QueryParserSTRING)
		}

	case QueryParserSTRING:
		localctx = NewValueQueryItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Match(QueryParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *QueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *Query_itemContext = nil
		if localctx != nil {
			t = localctx.(*Query_itemContext)
		}
		return p.Query_item_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QueryParser) Query_item_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
