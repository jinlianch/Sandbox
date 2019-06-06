// Code generated from SearchParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package search_parser // SearchParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSearchParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSearchParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitAndQueryCriterias(ctx *AndQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitNotQueryCriterias(ctx *NotQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitSingleQueryCriteria(ctx *SingleQueryCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitBracketQueryCriterias(ctx *BracketQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitOrQueryCriterias(ctx *OrQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitQueryCriteria(ctx *QueryCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitHeaderCriteria(ctx *HeaderCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitFieldCriteria(ctx *FieldCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchParserVisitor) VisitValueCriteria(ctx *ValueCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}
