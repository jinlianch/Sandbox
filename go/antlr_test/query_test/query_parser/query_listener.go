// Code generated from Query.g4 by ANTLR 4.7.2. DO NOT EDIT.

package query_parser // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterSingle is called when entering the single production.
	EnterSingle(c *SingleContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitSingle is called when exiting the single production.
	ExitSingle(c *SingleContext)
}
