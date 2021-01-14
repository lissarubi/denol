package main

import (
	"fmt"
	"parser/parser"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type denolListener struct {
	*parser.BaseDenolListener
}

func addText(tokenText string, tokenName string) string{
	if tokenName == "NL"{
		return "\n"
	}else if tokenName == "ID"{
		return "$" + tokenText
	}else if tokenName == "PI"{
		return "3.1415"
	}else if tokenText == "ola"{
		return "echo"
	}else if tokenText == "salve"{
		return "if"
	}else if tokenText == "pamonhosa"{
		return "for"
	}else if tokenText == "day"{
		return "do"
	}else if tokenText == "zoeira"{
		return "while"
	}
	return tokenText
}

func main() {
	input, _ := antlr.NewFileStream("hello.denol")
	lexer := parser.NewDenolLexer(input)
	output := "<?php\n"
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		output += addText(t.GetText(), lexer.SymbolicNames[t.GetTokenType()])
	}
	f, err := os.Create("index.php")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(output)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
