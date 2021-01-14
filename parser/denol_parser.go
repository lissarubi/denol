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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 118,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 52, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 59, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 70, 10, 5, 12,
	5, 14, 5, 73, 11, 5, 3, 6, 3, 6, 3, 6, 5, 6, 78, 10, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7,
	7, 106, 10, 7, 12, 7, 14, 7, 109, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 116, 10, 7, 3, 7, 2, 3, 8, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 132, 2,
	14, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 58, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2,
	10, 77, 3, 2, 2, 2, 12, 115, 3, 2, 2, 2, 14, 15, 7, 3, 2, 2, 15, 16, 5,
	4, 3, 2, 16, 17, 7, 4, 2, 2, 17, 3, 3, 2, 2, 2, 18, 52, 5, 6, 4, 2, 19,
	20, 7, 30, 2, 2, 20, 21, 7, 5, 2, 2, 21, 22, 5, 4, 3, 2, 22, 23, 7, 6,
	2, 2, 23, 52, 3, 2, 2, 2, 24, 25, 7, 30, 2, 2, 25, 26, 7, 7, 2, 2, 26,
	27, 5, 4, 3, 2, 27, 28, 7, 6, 2, 2, 28, 52, 3, 2, 2, 2, 29, 30, 7, 30,
	2, 2, 30, 31, 7, 8, 2, 2, 31, 32, 5, 4, 3, 2, 32, 33, 7, 6, 2, 2, 33, 52,
	3, 2, 2, 2, 34, 35, 7, 30, 2, 2, 35, 36, 7, 9, 2, 2, 36, 37, 5, 4, 3, 2,
	37, 38, 7, 6, 2, 2, 38, 52, 3, 2, 2, 2, 39, 40, 7, 30, 2, 2, 40, 41, 7,
	10, 2, 2, 41, 42, 5, 4, 3, 2, 42, 43, 7, 6, 2, 2, 43, 52, 3, 2, 2, 2, 44,
	45, 7, 30, 2, 2, 45, 46, 7, 11, 2, 2, 46, 47, 5, 4, 3, 2, 47, 48, 7, 6,
	2, 2, 48, 52, 3, 2, 2, 2, 49, 52, 7, 12, 2, 2, 50, 52, 7, 13, 2, 2, 51,
	18, 3, 2, 2, 2, 51, 19, 3, 2, 2, 2, 51, 24, 3, 2, 2, 2, 51, 29, 3, 2, 2,
	2, 51, 34, 3, 2, 2, 2, 51, 39, 3, 2, 2, 2, 51, 44, 3, 2, 2, 2, 51, 49,
	3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 5, 3, 2, 2, 2, 53, 59, 5, 8, 5, 2,
	54, 55, 5, 8, 5, 2, 55, 56, 7, 10, 2, 2, 56, 57, 5, 8, 5, 2, 57, 59, 3,
	2, 2, 2, 58, 53, 3, 2, 2, 2, 58, 54, 3, 2, 2, 2, 59, 7, 3, 2, 2, 2, 60,
	61, 8, 5, 1, 2, 61, 62, 5, 10, 6, 2, 62, 71, 3, 2, 2, 2, 63, 64, 12, 4,
	2, 2, 64, 65, 7, 14, 2, 2, 65, 70, 5, 10, 6, 2, 66, 67, 12, 3, 2, 2, 67,
	68, 7, 15, 2, 2, 68, 70, 5, 10, 6, 2, 69, 63, 3, 2, 2, 2, 69, 66, 3, 2,
	2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 9,
	3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 78, 7, 30, 2, 2, 75, 78, 7, 24, 2,
	2, 76, 78, 5, 2, 2, 2, 77, 74, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 76,
	3, 2, 2, 2, 78, 11, 3, 2, 2, 2, 79, 80, 7, 16, 2, 2, 80, 81, 5, 2, 2, 2,
	81, 82, 5, 12, 7, 2, 82, 116, 3, 2, 2, 2, 83, 84, 7, 16, 2, 2, 84, 85,
	5, 2, 2, 2, 85, 86, 5, 12, 7, 2, 86, 87, 7, 17, 2, 2, 87, 88, 5, 12, 7,
	2, 88, 116, 3, 2, 2, 2, 89, 90, 7, 18, 2, 2, 90, 91, 5, 2, 2, 2, 91, 92,
	5, 12, 7, 2, 92, 116, 3, 2, 2, 2, 93, 94, 7, 19, 2, 2, 94, 95, 5, 2, 2,
	2, 95, 96, 5, 12, 7, 2, 96, 116, 3, 2, 2, 2, 97, 98, 7, 20, 2, 2, 98, 99,
	5, 12, 7, 2, 99, 100, 7, 18, 2, 2, 100, 101, 5, 2, 2, 2, 101, 102, 7, 6,
	2, 2, 102, 116, 3, 2, 2, 2, 103, 107, 7, 21, 2, 2, 104, 106, 5, 12, 7,
	2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107,
	108, 3, 2, 2, 2, 108, 110, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 116,
	7, 22, 2, 2, 111, 112, 5, 4, 3, 2, 112, 113, 7, 6, 2, 2, 113, 116, 3, 2,
	2, 2, 114, 116, 7, 6, 2, 2, 115, 79, 3, 2, 2, 2, 115, 83, 3, 2, 2, 2, 115,
	89, 3, 2, 2, 2, 115, 93, 3, 2, 2, 2, 115, 97, 3, 2, 2, 2, 115, 103, 3,
	2, 2, 2, 115, 111, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 13, 3, 2, 2,
	2, 9, 51, 58, 69, 71, 77, 107, 115,
}
var literalNames = []string{
	"", "'('", "')'", "'=='", "';'", "'>='", "'<='", "'>'", "'<'", "'='", "'&&'",
	"'||'", "'+'", "'-'", "'salve'", "'naosalvou'", "'zoeira'", "'pamonhosa'",
	"'day'", "'{'", "'}'", "", "", "", "'pi'", "'^'", "'\n'", "", "", "' '",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "OLA", "INT", "DOUBLE", "PI", "POW", "NL", "WS", "ID", "SPACE",
	"STRING",
}

