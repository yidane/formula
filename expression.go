package formula

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
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

func (expression *Expression) Evaluate() (interface{}, error) {
	if expression.context.Error != nil {
		return nil, expression.context.Error
	}

	err := expression.Compile()
	if err != nil {
		return nil, err
	}

	result := (*expression.parsedExpression).Evaluate()
	if expression.context.Error != nil {
		return nil, expression.context.Error
	}

	return result.Value, nil
}
