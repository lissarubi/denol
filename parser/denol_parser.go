// Code generated from Denol.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Denol

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 130,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 5, 5, 64, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 71, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 82, 10, 7, 12,
	7, 14, 7, 85, 11, 7, 3, 8, 3, 8, 3, 8, 5, 8, 90, 10, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7,
	9, 118, 10, 9, 12, 9, 14, 9, 121, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 128, 10, 9, 3, 9, 2, 3, 12, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2,
	2, 142, 2, 18, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 63,
	3, 2, 2, 2, 10, 70, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 89, 3, 2, 2, 2,
	16, 127, 3, 2, 2, 2, 18, 19, 7, 31, 2, 2, 19, 20, 5, 8, 5, 2, 20, 21, 7,
	3, 2, 2, 21, 3, 3, 2, 2, 2, 22, 23, 7, 32, 2, 2, 23, 24, 5, 8, 5, 2, 24,
	25, 7, 4, 2, 2, 25, 5, 3, 2, 2, 2, 26, 27, 7, 33, 2, 2, 27, 28, 5, 8, 5,
	2, 28, 29, 7, 5, 2, 2, 29, 7, 3, 2, 2, 2, 30, 64, 5, 10, 6, 2, 31, 32,
	7, 28, 2, 2, 32, 33, 7, 6, 2, 2, 33, 34, 5, 8, 5, 2, 34, 35, 7, 7, 2, 2,
	35, 64, 3, 2, 2, 2, 36, 37, 7, 28, 2, 2, 37, 38, 7, 8, 2, 2, 38, 39, 5,
	8, 5, 2, 39, 40, 7, 7, 2, 2, 40, 64, 3, 2, 2, 2, 41, 42, 7, 28, 2, 2, 42,
	43, 7, 9, 2, 2, 43, 44, 5, 8, 5, 2, 44, 45, 7, 7, 2, 2, 45, 64, 3, 2, 2,
	2, 46, 47, 7, 28, 2, 2, 47, 48, 7, 10, 2, 2, 48, 49, 5, 8, 5, 2, 49, 50,
	7, 7, 2, 2, 50, 64, 3, 2, 2, 2, 51, 52, 7, 28, 2, 2, 52, 53, 7, 11, 2,
	2, 53, 54, 5, 8, 5, 2, 54, 55, 7, 7, 2, 2, 55, 64, 3, 2, 2, 2, 56, 57,
	7, 28, 2, 2, 57, 58, 7, 12, 2, 2, 58, 59, 5, 8, 5, 2, 59, 60, 7, 7, 2,
	2, 60, 64, 3, 2, 2, 2, 61, 64, 7, 13, 2, 2, 62, 64, 7, 14, 2, 2, 63, 30,
	3, 2, 2, 2, 63, 31, 3, 2, 2, 2, 63, 36, 3, 2, 2, 2, 63, 41, 3, 2, 2, 2,
	63, 46, 3, 2, 2, 2, 63, 51, 3, 2, 2, 2, 63, 56, 3, 2, 2, 2, 63, 61, 3,
	2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 9, 3, 2, 2, 2, 65, 71, 5, 12, 7, 2, 66,
	67, 5, 12, 7, 2, 67, 68, 7, 11, 2, 2, 68, 69, 5, 12, 7, 2, 69, 71, 3, 2,
	2, 2, 70, 65, 3, 2, 2, 2, 70, 66, 3, 2, 2, 2, 71, 11, 3, 2, 2, 2, 72, 73,
	8, 7, 1, 2, 73, 74, 5, 14, 8, 2, 74, 83, 3, 2, 2, 2, 75, 76, 12, 4, 2,
	2, 76, 77, 7, 15, 2, 2, 77, 82, 5, 14, 8, 2, 78, 79, 12, 3, 2, 2, 79, 80,
	7, 16, 2, 2, 80, 82, 5, 14, 8, 2, 81, 75, 3, 2, 2, 2, 81, 78, 3, 2, 2,
	2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 13,
	3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 90, 7, 28, 2, 2, 87, 90, 7, 22, 2,
	2, 88, 90, 5, 2, 2, 2, 89, 86, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 88,
	3, 2, 2, 2, 90, 15, 3, 2, 2, 2, 91, 92, 7, 17, 2, 2, 92, 93, 5, 2, 2, 2,
	93, 94, 5, 16, 9, 2, 94, 128, 3, 2, 2, 2, 95, 96, 7, 17, 2, 2, 96, 97,
	5, 2, 2, 2, 97, 98, 5, 16, 9, 2, 98, 99, 7, 18, 2, 2, 99, 100, 5, 16, 9,
	2, 100, 128, 3, 2, 2, 2, 101, 102, 7, 19, 2, 2, 102, 103, 5, 2, 2, 2, 103,
	104, 5, 16, 9, 2, 104, 128, 3, 2, 2, 2, 105, 106, 7, 20, 2, 2, 106, 107,
	5, 2, 2, 2, 107, 108, 5, 16, 9, 2, 108, 128, 3, 2, 2, 2, 109, 110, 7, 21,
	2, 2, 110, 111, 5, 16, 9, 2, 111, 112, 7, 19, 2, 2, 112, 113, 5, 2, 2,
	2, 113, 114, 7, 7, 2, 2, 114, 128, 3, 2, 2, 2, 115, 119, 7, 33, 2, 2, 116,
	118, 5, 16, 9, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121, 119, 3, 2,
	2, 2, 122, 128, 7, 5, 2, 2, 123, 124, 5, 8, 5, 2, 124, 125, 7, 7, 2, 2,
	125, 128, 3, 2, 2, 2, 126, 128, 7, 7, 2, 2, 127, 91, 3, 2, 2, 2, 127, 95,
	3, 2, 2, 2, 127, 101, 3, 2, 2, 2, 127, 105, 3, 2, 2, 2, 127, 109, 3, 2,
	2, 2, 127, 115, 3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2,
	128, 17, 3, 2, 2, 2, 9, 63, 70, 81, 83, 89, 119, 127,
}
var literalNames = []string{
	"", "')'", "']'", "'}'", "'=='", "';'", "'>='", "'<='", "'>'", "'<'", "'='",
	"'&&'", "'||'", "'+'", "'-'", "'salve'", "'naosalvou'", "'zoeira'", "'pamonhosa'",
	"'day'", "", "'pi'", "'^'", "", "", "':'", "", "' '", "", "'('", "'['",
	"'{'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "NUMBER", "PI", "POW", "NL", "WS", "TWOPOINS", "ID", "SPACE", "STRING",
	"PAREN_START", "SQRT_START", "KEYS_START",
}