var ruleNames = []string{
	"paren_expr", "expr", "test", "sum", "term", "statement",
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
	DenolParserEOF    = antlr.TokenEOF
	DenolParserT__0   = 1
	DenolParserT__1   = 2
	DenolParserT__2   = 3
	DenolParserT__3   = 4
	DenolParserT__4   = 5
	DenolParserT__5   = 6
	DenolParserT__6   = 7
	DenolParserT__7   = 8
	DenolParserT__8   = 9
	DenolParserT__9   = 10
	DenolParserT__10  = 11
	DenolParserT__11  = 12
	DenolParserT__12  = 13
	DenolParserT__13  = 14
	DenolParserT__14  = 15
	DenolParserT__15  = 16
	DenolParserT__16  = 17
	DenolParserT__17  = 18
	DenolParserT__18  = 19
	DenolParserT__19  = 20
	DenolParserOLA    = 21
	DenolParserINT    = 22
	DenolParserDOUBLE = 23
	DenolParserPI     = 24
	DenolParserPOW    = 25
	DenolParserNL     = 26
	DenolParserWS     = 27
	DenolParserID     = 28
	DenolParserSPACE  = 29
	DenolParserSTRING = 30
)

// DenolParser rules.
const (
	DenolParserRULE_paren_expr = 0
	DenolParserRULE_expr       = 1
	DenolParserRULE_test       = 2
	DenolParserRULE_sum        = 3
	DenolParserRULE_term       = 4
	DenolParserRULE_statement  = 5
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
		p.SetState(12)
		p.Match(DenolParserT__0)
	}
	{
		p.SetState(13)
		p.Expr()
	}
	{
		p.SetState(14)
		p.Match(DenolParserT__1)
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
	p.EnterRule(localctx, 2, DenolParserRULE_expr)

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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Test()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Match(DenolParserID)
		}
		{
			p.SetState(18)
			p.Match(DenolParserT__2)
		}
		{
			p.SetState(19)
			p.Expr()
		}
		{
			p.SetState(20)
			p.Match(DenolParserT__3)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Match(DenolParserID)
		}
		{
			p.SetState(23)
			p.Match(DenolParserT__4)
		}
		{
			p.SetState(24)
			p.Expr()
		}
		{
			p.SetState(25)
			p.Match(DenolParserT__3)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.Match(DenolParserID)
		}
		{
			p.SetState(28)
			p.Match(DenolParserT__5)
		}
		{
			p.SetState(29)
			p.Expr()
		}
		{
			p.SetState(30)
			p.Match(DenolParserT__3)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)
			p.Match(DenolParserID)
		}
		{
			p.SetState(33)
			p.Match(DenolParserT__6)
		}
		{
			p.SetState(34)
			p.Expr()
		}
		{
			p.SetState(35)
			p.Match(DenolParserT__3)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(37)
			p.Match(DenolParserID)
		}
		{
			p.SetState(38)
			p.Match(DenolParserT__7)
		}
		{
			p.SetState(39)
			p.Expr()
		}
		{
			p.SetState(40)
			p.Match(DenolParserT__3)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(42)
			p.Match(DenolParserID)
		}
		{
			p.SetState(43)
			p.Match(DenolParserT__8)
		}
		{
			p.SetState(44)
			p.Expr()
		}
		{
			p.SetState(45)
			p.Match(DenolParserT__3)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(47)
			p.Match(DenolParserT__9)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(48)
			p.Match(DenolParserT__10)
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
	p.EnterRule(localctx, 4, DenolParserRULE_test)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.sum(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.sum(0)
		}
		{
			p.SetState(53)
			p.Match(DenolParserT__7)
		}
		{
			p.SetState(54)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, DenolParserRULE_sum, _p)

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
		p.SetState(59)
		p.Term()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSumContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DenolParserRULE_sum)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(62)
					p.Match(DenolParserT__11)
				}
				{
					p.SetState(63)
					p.Term()
				}

			case 2:
				localctx = NewSumContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DenolParserRULE_sum)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(65)
					p.Match(DenolParserT__12)
				}
				{
					p.SetState(66)
					p.Term()
				}

			}

		}
		p.SetState(71)
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

