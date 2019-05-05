package main

import (
	"github.com/yidane/formula"
	"github.com/yidane/formula/opt"
	"log"
)

func init() {
	var f opt.Function = new(CustomFunction)
	err := formula.Register(&f)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	expression := formula.NewExpression("CustomFunction(1,2)")
	result, err := expression.Evaluate()
	if err != nil {
		log.Fatal(err)
	}

	v, err := result.Int64()
	if err != nil {
		log.Fatal(err)
	}

	if v != 4 { //CustomFunction: i+j+1
		log.Fatal("error")
	}

	log.Println("custom function succeed")
}
