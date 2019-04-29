// Generated from Formula.g4 by ANTLR 4.7.

package parser // Formula

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FormulaListener is a complete listener for a parse tree produced by FormulaParser.
type FormulaListener interface {
	antlr.ParseTreeListener

	// EnterCalc is called when entering the calc production.
	EnterCalc(c *CalcContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterBitOrExpr is called when entering the bitOrExpr production.
	EnterBitOrExpr(c *BitOrExprContext)

	// EnterBitXorExpr is called when entering the bitXorExpr production.
	EnterBitXorExpr(c *BitXorExprContext)

	// EnterBitAndExpr is called when entering the bitAndExpr production.
	EnterBitAndExpr(c *BitAndExprContext)

	// EnterEqExpr is called when entering the eqExpr production.
	EnterEqExpr(c *EqExprContext)

	// EnterRelExpr is called when entering the relExpr production.
	EnterRelExpr(c *RelExprContext)

	// EnterShiftExpr is called when entering the shiftExpr production.
	EnterShiftExpr(c *ShiftExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterMultExpr is called when entering the multExpr production.
	EnterMultExpr(c *MultExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterΠ is called when entering the π production.
	EnterΠ(c *ΠContext)

	// ExitCalc is called when exiting the calc production.
	ExitCalc(c *CalcContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitBitOrExpr is called when exiting the bitOrExpr production.
	ExitBitOrExpr(c *BitOrExprContext)

	// ExitBitXorExpr is called when exiting the bitXorExpr production.
	ExitBitXorExpr(c *BitXorExprContext)

	// ExitBitAndExpr is called when exiting the bitAndExpr production.
	ExitBitAndExpr(c *BitAndExprContext)

	// ExitEqExpr is called when exiting the eqExpr production.
	ExitEqExpr(c *EqExprContext)

	// ExitRelExpr is called when exiting the relExpr production.
	ExitRelExpr(c *RelExprContext)

	// ExitShiftExpr is called when exiting the shiftExpr production.
	ExitShiftExpr(c *ShiftExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitMultExpr is called when exiting the multExpr production.
	ExitMultExpr(c *MultExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitΠ is called when exiting the π production.
	ExitΠ(c *ΠContext)
}
