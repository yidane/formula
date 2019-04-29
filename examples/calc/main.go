package main

import (
	"bufio"
	"fmt"
	"github.com/yidane/formula"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
		}

		expression := formula.NewExpression(string(l))
		result, err := expression.Evaluate()
		if err != nil {
			log.Println(err)
		}

		fmt.Println(expression.OriginalString(), "= ", result.Value)
	}
}
