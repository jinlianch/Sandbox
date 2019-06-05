// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QueryParser#AndQueryCriterias.
	VisitAndQueryCriterias(ctx *AndQueryCriteriasContext) interface{}

	// Visit a parse tree produced by QueryParser#NotQueryCriterias.
	VisitNotQueryCriterias(ctx *NotQueryCriteriasContext) interface{}

	// Visit a parse tree produced by QueryParser#SingleQueryCriteria.
	VisitSingleQueryCriteria(ctx *SingleQueryCriteriaContext) interface{}

	// Visit a parse tree produced by QueryParser#BracketQueryCriterias.
	VisitBracketQueryCriterias(ctx *BracketQueryCriteriasContext) interface{}

	// Visit a parse tree produced by QueryParser#OrQueryCriterias.
	VisitOrQueryCriterias(ctx *OrQueryCriteriasContext) interface{}

	// Visit a parse tree produced by QueryParser#queryCriteria.
	VisitQueryCriteria(ctx *QueryCriteriaContext) interface{}

	// Visit a parse tree produced by QueryParser#keyCriteria.
	VisitKeyCriteria(ctx *KeyCriteriaContext) interface{}

	// Visit a parse tree produced by QueryParser#headerCriteria.
	VisitHeaderCriteria(ctx *HeaderCriteriaContext) interface{}

	// Visit a parse tree produced by QueryParser#valueCriteria.
	VisitValueCriteria(ctx *ValueCriteriaContext) interface{}
}
