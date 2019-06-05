// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterAndQueryCriterias is called when production AndQueryCriterias is entered.
func (s *BaseQueryListener) EnterAndQueryCriterias(ctx *AndQueryCriteriasContext) {}

// ExitAndQueryCriterias is called when production AndQueryCriterias is exited.
func (s *BaseQueryListener) ExitAndQueryCriterias(ctx *AndQueryCriteriasContext) {}

// EnterNotQueryCriterias is called when production NotQueryCriterias is entered.
func (s *BaseQueryListener) EnterNotQueryCriterias(ctx *NotQueryCriteriasContext) {}

// ExitNotQueryCriterias is called when production NotQueryCriterias is exited.
func (s *BaseQueryListener) ExitNotQueryCriterias(ctx *NotQueryCriteriasContext) {}

// EnterSingleQueryCriteria is called when production SingleQueryCriteria is entered.
func (s *BaseQueryListener) EnterSingleQueryCriteria(ctx *SingleQueryCriteriaContext) {}

// ExitSingleQueryCriteria is called when production SingleQueryCriteria is exited.
func (s *BaseQueryListener) ExitSingleQueryCriteria(ctx *SingleQueryCriteriaContext) {}

// EnterBracketQueryCriterias is called when production BracketQueryCriterias is entered.
func (s *BaseQueryListener) EnterBracketQueryCriterias(ctx *BracketQueryCriteriasContext) {}

// ExitBracketQueryCriterias is called when production BracketQueryCriterias is exited.
func (s *BaseQueryListener) ExitBracketQueryCriterias(ctx *BracketQueryCriteriasContext) {}

// EnterOrQueryCriterias is called when production OrQueryCriterias is entered.
func (s *BaseQueryListener) EnterOrQueryCriterias(ctx *OrQueryCriteriasContext) {}

// ExitOrQueryCriterias is called when production OrQueryCriterias is exited.
func (s *BaseQueryListener) ExitOrQueryCriterias(ctx *OrQueryCriteriasContext) {}

// EnterQueryCriteria is called when production queryCriteria is entered.
func (s *BaseQueryListener) EnterQueryCriteria(ctx *QueryCriteriaContext) {}

// ExitQueryCriteria is called when production queryCriteria is exited.
func (s *BaseQueryListener) ExitQueryCriteria(ctx *QueryCriteriaContext) {}

// EnterKeyCriteria is called when production keyCriteria is entered.
func (s *BaseQueryListener) EnterKeyCriteria(ctx *KeyCriteriaContext) {}

// ExitKeyCriteria is called when production keyCriteria is exited.
func (s *BaseQueryListener) ExitKeyCriteria(ctx *KeyCriteriaContext) {}

// EnterHeaderCriteria is called when production headerCriteria is entered.
func (s *BaseQueryListener) EnterHeaderCriteria(ctx *HeaderCriteriaContext) {}

// ExitHeaderCriteria is called when production headerCriteria is exited.
func (s *BaseQueryListener) ExitHeaderCriteria(ctx *HeaderCriteriaContext) {}

// EnterValueCriteria is called when production valueCriteria is entered.
func (s *BaseQueryListener) EnterValueCriteria(ctx *ValueCriteriaContext) {}

// ExitValueCriteria is called when production valueCriteria is exited.
func (s *BaseQueryListener) ExitValueCriteria(ctx *ValueCriteriaContext) {}
