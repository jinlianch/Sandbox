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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 102,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 40, 10, 5, 3, 6, 3, 6, 3, 6, 5, 6, 45, 10, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 51, 10, 7, 3, 8, 3, 8, 7, 8, 55, 10, 8, 12, 8,
	14, 8, 58, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 63, 10, 9, 12, 9, 14, 9, 66,
	11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 83, 10, 10, 3, 11, 3, 11,
	3, 11, 5, 11, 88, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 6, 14, 99, 10, 14, 13, 14, 14, 14, 100, 2, 2, 15, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 2, 23, 2,
	25, 2, 27, 12, 3, 2, 9, 4, 2, 34, 34, 40, 40, 6, 2, 47, 47, 67, 92, 97,
	97, 99, 124, 7, 2, 47, 47, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 36, 36,
	94, 94, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112, 112, 116,
	116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 107, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2,
	5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 44, 3, 2,
	2, 2, 13, 50, 3, 2, 2, 2, 15, 52, 3, 2, 2, 2, 17, 59, 3, 2, 2, 2, 19, 82,
	3, 2, 2, 2, 21, 84, 3, 2, 2, 2, 23, 89, 3, 2, 2, 2, 25, 95, 3, 2, 2, 2,
	27, 98, 3, 2, 2, 2, 29, 30, 7, 42, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7,
	43, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 60, 2, 2, 34, 8, 3, 2, 2, 2, 35,
	36, 7, 99, 2, 2, 36, 37, 7, 112, 2, 2, 37, 40, 7, 102, 2, 2, 38, 40, 9,
	2, 2, 2, 39, 35, 3, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 10, 3, 2, 2, 2, 41,
	42, 7, 113, 2, 2, 42, 45, 7, 116, 2, 2, 43, 45, 7, 126, 2, 2, 44, 41, 3,
	2, 2, 2, 44, 43, 3, 2, 2, 2, 45, 12, 3, 2, 2, 2, 46, 47, 7, 112, 2, 2,
	47, 48, 7, 113, 2, 2, 48, 51, 7, 118, 2, 2, 49, 51, 7, 47, 2, 2, 50, 46,
	3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 14, 3, 2, 2, 2, 52, 56, 9, 3, 2, 2,
	53, 55, 9, 4, 2, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3,
	2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 16, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59,
	64, 7, 36, 2, 2, 60, 63, 5, 21, 11, 2, 61, 63, 10, 5, 2, 2, 62, 60, 3,
	2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64,
	65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 68, 7, 36,
	2, 2, 68, 18, 3, 2, 2, 2, 69, 70, 7, 37, 2, 2, 70, 83, 5, 17, 9, 2, 71,
	72, 7, 106, 2, 2, 72, 73, 7, 103, 2, 2, 73, 74, 7, 99, 2, 2, 74, 75, 7,
	102, 2, 2, 75, 76, 7, 103, 2, 2, 76, 77, 7, 116, 2, 2, 77, 78, 7, 93, 2,
	2, 78, 79, 3, 2, 2, 2, 79, 80, 5, 17, 9, 2, 80, 81, 7, 95, 2, 2, 81, 83,
	3, 2, 2, 2, 82, 69, 3, 2, 2, 2, 82, 71, 3, 2, 2, 2, 83, 20, 3, 2, 2, 2,
	84, 87, 7, 94, 2, 2, 85, 88, 9, 6, 2, 2, 86, 88, 5, 23, 12, 2, 87, 85,
	3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 22, 3, 2, 2, 2, 89, 90, 7, 119, 2,
	2, 90, 91, 5, 25, 13, 2, 91, 92, 5, 25, 13, 2, 92, 93, 5, 25, 13, 2, 93,
	94, 5, 25, 13, 2, 94, 24, 3, 2, 2, 2, 95, 96, 9, 7, 2, 2, 96, 26, 3, 2,
	2, 2, 97, 99, 9, 8, 2, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100,
	98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 28, 3, 2, 2, 2, 12, 2, 39, 44,
	50, 56, 62, 64, 82, 87, 100, 2,
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
	"", "", "", "", "AND", "OR", "NOT", "ID", "STRING", "HEADER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "AND", "OR", "NOT", "ID", "STRING", "HEADER", "ESC",
	"UNICODE", "HEX", "WS",
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
	QueryLexerT__1   = 2
	QueryLexerT__2   = 3
	QueryLexerAND    = 4
	QueryLexerOR     = 5
	QueryLexerNOT    = 6
	QueryLexerID     = 7
	QueryLexerSTRING = 8
	QueryLexerHEADER = 9
	QueryLexerWS     = 10
)