func (s *TermContext) INT() antlr.TerminalNode {
	return s.GetToken(DenolParserINT, 0)
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
	p.EnterRule(localctx, 8, DenolParserRULE_term)

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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DenolParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(DenolParserID)
		}

	case DenolParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(DenolParserINT)
		}

	case DenolParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
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
	p.EnterRule(localctx, 10, DenolParserRULE_statement)
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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(DenolParserT__13)
		}
		{
			p.SetState(78)
			p.Paren_expr()
		}
		{
			p.SetState(79)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(DenolParserT__13)
		}
		{
			p.SetState(82)
			p.Paren_expr()
		}
		{
			p.SetState(83)
			p.Statement()
		}
		{
			p.SetState(84)
			p.Match(DenolParserT__14)
		}
		{
			p.SetState(85)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Match(DenolParserT__15)
		}
		{
			p.SetState(88)
			p.Paren_expr()
		}
		{
			p.SetState(89)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Match(DenolParserT__16)
		}
		{
			p.SetState(92)
			p.Paren_expr()
		}
		{
			p.SetState(93)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)
			p.Match(DenolParserT__17)
		}
		{
			p.SetState(96)
			p.Statement()
		}
		{
			p.SetState(97)
			p.Match(DenolParserT__15)
		}
		{
			p.SetState(98)
			p.Paren_expr()
		}
		{
			p.SetState(99)
			p.Match(DenolParserT__3)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.Match(DenolParserT__18)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DenolParserT__0)|(1<<DenolParserT__3)|(1<<DenolParserT__9)|(1<<DenolParserT__10)|(1<<DenolParserT__13)|(1<<DenolParserT__15)|(1<<DenolParserT__16)|(1<<DenolParserT__17)|(1<<DenolParserT__18)|(1<<DenolParserINT)|(1<<DenolParserID))) != 0 {
			{
				p.SetState(102)
				p.Statement()
			}

			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(108)
			p.Match(DenolParserT__19)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Expr()
		}
		{
			p.SetState(110)
			p.Match(DenolParserT__3)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(112)
			p.Match(DenolParserT__3)
		}

	}

	return localctx
}

func (p *DenolParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
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
