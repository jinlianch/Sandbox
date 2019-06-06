// Code generated from SearchParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package search_parser // SearchParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SearchParserListener is a complete listener for a parse tree produced by SearchParser.
type SearchParserListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterAndQueryCriterias is called when entering the AndQueryCriterias production.
	EnterAndQueryCriterias(c *AndQueryCriteriasContext)

	// EnterNotQueryCriterias is called when entering the NotQueryCriterias production.
	EnterNotQueryCriterias(c *NotQueryCriteriasContext)

	// EnterSingleQueryCriteria is called when entering the SingleQueryCriteria production.
	EnterSingleQueryCriteria(c *SingleQueryCriteriaContext)

	// EnterBracketQueryCriterias is called when entering the BracketQueryCriterias production.
	EnterBracketQueryCriterias(c *BracketQueryCriteriasContext)

	// EnterOrQueryCriterias is called when entering the OrQueryCriterias production.
	EnterOrQueryCriterias(c *OrQueryCriteriasContext)

	// EnterQueryCriteria is called when entering the queryCriteria production.
	EnterQueryCriteria(c *QueryCriteriaContext)

	// EnterHeaderCriteria is called when entering the headerCriteria production.
	EnterHeaderCriteria(c *HeaderCriteriaContext)

	// EnterFieldCriteria is called when entering the fieldCriteria production.
	EnterFieldCriteria(c *FieldCriteriaContext)

	// EnterValueCriteria is called when entering the valueCriteria production.
	EnterValueCriteria(c *ValueCriteriaContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitAndQueryCriterias is called when exiting the AndQueryCriterias production.
	ExitAndQueryCriterias(c *AndQueryCriteriasContext)

	// ExitNotQueryCriterias is called when exiting the NotQueryCriterias production.
	ExitNotQueryCriterias(c *NotQueryCriteriasContext)

	// ExitSingleQueryCriteria is called when exiting the SingleQueryCriteria production.
	ExitSingleQueryCriteria(c *SingleQueryCriteriaContext)

	// ExitBracketQueryCriterias is called when exiting the BracketQueryCriterias production.
	ExitBracketQueryCriterias(c *BracketQueryCriteriasContext)

	// ExitOrQueryCriterias is called when exiting the OrQueryCriterias production.
	ExitOrQueryCriterias(c *OrQueryCriteriasContext)

	// ExitQueryCriteria is called when exiting the queryCriteria production.
	ExitQueryCriteria(c *QueryCriteriaContext)

	// ExitHeaderCriteria is called when exiting the headerCriteria production.
	ExitHeaderCriteria(c *HeaderCriteriaContext)

	// ExitFieldCriteria is called when exiting the fieldCriteria production.
	ExitFieldCriteria(c *FieldCriteriaContext)

	// ExitValueCriteria is called when exiting the valueCriteria production.
	ExitValueCriteria(c *ValueCriteriaContext)
}
