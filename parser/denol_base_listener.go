// Code generated from Denol.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Denol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDenolListener is a complete listener for a parse tree produced by DenolParser.
type BaseDenolListener struct{}

var _ DenolListener = &BaseDenolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDenolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDenolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDenolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDenolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParen_expr is called when production paren_expr is entered.
func (s *BaseDenolListener) EnterParen_expr(ctx *Paren_exprContext) {}

// ExitParen_expr is called when production paren_expr is exited.
func (s *BaseDenolListener) ExitParen_expr(ctx *Paren_exprContext) {}

// EnterSqrt_expr is called when production sqrt_expr is entered.
func (s *BaseDenolListener) EnterSqrt_expr(ctx *Sqrt_exprContext) {}

// ExitSqrt_expr is called when production sqrt_expr is exited.
func (s *BaseDenolListener) ExitSqrt_expr(ctx *Sqrt_exprContext) {}

// EnterKeys_expr is called when production keys_expr is entered.
func (s *BaseDenolListener) EnterKeys_expr(ctx *Keys_exprContext) {}

// ExitKeys_expr is called when production keys_expr is exited.
func (s *BaseDenolListener) ExitKeys_expr(ctx *Keys_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseDenolListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseDenolListener) ExitExpr(ctx *ExprContext) {}

// EnterTest is called when production test is entered.
func (s *BaseDenolListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BaseDenolListener) ExitTest(ctx *TestContext) {}

// EnterSum is called when production sum is entered.
func (s *BaseDenolListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BaseDenolListener) ExitSum(ctx *SumContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseDenolListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseDenolListener) ExitTerm(ctx *TermContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDenolListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDenolListener) ExitStatement(ctx *StatementContext) {}
