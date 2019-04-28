package formula

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	_ "github.com/yidane/formula/internal/fs"
	"github.com/yidane/formula/internal/parser"
	"github.com/yidane/formula/opt"
	"strings"
)

type Expression struct {
	context            *opt.FormulaContext
	originalExpression string
	parsedExpression   *opt.LogicalExpression
}

func NewExpression(expression string, options ...opt.Option) *Expression {
	return &Expression{
		originalExpression: strings.TrimSpace(expression),
		context:            opt.NewFormulaContext(options...),
	}
}

func (expression *Expression) Compile() error {
	lexer := parser.NewFormulaLexer(antlr.NewInputStream(expression.originalExpression))
	formulaParser := parser.NewFormulaParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	calcContext := formulaParser.Calc()
	//formulaParser.AddErrorListener(antlr.ErrorListener())

	//编译过程中，怎么解决异常？
	//if calcContext!=nil{
	//
	//}

	expression.parsedExpression = calcContext.GetRetValue()

	return nil
}

func (expression *Expression) Evaluate() (*opt.Argument, error) {
	err := expression.Compile()
	if err != nil {
		return nil, err
	}

	result, err := (*expression.parsedExpression).Evaluate(expression.context)
	if err != nil {
		return nil, err
	}

	return result, nil
}