var ruleNames = []string{
	"paren_expr", "sqrt_expr", "keys_expr", "expr", "test", "sum", "term",
	"statement",
}

type DenolParser struct {
	*antlr.BaseParser
}

// NewDenolParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DenolParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDenolParser(input antlr.TokenStream) *DenolParser {
	this := new(DenolParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Denol.g4"

	return this
}

// DenolParser tokens.
const (
	DenolParserEOF         = antlr.TokenEOF
	DenolParserT__0        = 1
	DenolParserT__1        = 2
	DenolParserT__2        = 3
	DenolParserT__3        = 4
	DenolParserT__4        = 5
	DenolParserT__5        = 6
	DenolParserT__6        = 7
	DenolParserT__7        = 8
	DenolParserT__8        = 9
	DenolParserT__9        = 10
	DenolParserT__10       = 11
	DenolParserT__11       = 12
	DenolParserT__12       = 13
	DenolParserT__13       = 14
	DenolParserT__14       = 15
	DenolParserT__15       = 16
	DenolParserT__16       = 17
	DenolParserT__17       = 18
	DenolParserT__18       = 19
	DenolParserNUMBER      = 20
	DenolParserPI          = 21
	DenolParserPOW         = 22
	DenolParserNL          = 23
	DenolParserWS          = 24
	DenolParserTWOPOINS    = 25
	DenolParserID          = 26
	DenolParserSPACE       = 27
	DenolParserSTRING      = 28
	DenolParserPAREN_START = 29
	DenolParserSQRT_START  = 30
	DenolParserKEYS_START  = 31
)

// DenolParser rules.
const (
	DenolParserRULE_paren_expr = 0
	DenolParserRULE_sqrt_expr  = 1
	DenolParserRULE_keys_expr  = 2
	DenolParserRULE_expr       = 3
	DenolParserRULE_test       = 4
	DenolParserRULE_sum        = 5
	DenolParserRULE_term       = 6
	DenolParserRULE_statement  = 7
)

// IParen_exprContext is an interface to support dynamic dispatch.
type IParen_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParen_exprContext differentiates from other interfaces.
	IsParen_exprContext()
}

