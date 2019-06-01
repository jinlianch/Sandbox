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

// EnterBracketQueryItem is called when production BracketQueryItem is entered.
func (s *BaseQueryListener) EnterBracketQueryItem(ctx *BracketQueryItemContext) {}

// ExitBracketQueryItem is called when production BracketQueryItem is exited.
func (s *BaseQueryListener) ExitBracketQueryItem(ctx *BracketQueryItemContext) {}

// EnterSingleQueryItem is called when production SingleQueryItem is entered.
func (s *BaseQueryListener) EnterSingleQueryItem(ctx *SingleQueryItemContext) {}

// ExitSingleQueryItem is called when production SingleQueryItem is exited.
func (s *BaseQueryListener) ExitSingleQueryItem(ctx *SingleQueryItemContext) {}

// EnterNotQueryItem is called when production NotQueryItem is entered.
func (s *BaseQueryListener) EnterNotQueryItem(ctx *NotQueryItemContext) {}

// ExitNotQueryItem is called when production NotQueryItem is exited.
func (s *BaseQueryListener) ExitNotQueryItem(ctx *NotQueryItemContext) {}

// EnterOrQueryItem is called when production OrQueryItem is entered.
func (s *BaseQueryListener) EnterOrQueryItem(ctx *OrQueryItemContext) {}

// ExitOrQueryItem is called when production OrQueryItem is exited.
func (s *BaseQueryListener) ExitOrQueryItem(ctx *OrQueryItemContext) {}

// EnterAndQueryItem is called when production AndQueryItem is entered.
func (s *BaseQueryListener) EnterAndQueryItem(ctx *AndQueryItemContext) {}

// ExitAndQueryItem is called when production AndQueryItem is exited.
func (s *BaseQueryListener) ExitAndQueryItem(ctx *AndQueryItemContext) {}

// EnterHeaderQueryItem is called when production HeaderQueryItem is entered.
func (s *BaseQueryListener) EnterHeaderQueryItem(ctx *HeaderQueryItemContext) {}

// ExitHeaderQueryItem is called when production HeaderQueryItem is exited.
func (s *BaseQueryListener) ExitHeaderQueryItem(ctx *HeaderQueryItemContext) {}

// EnterKeyQueryItem is called when production KeyQueryItem is entered.
func (s *BaseQueryListener) EnterKeyQueryItem(ctx *KeyQueryItemContext) {}

// ExitKeyQueryItem is called when production KeyQueryItem is exited.
func (s *BaseQueryListener) ExitKeyQueryItem(ctx *KeyQueryItemContext) {}

// EnterValueQueryItem is called when production ValueQueryItem is entered.
func (s *BaseQueryListener) EnterValueQueryItem(ctx *ValueQueryItemContext) {}

// ExitValueQueryItem is called when production ValueQueryItem is exited.
func (s *BaseQueryListener) ExitValueQueryItem(ctx *ValueQueryItemContext) {}
