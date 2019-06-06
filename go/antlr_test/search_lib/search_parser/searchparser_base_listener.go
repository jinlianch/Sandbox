// Code generated from SearchParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package search_parser // SearchParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSearchParserListener is a complete listener for a parse tree produced by SearchParser.
type BaseSearchParserListener struct{}

var _ SearchParserListener = &BaseSearchParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSearchParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSearchParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSearchParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSearchParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSearchParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSearchParserListener) ExitQuery(ctx *QueryContext) {}

// EnterAndQueryCriterias is called when production AndQueryCriterias is entered.
func (s *BaseSearchParserListener) EnterAndQueryCriterias(ctx *AndQueryCriteriasContext) {}

// ExitAndQueryCriterias is called when production AndQueryCriterias is exited.
func (s *BaseSearchParserListener) ExitAndQueryCriterias(ctx *AndQueryCriteriasContext) {}

// EnterNotQueryCriterias is called when production NotQueryCriterias is entered.
func (s *BaseSearchParserListener) EnterNotQueryCriterias(ctx *NotQueryCriteriasContext) {}

// ExitNotQueryCriterias is called when production NotQueryCriterias is exited.
func (s *BaseSearchParserListener) ExitNotQueryCriterias(ctx *NotQueryCriteriasContext) {}

// EnterSingleQueryCriteria is called when production SingleQueryCriteria is entered.
func (s *BaseSearchParserListener) EnterSingleQueryCriteria(ctx *SingleQueryCriteriaContext) {}

// ExitSingleQueryCriteria is called when production SingleQueryCriteria is exited.
func (s *BaseSearchParserListener) ExitSingleQueryCriteria(ctx *SingleQueryCriteriaContext) {}

// EnterBracketQueryCriterias is called when production BracketQueryCriterias is entered.
func (s *BaseSearchParserListener) EnterBracketQueryCriterias(ctx *BracketQueryCriteriasContext) {}

// ExitBracketQueryCriterias is called when production BracketQueryCriterias is exited.
func (s *BaseSearchParserListener) ExitBracketQueryCriterias(ctx *BracketQueryCriteriasContext) {}

// EnterOrQueryCriterias is called when production OrQueryCriterias is entered.
func (s *BaseSearchParserListener) EnterOrQueryCriterias(ctx *OrQueryCriteriasContext) {}

// ExitOrQueryCriterias is called when production OrQueryCriterias is exited.
func (s *BaseSearchParserListener) ExitOrQueryCriterias(ctx *OrQueryCriteriasContext) {}

// EnterQueryCriteria is called when production queryCriteria is entered.
func (s *BaseSearchParserListener) EnterQueryCriteria(ctx *QueryCriteriaContext) {}

// ExitQueryCriteria is called when production queryCriteria is exited.
func (s *BaseSearchParserListener) ExitQueryCriteria(ctx *QueryCriteriaContext) {}

// EnterHeaderCriteria is called when production headerCriteria is entered.
func (s *BaseSearchParserListener) EnterHeaderCriteria(ctx *HeaderCriteriaContext) {}

// ExitHeaderCriteria is called when production headerCriteria is exited.
func (s *BaseSearchParserListener) ExitHeaderCriteria(ctx *HeaderCriteriaContext) {}

// EnterFieldCriteria is called when production fieldCriteria is entered.
func (s *BaseSearchParserListener) EnterFieldCriteria(ctx *FieldCriteriaContext) {}

// ExitFieldCriteria is called when production fieldCriteria is exited.
func (s *BaseSearchParserListener) ExitFieldCriteria(ctx *FieldCriteriaContext) {}

// EnterValueCriteria is called when production valueCriteria is entered.
func (s *BaseSearchParserListener) EnterValueCriteria(ctx *ValueCriteriaContext) {}

// ExitValueCriteria is called when production valueCriteria is exited.
func (s *BaseSearchParserListener) ExitValueCriteria(ctx *ValueCriteriaContext) {}