type Paren_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_exprContext() *Paren_exprContext {
	var p = new(Paren_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_paren_expr
	return p
}

func (*Paren_exprContext) IsParen_exprContext() {}

func NewParen_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_exprContext {
	var p = new(Paren_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_paren_expr

	return p
}

func (s *Paren_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_exprContext) PAREN_START() antlr.TerminalNode {
	return s.GetToken(DenolParserPAREN_START, 0)
}

func (s *Paren_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Paren_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterParen_expr(s)
	}
}

func (s *Paren_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitParen_expr(s)
	}
}

func (p *DenolParser) Paren_expr() (localctx IParen_exprContext) {
	localctx = NewParen_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DenolParserRULE_paren_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(DenolParserPAREN_START)
	}
	{
		p.SetState(17)
		p.Expr()
	}
	{
		p.SetState(18)
		p.Match(DenolParserT__0)
	}

	return localctx
}

// ISqrt_exprContext is an interface to support dynamic dispatch.
type ISqrt_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqrt_exprContext differentiates from other interfaces.
	IsSqrt_exprContext()
}

type Sqrt_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqrt_exprContext() *Sqrt_exprContext {
	var p = new(Sqrt_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_sqrt_expr
	return p
}

func (*Sqrt_exprContext) IsSqrt_exprContext() {}

func NewSqrt_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sqrt_exprContext {
	var p = new(Sqrt_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_sqrt_expr

	return p
}

func (s *Sqrt_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Sqrt_exprContext) SQRT_START() antlr.TerminalNode {
	return s.GetToken(DenolParserSQRT_START, 0)
}

func (s *Sqrt_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Sqrt_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sqrt_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sqrt_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterSqrt_expr(s)
	}
}

func (s *Sqrt_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitSqrt_expr(s)
	}
}

func (p *DenolParser) Sqrt_expr() (localctx ISqrt_exprContext) {
	localctx = NewSqrt_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DenolParserRULE_sqrt_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(DenolParserSQRT_START)
	}
	{
		p.SetState(21)
		p.Expr()
	}
	{
		p.SetState(22)
		p.Match(DenolParserT__1)
	}

	return localctx
}

// IKeys_exprContext is an interface to support dynamic dispatch.
type IKeys_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeys_exprContext differentiates from other interfaces.
	IsKeys_exprContext()
}

type Keys_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeys_exprContext() *Keys_exprContext {
	var p = new(Keys_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_keys_expr
	return p
}

func (*Keys_exprContext) IsKeys_exprContext() {}

func NewKeys_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Keys_exprContext {
	var p = new(Keys_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_keys_expr

	return p
}

func (s *Keys_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Keys_exprContext) KEYS_START() antlr.TerminalNode {
	return s.GetToken(DenolParserKEYS_START, 0)
}

func (s *Keys_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Keys_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Keys_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Keys_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterKeys_expr(s)
	}
}

func (s *Keys_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitKeys_expr(s)
	}
}

func (p *DenolParser) Keys_expr() (localctx IKeys_exprContext) {
	localctx = NewKeys_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DenolParserRULE_keys_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(DenolParserKEYS_START)
	}
	{
		p.SetState(25)
		p.Expr()
	}
	{
		p.SetState(26)
		p.Match(DenolParserT__2)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(DenolParserID, 0)
}

func (s *ExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *DenolParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DenolParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Test()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Match(DenolParserID)
		}
		{
			p.SetState(30)
			p.Match(DenolParserT__3)
		}
		{
			p.SetState(31)
			p.Expr()
		}
		{
			p.SetState(32)
			p.Match(DenolParserT__4)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(DenolParserID)
		}
		{
			p.SetState(35)
			p.Match(DenolParserT__5)
		}
		{
			p.SetState(36)
			p.Expr()
		}
		{
			p.SetState(37)
			p.Match(DenolParserT__4)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Match(DenolParserID)
		}
		{
			p.SetState(40)
			p.Match(DenolParserT__6)
		}
		{
			p.SetState(41)
			p.Expr()
		}
		{
			p.SetState(42)
			p.Match(DenolParserT__4)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(44)
			p.Match(DenolParserID)
		}
		{
			p.SetState(45)
			p.Match(DenolParserT__7)
		}
		{
			p.SetState(46)
			p.Expr()
		}
		{
			p.SetState(47)
			p.Match(DenolParserT__4)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.Match(DenolParserID)
		}
		{
			p.SetState(50)
			p.Match(DenolParserT__8)
		}
		{
			p.SetState(51)
			p.Expr()
		}
		{
			p.SetState(52)
			p.Match(DenolParserT__4)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.Match(DenolParserID)
		}
		{
			p.SetState(55)
			p.Match(DenolParserT__9)
		}
		{
			p.SetState(56)
			p.Expr()
		}
		{
			p.SetState(57)
			p.Match(DenolParserT__4)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(59)
			p.Match(DenolParserT__10)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(60)
			p.Match(DenolParserT__11)
		}

	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) AllSum() []ISumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISumContext)(nil)).Elem())
	var tst = make([]ISumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISumContext)
		}
	}

	return tst
}

