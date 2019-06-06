// Code generated from SearchParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package search_parser // SearchParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SearchParser.
type SearchParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SearchParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by SearchParser#AndQueryCriterias.
	VisitAndQueryCriterias(ctx *AndQueryCriteriasContext) interface{}

	// Visit a parse tree produced by SearchParser#NotQueryCriterias.
	VisitNotQueryCriterias(ctx *NotQueryCriteriasContext) interface{}

	// Visit a parse tree produced by SearchParser#SingleQueryCriteria.
	VisitSingleQueryCriteria(ctx *SingleQueryCriteriaContext) interface{}

	// Visit a parse tree produced by SearchParser#BracketQueryCriterias.
	VisitBracketQueryCriterias(ctx *BracketQueryCriteriasContext) interface{}

	// Visit a parse tree produced by SearchParser#OrQueryCriterias.
	VisitOrQueryCriterias(ctx *OrQueryCriteriasContext) interface{}

	// Visit a parse tree produced by SearchParser#queryCriteria.
	VisitQueryCriteria(ctx *QueryCriteriaContext) interface{}

	// Visit a parse tree produced by SearchParser#headerCriteria.
	VisitHeaderCriteria(ctx *HeaderCriteriaContext) interface{}

	// Visit a parse tree produced by SearchParser#fieldCriteria.
	VisitFieldCriteria(ctx *FieldCriteriaContext) interface{}

	// Visit a parse tree produced by SearchParser#valueCriteria.
	VisitValueCriteria(ctx *ValueCriteriaContext) interface{}
}
