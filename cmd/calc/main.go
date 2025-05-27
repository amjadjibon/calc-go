package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amjadjibon/calc-go"
)

func formatResult(result float64) string {
	str := fmt.Sprintf("%.10f", result)
	str = strings.TrimRight(str, "0")
	str = strings.TrimSuffix(str, ".")
	if str == "" {
		return "0"
	}
	return str
}

func safeEvalPrint(expr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
		}
	}()
	root := calc.NewParser(strings.TrimSpace(expr)).Parse()
	result := root.Eval()
	fmt.Println(formatResult(result))
}

func main() {
	if len(os.Args) > 1 {
		safeEvalPrint(strings.Join(os.Args[1:], " "))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter expressions or 'q'/'exit' to quit.")
	for {
		fmt.Print("» ")
		if !scanner.Scan() {
			break
		}

		expr := scanner.Text()
		if strings.TrimSpace(expr) == "" {
			continue
		}

		if expr == "q" || expr == "exit" {
			fmt.Println("sayonara...")
			break
		}

		safeEvalPrint(expr)
	}
}
