// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QueryParser#BracketQueryItem.
	VisitBracketQueryItem(ctx *BracketQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#SingleQueryItem.
	VisitSingleQueryItem(ctx *SingleQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#NotQueryItem.
	VisitNotQueryItem(ctx *NotQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#OrQueryItem.
	VisitOrQueryItem(ctx *OrQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#AndQueryItem.
	VisitAndQueryItem(ctx *AndQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#HeaderQueryItem.
	VisitHeaderQueryItem(ctx *HeaderQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#KeyQueryItem.
	VisitKeyQueryItem(ctx *KeyQueryItemContext) interface{}

	// Visit a parse tree produced by QueryParser#ValueQueryItem.
	VisitValueQueryItem(ctx *ValueQueryItemContext) interface{}
}
