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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 109,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 5, 5, 41, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6,
	47, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 53, 10, 7, 3, 8, 6, 8, 56, 10,
	8, 13, 8, 14, 8, 57, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 73, 10, 9, 3, 10, 6, 10, 76, 10, 10, 13,
	10, 14, 10, 77, 3, 11, 3, 11, 3, 11, 7, 11, 83, 10, 11, 12, 11, 14, 11,
	86, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 93, 10, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 6, 15, 104, 10,
	15, 13, 15, 14, 15, 105, 3, 15, 3, 15, 3, 84, 2, 16, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 2, 25, 2, 27, 2,
	29, 13, 3, 2, 7, 7, 2, 11, 12, 15, 15, 34, 34, 36, 36, 60, 60, 6, 2, 11,
	12, 15, 15, 34, 34, 36, 36, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104,
	104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 115, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 3, 31, 3, 2, 2, 2, 5, 33, 3, 2, 2, 2, 7, 35, 3, 2,
	2, 2, 9, 40, 3, 2, 2, 2, 11, 46, 3, 2, 2, 2, 13, 52, 3, 2, 2, 2, 15, 55,
	3, 2, 2, 2, 17, 72, 3, 2, 2, 2, 19, 75, 3, 2, 2, 2, 21, 79, 3, 2, 2, 2,
	23, 89, 3, 2, 2, 2, 25, 94, 3, 2, 2, 2, 27, 100, 3, 2, 2, 2, 29, 103, 3,
	2, 2, 2, 31, 32, 7, 42, 2, 2, 32, 4, 3, 2, 2, 2, 33, 34, 7, 43, 2, 2, 34,
	6, 3, 2, 2, 2, 35, 36, 7, 60, 2, 2, 36, 8, 3, 2, 2, 2, 37, 38, 7, 113,
	2, 2, 38, 41, 7, 116, 2, 2, 39, 41, 7, 126, 2, 2, 40, 37, 3, 2, 2, 2, 40,
	39, 3, 2, 2, 2, 41, 10, 3, 2, 2, 2, 42, 43, 7, 112, 2, 2, 43, 44, 7, 113,
	2, 2, 44, 47, 7, 118, 2, 2, 45, 47, 7, 47, 2, 2, 46, 42, 3, 2, 2, 2, 46,
	45, 3, 2, 2, 2, 47, 12, 3, 2, 2, 2, 48, 49, 7, 99, 2, 2, 49, 50, 7, 112,
	2, 2, 50, 53, 7, 102, 2, 2, 51, 53, 7, 40, 2, 2, 52, 48, 3, 2, 2, 2, 52,
	51, 3, 2, 2, 2, 53, 14, 3, 2, 2, 2, 54, 56, 10, 2, 2, 2, 55, 54, 3, 2,
	2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 16,
	3, 2, 2, 2, 59, 60, 7, 37, 2, 2, 60, 73, 5, 15, 8, 2, 61, 62, 7, 106, 2,
	2, 62, 63, 7, 103, 2, 2, 63, 64, 7, 99, 2, 2, 64, 65, 7, 102, 2, 2, 65,
	66, 7, 103, 2, 2, 66, 67, 7, 116, 2, 2, 67, 68, 7, 93, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 70, 5, 15, 8, 2, 70, 71, 7, 95, 2, 2, 71, 73, 3, 2, 2, 2,
	72, 59, 3, 2, 2, 2, 72, 61, 3, 2, 2, 2, 73, 18, 3, 2, 2, 2, 74, 76, 10,
	3, 2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77,
	78, 3, 2, 2, 2, 78, 20, 3, 2, 2, 2, 79, 84, 7, 36, 2, 2, 80, 83, 5, 23,
	12, 2, 81, 83, 11, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83,
	86, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 87, 3, 2, 2,
	2, 86, 84, 3, 2, 2, 2, 87, 88, 7, 36, 2, 2, 88, 22, 3, 2, 2, 2, 89, 92,
	7, 94, 2, 2, 90, 93, 9, 4, 2, 2, 91, 93, 5, 25, 13, 2, 92, 90, 3, 2, 2,
	2, 92, 91, 3, 2, 2, 2, 93, 24, 3, 2, 2, 2, 94, 95, 7, 119, 2, 2, 95, 96,
	5, 27, 14, 2, 96, 97, 5, 27, 14, 2, 97, 98, 5, 27, 14, 2, 98, 99, 5, 27,
	14, 2, 99, 26, 3, 2, 2, 2, 100, 101, 9, 5, 2, 2, 101, 28, 3, 2, 2, 2, 102,
	104, 9, 6, 2, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 8, 15,
	2, 2, 108, 30, 3, 2, 2, 2, 13, 2, 40, 46, 52, 57, 72, 77, 82, 84, 92, 105,
	3, 8, 2, 2,
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
	"", "'('", "')'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "OR", "NOT", "AND", "KEY", "HEADER", "NOQUOTE_STRING",
	"QUOTE_STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "OR", "NOT", "AND", "KEY", "HEADER", "NOQUOTE_STRING",
	"QUOTE_STRING", "ESC", "UNICODE", "HEX", "WS",
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
	QueryLexerT__0           = 1
	QueryLexerT__1           = 2
	QueryLexerT__2           = 3
	QueryLexerOR             = 4
	QueryLexerNOT            = 5
	QueryLexerAND            = 6
	QueryLexerKEY            = 7
	QueryLexerHEADER         = 8
	QueryLexerNOQUOTE_STRING = 9
	QueryLexerQUOTE_STRING   = 10
	QueryLexerWS             = 11
)
