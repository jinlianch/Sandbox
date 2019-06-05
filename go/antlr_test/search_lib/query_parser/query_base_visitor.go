// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitAndQueryCriterias(ctx *AndQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitNotQueryCriterias(ctx *NotQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitSingleQueryCriteria(ctx *SingleQueryCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitBracketQueryCriterias(ctx *BracketQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitOrQueryCriterias(ctx *OrQueryCriteriasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitQueryCriteria(ctx *QueryCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitKeyCriteria(ctx *KeyCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitHeaderCriteria(ctx *HeaderCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitValueCriteria(ctx *ValueCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}