func (s *TestContext) Sum(i int) ISumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitTest(s)
	}
}

func (p *DenolParser) Test() (localctx ITestContext) {
	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DenolParserRULE_test)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.sum(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.sum(0)
		}
		{
			p.SetState(65)
			p.Match(DenolParserT__8)
		}
		{
			p.SetState(66)
			p.sum(0)
		}

	}

	return localctx
}

// ISumContext is an interface to support dynamic dispatch.
type ISumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSumContext differentiates from other interfaces.
	IsSumContext()
}

type SumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_sum
	return p
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_sum

	return p
}

func (s *SumContext) GetParser() antlr.Parser { return s.parser }

func (s *SumContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SumContext) Sum() ISumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitSum(s)
	}
}

func (p *DenolParser) Sum() (localctx ISumContext) {
	return p.sum(0)
}

func (p *DenolParser) sum(_p int) (localctx ISumContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSumContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISumContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, DenolParserRULE_sum, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Term()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSumContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DenolParserRULE_sum)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(74)
					p.Match(DenolParserT__12)
				}
				{
					p.SetState(75)
					p.Term()
				}

			case 2:
				localctx = NewSumContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DenolParserRULE_sum)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(77)
					p.Match(DenolParserT__13)
				}
				{
					p.SetState(78)
					p.Term()
				}

			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) ID() antlr.TerminalNode {
	return s.GetToken(DenolParserID, 0)
}

func (s *TermContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DenolParserNUMBER, 0)
}

func (s *TermContext) Paren_expr() IParen_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParen_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *DenolParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DenolParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DenolParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(DenolParserID)
		}

	case DenolParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(DenolParserNUMBER)
		}

	case DenolParserPAREN_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Paren_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DenolParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DenolParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Paren_expr() IParen_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParen_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) KEYS_START() antlr.TerminalNode {
	return s.GetToken(DenolParserKEYS_START, 0)
}

func (s *StatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DenolListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *DenolParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DenolParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Match(DenolParserT__14)
		}
		{
			p.SetState(90)
			p.Paren_expr()
		}
		{
			p.SetState(91)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(DenolParserT__14)
		}
		{
			p.SetState(94)
			p.Paren_expr()
		}
		{
			p.SetState(95)
			p.Statement()
		}
		{
			p.SetState(96)
			p.Match(DenolParserT__15)
		}
		{
			p.SetState(97)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Match(DenolParserT__16)
		}
		{
			p.SetState(100)
			p.Paren_expr()
		}
		{
			p.SetState(101)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Match(DenolParserT__17)
		}
		{
			p.SetState(104)
			p.Paren_expr()
		}
		{
			p.SetState(105)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Match(DenolParserT__18)
		}
		{
			p.SetState(108)
			p.Statement()
		}
		{
			p.SetState(109)
			p.Match(DenolParserT__16)
		}
		{
			p.SetState(110)
			p.Paren_expr()
		}
		{
			p.SetState(111)
			p.Match(DenolParserT__4)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.Match(DenolParserKEYS_START)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DenolParserT__4)|(1<<DenolParserT__10)|(1<<DenolParserT__11)|(1<<DenolParserT__14)|(1<<DenolParserT__16)|(1<<DenolParserT__17)|(1<<DenolParserT__18)|(1<<DenolParserNUMBER)|(1<<DenolParserID)|(1<<DenolParserPAREN_START)|(1<<DenolParserKEYS_START))) != 0 {
			{
				p.SetState(114)
				p.Statement()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
			p.Match(DenolParserT__2)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(121)
			p.Expr()
		}
		{
			p.SetState(122)
			p.Match(DenolParserT__4)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(124)
			p.Match(DenolParserT__4)
		}

	}

	return localctx
}

func (p *DenolParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *SumContext = nil
		if localctx != nil {
			t = localctx.(*SumContext)
		}
		return p.Sum_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DenolParser) Sum_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
