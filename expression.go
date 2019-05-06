package formula

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yidane/formula/internal/cache"
	"github.com/yidane/formula/internal/exp"
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

func (expression *Expression) compile() error {
	//handle empty expression
	if expression.originalExpression == "" {
		expression.parsedExpression = exp.NewEmptyExpression()
		return nil
	}

	//restore expression from cache
	logicExpression := cache.Restore(expression.context.Option, expression.originalExpression)
	if logicExpression != nil {
		expression.parsedExpression = logicExpression
		return nil
	}

	//compile expression
	lexer := parser.NewFormulaLexer(antlr.NewInputStream(expression.originalExpression))
	formulaParser := parser.NewFormulaParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	//handle compile error
	errListener := newFormulaErrorListener()
	formulaParser.AddErrorListener(errListener)
	calcContext := formulaParser.Calc()

	if errListener.HasError() {
		return errListener.Error()
	}

	expression.parsedExpression = calcContext.GetRetValue()
	return nil
}

func (expression *Expression) OriginalString() string {
	return expression.originalExpression
}

func (expression *Expression) AddParameter(name string, value interface{}) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("argument name can not be empty")
	}

	if _, ok := expression.context.Parameters[name]; ok {
		return fmt.Errorf("argument %s dureplate", name)
	}

	expression.context.Parameters[name] = value
	return nil
}

func (expression *Expression) ResetParameters() {
	expression.context.Parameters = make(map[string]interface{})
}

func (expression *Expression) Evaluate() (*opt.Argument, error) {
	err := expression.compile()
	if err != nil {
		return nil, err
	}

	result, err := (*expression.parsedExpression).Evaluate(expression.context)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type formulaErrorListener struct {
	buf bytes.Buffer
}

func newFormulaErrorListener() *formulaErrorListener {
	return new(formulaErrorListener)
}

func (l *formulaErrorListener) HasError() bool {
	return l.buf.Len() > 0
}

func (l *formulaErrorListener) Error() error {
	return errors.New(l.buf.String())
}

func (l *formulaErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.buf.WriteString(msg)
}

func (*formulaErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (*formulaErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (*formulaErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}
