package main

import (
	"bufio"
	"fmt"
	"github.com/yidane/formula"
	"io"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
		}

		if len(l) == 0 {
			continue
		}

		s := string(l)
		fmt.Println(s)

		expression := formula.NewExpression(s)
		result, err := expression.Evaluate()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(expression.OriginalString(), "= ", result.Value)
		}
	}
}
