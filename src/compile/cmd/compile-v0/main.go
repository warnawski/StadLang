package main

import (
	"compiler/src/compile/internal/parser"
	"fmt"
)

func main() {

	tokList := parser.TokenizationV3(
		`public int x:=24 //Test comment skip
public int y:=56
public int c := x / y`)

	fmt.Println(tokList)
}
