// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QueryParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by QueryParser#single.
	VisitSingle(ctx *SingleContext) interface{}
}
