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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 33, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 5, 2, 10, 10, 2, 3, 2, 3, 2, 5,
	2, 14, 10, 2, 3, 2, 5, 2, 17, 10, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22,
	11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 29, 10, 3, 3, 4, 3, 4, 3, 4,
	2, 2, 5, 2, 4, 6, 2, 3, 3, 2, 3, 5, 2, 34, 2, 9, 3, 2, 2, 2, 4, 28, 3,
	2, 2, 2, 6, 30, 3, 2, 2, 2, 8, 10, 7, 6, 2, 2, 9, 8, 3, 2, 2, 2, 9, 10,
	3, 2, 2, 2, 10, 20, 3, 2, 2, 2, 11, 14, 5, 4, 3, 2, 12, 14, 5, 6, 4, 2,
	13, 11, 3, 2, 2, 2, 13, 12, 3, 2, 2, 2, 14, 16, 3, 2, 2, 2, 15, 17, 7,
	6, 2, 2, 16, 15, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18,
	13, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2,
	2, 21, 3, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 7, 4, 2, 2, 24, 25, 7,
	3, 2, 2, 25, 29, 5, 6, 4, 2, 26, 27, 7, 4, 2, 2, 27, 29, 7, 3, 2, 2, 28,
	23, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 31, 9, 2, 2,
	2, 31, 7, 3, 2, 2, 2, 7, 9, 13, 16, 20, 28,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'",
}
var symbolicNames = []string{
	"", "", "SYMBOL", "STRING", "WS",
}

var ruleNames = []string{
	"query", "pair", "single",
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
	QueryParserSYMBOL = 2
	QueryParserSTRING = 3
	QueryParserWS     = 4
)

// QueryParser rules.
const (
	QueryParserRULE_query  = 0
	QueryParserRULE_pair   = 1
	QueryParserRULE_single = 2
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

func (s *QueryContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(QueryParserWS)
}

func (s *QueryContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(QueryParserWS, i)
}

func (s *QueryContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *QueryContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *QueryContext) AllSingle() []ISingleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleContext)(nil)).Elem())
	var tst = make([]ISingleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleContext)
		}
	}

	return tst
}

func (s *QueryContext) Single(i int) ISingleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleContext)
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
	fmt.Println("Start query")
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryParserRULE_query)
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
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QueryParserWS {
		{
			p.SetState(6)
			p.Match(QueryParserWS)
		}

	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserT__0)|(1<<QueryParserSYMBOL)|(1<<QueryParserSTRING))) != 0 {
		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(9)
				p.Pair()
			}

		case 2:
			{
				p.SetState(10)
				p.Single()
			}

		}
		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserWS {
			{
				p.SetState(13)
				p.Match(QueryParserWS)
			}

		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(QueryParserSYMBOL, 0)
}

func (s *PairContext) Single() ISingleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryParserRULE_pair)

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

	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Match(QueryParserSYMBOL)
		}
		{
			p.SetState(22)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(23)
			p.Single()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(QueryParserSYMBOL)
		}
		{
			p.SetState(25)
			p.Match(QueryParserT__0)
		}

	}

	return localctx
}

// ISingleContext is an interface to support dynamic dispatch.
type ISingleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleContext differentiates from other interfaces.
	IsSingleContext()
}

type SingleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleContext() *SingleContext {
	var p = new(SingleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_single
	return p
}

func (*SingleContext) IsSingleContext() {}

func NewSingleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleContext {
	var p = new(SingleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_single

	return p
}

func (s *SingleContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(QueryParserSYMBOL, 0)
}

func (s *SingleContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *SingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterSingle(s)
	}
}

func (s *SingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitSingle(s)
	}
}

func (s *SingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Single() (localctx ISingleContext) {
	localctx = NewSingleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_single)
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
		p.SetState(28)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserT__0)|(1<<QueryParserSYMBOL)|(1<<QueryParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
