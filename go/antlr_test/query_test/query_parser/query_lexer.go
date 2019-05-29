// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 52, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 6, 3, 21, 10, 3, 13, 3, 14, 3, 22, 3,
	4, 3, 4, 3, 4, 7, 4, 28, 10, 4, 12, 4, 14, 4, 31, 11, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 5, 5, 38, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 6, 8, 49, 10, 8, 13, 8, 14, 8, 50, 2, 2, 9, 3, 3, 5, 4,
	7, 5, 9, 2, 11, 2, 13, 2, 15, 6, 3, 2, 7, 7, 2, 11, 12, 15, 15, 34, 34,
	36, 36, 60, 60, 4, 2, 36, 36, 94, 94, 10, 2, 36, 36, 49, 49, 94, 94, 100,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99,
	104, 5, 2, 11, 12, 15, 15, 34, 34, 2, 53, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 20, 3,
	2, 2, 2, 7, 24, 3, 2, 2, 2, 9, 34, 3, 2, 2, 2, 11, 39, 3, 2, 2, 2, 13,
	45, 3, 2, 2, 2, 15, 48, 3, 2, 2, 2, 17, 18, 7, 60, 2, 2, 18, 4, 3, 2, 2,
	2, 19, 21, 10, 2, 2, 2, 20, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 20,
	3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 6, 3, 2, 2, 2, 24, 29, 7, 36, 2, 2,
	25, 28, 5, 9, 5, 2, 26, 28, 10, 3, 2, 2, 27, 25, 3, 2, 2, 2, 27, 26, 3,
	2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30,
	32, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 7, 36, 2, 2, 33, 8, 3, 2, 2,
	2, 34, 37, 7, 94, 2, 2, 35, 38, 9, 4, 2, 2, 36, 38, 5, 11, 6, 2, 37, 35,
	3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 10, 3, 2, 2, 2, 39, 40, 7, 119, 2,
	2, 40, 41, 5, 13, 7, 2, 41, 42, 5, 13, 7, 2, 42, 43, 5, 13, 7, 2, 43, 44,
	5, 13, 7, 2, 44, 12, 3, 2, 2, 2, 45, 46, 9, 5, 2, 2, 46, 14, 3, 2, 2, 2,
	47, 49, 9, 6, 2, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3,
	2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 16, 3, 2, 2, 2, 8, 2, 22, 27, 29, 37,
	50, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "SYMBOL", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "SYMBOL", "STRING", "ESC", "UNICODE", "HEX", "WS",
}

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewQueryLexer(input antlr.CharStream) *QueryLexer {

	l := new(QueryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerT__0   = 1
	QueryLexerSYMBOL = 2
	QueryLexerSTRING = 3
	QueryLexerWS     = 4
)
