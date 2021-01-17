// Code generated from Denol.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Denol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DenolListener is a complete listener for a parse tree produced by DenolParser.
type DenolListener interface {
	antlr.ParseTreeListener

	// EnterParen_expr is called when entering the paren_expr production.
	EnterParen_expr(c *Paren_exprContext)

	// EnterSqrt_expr is called when entering the sqrt_expr production.
	EnterSqrt_expr(c *Sqrt_exprContext)

	// EnterKeys_expr is called when entering the keys_expr production.
	EnterKeys_expr(c *Keys_exprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// ExitParen_expr is called when exiting the paren_expr production.
	ExitParen_expr(c *Paren_exprContext)

	// ExitSqrt_expr is called when exiting the sqrt_expr production.
	ExitSqrt_expr(c *Sqrt_exprContext)

	// ExitKeys_expr is called when exiting the keys_expr production.
	ExitKeys_expr(c *Keys_exprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)
}
