package main

import (
	"fmt"
	"os"
	"github.com/edersonferreira/denol/parser"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var tokensName []string
var tokensText []string
var functions []string

var phpWords = []string{"echo", "if", "else", "elseif", "for", "do", "switch", "case", "try", "catch", "continue", "finally", "default", "while", "break", "exit", "function", "import", "print", "return", "require", "halt_compiler()","abstract","array()","as","break","callable","class","clone","const","declare","die()","empty()","enddeclare","endfor","endforeach","endif","endswitch","endwhile","eval()","extends","final","fn","foreach","global","goto","implements","include","include_once","instanceof","insteadof","interface","isset()","list()","namespace","new","or","private","protected","public","require_once","static","throw","trait","unset()","use","var","xor","yield"}
var denolWords = []string{"ola", "salve", "naosalvou", "outrosalve","pamonhosa", "day","teste", "sepa", "tenta", "pega", "mandabala", "amain", "padraozao", "zoeira", "cancelar", "deno", "esporro", "!important", "printa", "veio", "pediu"}

type denolListener struct {
	*parser.BaseDenolListener
}

func addText(tokenText string, tokenName string) string{
		tokensName = append(tokensName, tokenName)
		if tokenName != "WS"{
			tokensText = append(tokensText, tokenText)
		}
	if tokenName == "NL"{
		if tokensName[len(tokensName) - 2] != "TWOPOINS" && tokensName[len(tokensName) - 2] != "PAREN_START" && tokensName[len(tokensName) - 2] != "SQRT_START" && tokensName[len(tokensName) - 2] != "KEYS_START"{
			return ";\n"
		}
		return "\n"
	}else if contains(phpWords, tokenText) && tokenName != "STRING"{
		return tokenText
	}else if contains(functions, tokenText){
		return tokenText
	}else if tokenText == "+" && tokensName[len(tokensName) - 2] == "STRING"{
		return "."
	}else if tokenName == "PI"{
		return "3.1415"
	}else if contains(denolWords, tokenText) && tokenName != "STRING"{
		return phpWords[IndexOf(denolWords, tokenText)]
	}else if tokenName == "ID"{
		if tokensText[len(tokensText) - 2] == "esporro"{
			functions = append(functions, tokenText)	
			return tokenText
		}
			return "$" + tokenText
	}else if tokenName == "WS"{
		return " "
	}
	return tokenText
}

func main() {

	os.Mkdir("build", 0755)

	for i := 1; i < len(os.Args); i++{

		input, _ := antlr.NewFileStream(os.Args[i])
		lexer := parser.NewDenolLexer(input)
		fmt.Println(lexer.GetErrorHeader())
		output := "<?php\n"
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			output += addText(t.GetText(), lexer.SymbolicNames[t.GetTokenType()])
		}
		f, err := os.Create( "build/" + strings.Replace(os.Args[i], ".denol", ".php", -1))
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
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
 func IndexOf(haystack []string, str string) int {
    for i, v := range haystack {
        if v == str {
            return i
        }
    }
    return -1
}
