// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterBracketQueryItem is called when entering the BracketQueryItem production.
	EnterBracketQueryItem(c *BracketQueryItemContext)

	// EnterSingleQueryItem is called when entering the SingleQueryItem production.
	EnterSingleQueryItem(c *SingleQueryItemContext)

	// EnterNotQueryItem is called when entering the NotQueryItem production.
	EnterNotQueryItem(c *NotQueryItemContext)

	// EnterOrQueryItem is called when entering the OrQueryItem production.
	EnterOrQueryItem(c *OrQueryItemContext)

	// EnterAndQueryItem is called when entering the AndQueryItem production.
	EnterAndQueryItem(c *AndQueryItemContext)

	// EnterHeaderQueryItem is called when entering the HeaderQueryItem production.
	EnterHeaderQueryItem(c *HeaderQueryItemContext)

	// EnterKeyQueryItem is called when entering the KeyQueryItem production.
	EnterKeyQueryItem(c *KeyQueryItemContext)

	// EnterValueQueryItem is called when entering the ValueQueryItem production.
	EnterValueQueryItem(c *ValueQueryItemContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitBracketQueryItem is called when exiting the BracketQueryItem production.
	ExitBracketQueryItem(c *BracketQueryItemContext)

	// ExitSingleQueryItem is called when exiting the SingleQueryItem production.
	ExitSingleQueryItem(c *SingleQueryItemContext)

	// ExitNotQueryItem is called when exiting the NotQueryItem production.
	ExitNotQueryItem(c *NotQueryItemContext)

	// ExitOrQueryItem is called when exiting the OrQueryItem production.
	ExitOrQueryItem(c *OrQueryItemContext)

	// ExitAndQueryItem is called when exiting the AndQueryItem production.
	ExitAndQueryItem(c *AndQueryItemContext)

	// ExitHeaderQueryItem is called when exiting the HeaderQueryItem production.
	ExitHeaderQueryItem(c *HeaderQueryItemContext)

	// ExitKeyQueryItem is called when exiting the KeyQueryItem production.
	ExitKeyQueryItem(c *KeyQueryItemContext)

	// ExitValueQueryItem is called when exiting the ValueQueryItem production.
	ExitValueQueryItem(c *ValueQueryItemContext)
}
