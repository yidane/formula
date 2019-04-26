// Generated from Formula.g4 by ANTLR 4.7.

package parser // Formula

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFormulaListener is a complete listener for a parse tree produced by FormulaParser.
type BaseFormulaListener struct{}

var _ FormulaListener = &BaseFormulaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFormulaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFormulaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFormulaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFormulaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNcalc is called when production ncalc is entered.
func (s *BaseFormulaListener) EnterNcalc(ctx *NcalcContext) {}

// ExitNcalc is called when production ncalc is exited.
func (s *BaseFormulaListener) ExitNcalc(ctx *NcalcContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseFormulaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseFormulaListener) ExitExpr(ctx *ExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseFormulaListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseFormulaListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseFormulaListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseFormulaListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterBitOrExpr is called when production bitOrExpr is entered.
func (s *BaseFormulaListener) EnterBitOrExpr(ctx *BitOrExprContext) {}

// ExitBitOrExpr is called when production bitOrExpr is exited.
func (s *BaseFormulaListener) ExitBitOrExpr(ctx *BitOrExprContext) {}

// EnterBitXorExpr is called when production bitXorExpr is entered.
func (s *BaseFormulaListener) EnterBitXorExpr(ctx *BitXorExprContext) {}

// ExitBitXorExpr is called when production bitXorExpr is exited.
func (s *BaseFormulaListener) ExitBitXorExpr(ctx *BitXorExprContext) {}

// EnterBitAndExpr is called when production bitAndExpr is entered.
func (s *BaseFormulaListener) EnterBitAndExpr(ctx *BitAndExprContext) {}

// ExitBitAndExpr is called when production bitAndExpr is exited.
func (s *BaseFormulaListener) ExitBitAndExpr(ctx *BitAndExprContext) {}

// EnterEqExpr is called when production eqExpr is entered.
func (s *BaseFormulaListener) EnterEqExpr(ctx *EqExprContext) {}

// ExitEqExpr is called when production eqExpr is exited.
func (s *BaseFormulaListener) ExitEqExpr(ctx *EqExprContext) {}

// EnterRelExpr is called when production relExpr is entered.
func (s *BaseFormulaListener) EnterRelExpr(ctx *RelExprContext) {}

// ExitRelExpr is called when production relExpr is exited.
func (s *BaseFormulaListener) ExitRelExpr(ctx *RelExprContext) {}

// EnterShiftExpr is called when production shiftExpr is entered.
func (s *BaseFormulaListener) EnterShiftExpr(ctx *ShiftExprContext) {}

// ExitShiftExpr is called when production shiftExpr is exited.
func (s *BaseFormulaListener) ExitShiftExpr(ctx *ShiftExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseFormulaListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseFormulaListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterMultExpr is called when production multExpr is entered.
func (s *BaseFormulaListener) EnterMultExpr(ctx *MultExprContext) {}

// ExitMultExpr is called when production multExpr is exited.
func (s *BaseFormulaListener) ExitMultExpr(ctx *MultExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseFormulaListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseFormulaListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseFormulaListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseFormulaListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFormulaListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFormulaListener) ExitValue(ctx *ValueContext) {}

// EnterId is called when production id is entered.
func (s *BaseFormulaListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseFormulaListener) ExitId(ctx *IdContext) {}
