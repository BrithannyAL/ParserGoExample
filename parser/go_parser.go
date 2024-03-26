// Code generated from C:/Users/oviquez/Documents/00.Cursos/Compiladores (IC-5701)/2024/I Semestre/tests/GoExample/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // goParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type goParser struct {
	*antlr.BaseParser
}

var GoParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goparserParserInit() {
	staticData := &GoParserParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "':='", "'('", "')'", "'~'", "':'", "'+'", "'-'", "'*'",
		"'/'", "','", "'if'", "'while'", "'let'", "'then'", "'else'", "'in'",
		"'do'", "'begin'", "'end'", "'const'", "'var'",
	}
	staticData.SymbolicNames = []string{
		"", "PyCOMA", "ASSIGN", "PIZQ", "PDER", "VIR", "DOSPUN", "SUM", "SUB",
		"MUL", "DIV", "COMA", "IF", "WHILE", "LET", "THEN", "ELSE", "IN", "DO",
		"BEGIN", "END", "CONST", "VAR", "ID", "NUM", "LINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "command", "singleCommand", "declaration", "singleDeclaration",
		"typeDenoter", "expression", "primaryExpression", "operator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 100, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 36, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 64, 8, 3, 10, 3,
		12, 3, 67, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		77, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 85, 8, 6, 10, 6, 12,
		6, 88, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 96, 8, 7, 1, 8,
		1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 7, 10,
		101, 0, 18, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 60, 1,
		0, 0, 0, 8, 76, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 80, 1, 0, 0, 0, 14,
		95, 1, 0, 0, 0, 16, 97, 1, 0, 0, 0, 18, 19, 3, 4, 2, 0, 19, 1, 1, 0, 0,
		0, 20, 25, 3, 4, 2, 0, 21, 22, 5, 1, 0, 0, 22, 24, 3, 4, 2, 0, 23, 21,
		1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0,
		26, 3, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 35, 5, 23, 0, 0, 29, 30, 5,
		2, 0, 0, 30, 36, 3, 12, 6, 0, 31, 32, 5, 3, 0, 0, 32, 33, 3, 12, 6, 0,
		33, 34, 5, 4, 0, 0, 34, 36, 1, 0, 0, 0, 35, 29, 1, 0, 0, 0, 35, 31, 1,
		0, 0, 0, 36, 59, 1, 0, 0, 0, 37, 38, 5, 12, 0, 0, 38, 39, 3, 12, 6, 0,
		39, 40, 5, 15, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 5, 16, 0, 0, 42, 43, 3,
		4, 2, 0, 43, 59, 1, 0, 0, 0, 44, 45, 5, 13, 0, 0, 45, 46, 3, 12, 6, 0,
		46, 47, 5, 18, 0, 0, 47, 48, 3, 4, 2, 0, 48, 59, 1, 0, 0, 0, 49, 50, 5,
		14, 0, 0, 50, 51, 3, 6, 3, 0, 51, 52, 5, 17, 0, 0, 52, 53, 3, 4, 2, 0,
		53, 59, 1, 0, 0, 0, 54, 55, 5, 19, 0, 0, 55, 56, 3, 2, 1, 0, 56, 57, 5,
		20, 0, 0, 57, 59, 1, 0, 0, 0, 58, 28, 1, 0, 0, 0, 58, 37, 1, 0, 0, 0, 58,
		44, 1, 0, 0, 0, 58, 49, 1, 0, 0, 0, 58, 54, 1, 0, 0, 0, 59, 5, 1, 0, 0,
		0, 60, 65, 3, 8, 4, 0, 61, 62, 5, 1, 0, 0, 62, 64, 3, 8, 4, 0, 63, 61,
		1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 7, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5, 21, 0, 0, 69, 70, 5,
		23, 0, 0, 70, 71, 5, 5, 0, 0, 71, 77, 3, 12, 6, 0, 72, 73, 5, 22, 0, 0,
		73, 74, 5, 23, 0, 0, 74, 75, 5, 6, 0, 0, 75, 77, 3, 10, 5, 0, 76, 68, 1,
		0, 0, 0, 76, 72, 1, 0, 0, 0, 77, 9, 1, 0, 0, 0, 78, 79, 5, 23, 0, 0, 79,
		11, 1, 0, 0, 0, 80, 86, 3, 14, 7, 0, 81, 82, 3, 16, 8, 0, 82, 83, 3, 14,
		7, 0, 83, 85, 1, 0, 0, 0, 84, 81, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 13, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0,
		89, 96, 5, 24, 0, 0, 90, 96, 5, 23, 0, 0, 91, 92, 5, 3, 0, 0, 92, 93, 3,
		12, 6, 0, 93, 94, 5, 4, 0, 0, 94, 96, 1, 0, 0, 0, 95, 89, 1, 0, 0, 0, 95,
		90, 1, 0, 0, 0, 95, 91, 1, 0, 0, 0, 96, 15, 1, 0, 0, 0, 97, 98, 7, 0, 0,
		0, 98, 17, 1, 0, 0, 0, 7, 25, 35, 58, 65, 76, 86, 95,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// goParserInit initializes any static state used to implement goParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoParserInit() {
	staticData := &GoParserParserStaticData
	staticData.once.Do(goparserParserInit)
}

// NewgoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgoParser(input antlr.TokenStream) *goParser {
	GoParserInit()
	this := new(goParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "goParser.g4"

	return this
}

// goParser tokens.
const (
	goParserEOF          = antlr.TokenEOF
	goParserPyCOMA       = 1
	goParserASSIGN       = 2
	goParserPIZQ         = 3
	goParserPDER         = 4
	goParserVIR          = 5
	goParserDOSPUN       = 6
	goParserSUM          = 7
	goParserSUB          = 8
	goParserMUL          = 9
	goParserDIV          = 10
	goParserCOMA         = 11
	goParserIF           = 12
	goParserWHILE        = 13
	goParserLET          = 14
	goParserTHEN         = 15
	goParserELSE         = 16
	goParserIN           = 17
	goParserDO           = 18
	goParserBEGIN        = 19
	goParserEND          = 20
	goParserCONST        = 21
	goParserVAR          = 22
	goParserID           = 23
	goParserNUM          = 24
	goParserLINE_COMMENT = 25
	goParserWS           = 26
)

// goParser rules.
const (
	goParserRULE_program           = 0
	goParserRULE_command           = 1
	goParserRULE_singleCommand     = 2
	goParserRULE_declaration       = 3
	goParserRULE_singleDeclaration = 4
	goParserRULE_typeDenoter       = 5
	goParserRULE_expression        = 6
	goParserRULE_primaryExpression = 7
	goParserRULE_operator          = 8
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleCommand() ISingleCommandContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) SingleCommand() ISingleCommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleCommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleCommandContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, goParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.SingleCommand()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleCommand() []ISingleCommandContext
	SingleCommand(i int) ISingleCommandContext
	AllPyCOMA() []antlr.TerminalNode
	PyCOMA(i int) antlr.TerminalNode

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) AllSingleCommand() []ISingleCommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleCommandContext); ok {
			len++
		}
	}

	tst := make([]ISingleCommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleCommandContext); ok {
			tst[i] = t.(ISingleCommandContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) SingleCommand(i int) ISingleCommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleCommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleCommandContext)
}

func (s *CommandContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(goParserPyCOMA)
}

func (s *CommandContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(goParserPyCOMA, i)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, goParserRULE_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.SingleCommand()
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserPyCOMA {
		{
			p.SetState(21)
			p.Match(goParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.SingleCommand()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleCommandContext is an interface to support dynamic dispatch.
type ISingleCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	PIZQ() antlr.TerminalNode
	PDER() antlr.TerminalNode
	IF() antlr.TerminalNode
	THEN() antlr.TerminalNode
	AllSingleCommand() []ISingleCommandContext
	SingleCommand(i int) ISingleCommandContext
	ELSE() antlr.TerminalNode
	WHILE() antlr.TerminalNode
	DO() antlr.TerminalNode
	LET() antlr.TerminalNode
	Declaration() IDeclarationContext
	IN() antlr.TerminalNode
	BEGIN() antlr.TerminalNode
	Command() ICommandContext
	END() antlr.TerminalNode

	// IsSingleCommandContext differentiates from other interfaces.
	IsSingleCommandContext()
}

type SingleCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleCommandContext() *SingleCommandContext {
	var p = new(SingleCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleCommand
	return p
}

func InitEmptySingleCommandContext(p *SingleCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleCommand
}

func (*SingleCommandContext) IsSingleCommandContext() {}

func NewSingleCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleCommandContext {
	var p = new(SingleCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_singleCommand

	return p
}

func (s *SingleCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleCommandContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *SingleCommandContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goParserASSIGN, 0)
}

func (s *SingleCommandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SingleCommandContext) PIZQ() antlr.TerminalNode {
	return s.GetToken(goParserPIZQ, 0)
}

func (s *SingleCommandContext) PDER() antlr.TerminalNode {
	return s.GetToken(goParserPDER, 0)
}

func (s *SingleCommandContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *SingleCommandContext) THEN() antlr.TerminalNode {
	return s.GetToken(goParserTHEN, 0)
}

func (s *SingleCommandContext) AllSingleCommand() []ISingleCommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleCommandContext); ok {
			len++
		}
	}

	tst := make([]ISingleCommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleCommandContext); ok {
			tst[i] = t.(ISingleCommandContext)
			i++
		}
	}

	return tst
}

func (s *SingleCommandContext) SingleCommand(i int) ISingleCommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleCommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleCommandContext)
}

func (s *SingleCommandContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *SingleCommandContext) WHILE() antlr.TerminalNode {
	return s.GetToken(goParserWHILE, 0)
}

func (s *SingleCommandContext) DO() antlr.TerminalNode {
	return s.GetToken(goParserDO, 0)
}

func (s *SingleCommandContext) LET() antlr.TerminalNode {
	return s.GetToken(goParserLET, 0)
}

func (s *SingleCommandContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *SingleCommandContext) IN() antlr.TerminalNode {
	return s.GetToken(goParserIN, 0)
}

func (s *SingleCommandContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(goParserBEGIN, 0)
}

func (s *SingleCommandContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *SingleCommandContext) END() antlr.TerminalNode {
	return s.GetToken(goParserEND, 0)
}

func (s *SingleCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) SingleCommand() (localctx ISingleCommandContext) {
	localctx = NewSingleCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, goParserRULE_singleCommand)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserASSIGN:
			{
				p.SetState(29)
				p.Match(goParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(30)
				p.Expression()
			}

		case goParserPIZQ:
			{
				p.SetState(31)
				p.Match(goParserPIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(32)
				p.Expression()
			}
			{
				p.SetState(33)
				p.Match(goParserPDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case goParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Expression()
		}
		{
			p.SetState(39)
			p.Match(goParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.SingleCommand()
		}
		{
			p.SetState(41)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.SingleCommand()
		}

	case goParserWHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Match(goParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Expression()
		}
		{
			p.SetState(46)
			p.Match(goParserDO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.SingleCommand()
		}

	case goParserLET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(goParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Declaration()
		}
		{
			p.SetState(51)
			p.Match(goParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.SingleCommand()
		}

	case goParserBEGIN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)
			p.Match(goParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Command()
		}
		{
			p.SetState(56)
			p.Match(goParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleDeclaration() []ISingleDeclarationContext
	SingleDeclaration(i int) ISingleDeclarationContext
	AllPyCOMA() []antlr.TerminalNode
	PyCOMA(i int) antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) AllSingleDeclaration() []ISingleDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ISingleDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleDeclarationContext); ok {
			tst[i] = t.(ISingleDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) SingleDeclaration(i int) ISingleDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleDeclarationContext)
}

func (s *DeclarationContext) AllPyCOMA() []antlr.TerminalNode {
	return s.GetTokens(goParserPyCOMA)
}

func (s *DeclarationContext) PyCOMA(i int) antlr.TerminalNode {
	return s.GetToken(goParserPyCOMA, i)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, goParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.SingleDeclaration()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserPyCOMA {
		{
			p.SetState(61)
			p.Match(goParserPyCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.SingleDeclaration()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleDeclarationContext is an interface to support dynamic dispatch.
type ISingleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	ID() antlr.TerminalNode
	VIR() antlr.TerminalNode
	Expression() IExpressionContext
	VAR() antlr.TerminalNode
	DOSPUN() antlr.TerminalNode
	TypeDenoter() ITypeDenoterContext

	// IsSingleDeclarationContext differentiates from other interfaces.
	IsSingleDeclarationContext()
}

type SingleDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleDeclarationContext() *SingleDeclarationContext {
	var p = new(SingleDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleDeclaration
	return p
}

func InitEmptySingleDeclarationContext(p *SingleDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleDeclaration
}

func (*SingleDeclarationContext) IsSingleDeclarationContext() {}

func NewSingleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleDeclarationContext {
	var p = new(SingleDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_singleDeclaration

	return p
}

func (s *SingleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleDeclarationContext) CONST() antlr.TerminalNode {
	return s.GetToken(goParserCONST, 0)
}

func (s *SingleDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *SingleDeclarationContext) VIR() antlr.TerminalNode {
	return s.GetToken(goParserVIR, 0)
}

func (s *SingleDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SingleDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(goParserVAR, 0)
}

func (s *SingleDeclarationContext) DOSPUN() antlr.TerminalNode {
	return s.GetToken(goParserDOSPUN, 0)
}

func (s *SingleDeclarationContext) TypeDenoter() ITypeDenoterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDenoterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDenoterContext)
}

func (s *SingleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) SingleDeclaration() (localctx ISingleDeclarationContext) {
	localctx = NewSingleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, goParserRULE_singleDeclaration)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(goParserCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(goParserVIR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Expression()
		}

	case goParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(goParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(goParserDOSPUN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.TypeDenoter()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDenoterContext is an interface to support dynamic dispatch.
type ITypeDenoterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsTypeDenoterContext differentiates from other interfaces.
	IsTypeDenoterContext()
}

type TypeDenoterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDenoterContext() *TypeDenoterContext {
	var p = new(TypeDenoterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_typeDenoter
	return p
}

func InitEmptyTypeDenoterContext(p *TypeDenoterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_typeDenoter
}

func (*TypeDenoterContext) IsTypeDenoterContext() {}

func NewTypeDenoterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDenoterContext {
	var p = new(TypeDenoterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_typeDenoter

	return p
}

func (s *TypeDenoterContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDenoterContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *TypeDenoterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDenoterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) TypeDenoter() (localctx ITypeDenoterContext) {
	localctx = NewTypeDenoterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, goParserRULE_typeDenoter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrimaryExpression() []IPrimaryExpressionContext
	PrimaryExpression(i int) IPrimaryExpressionContext
	AllOperator() []IOperatorContext
	Operator(i int) IOperatorContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExpressionContext); ok {
			tst[i] = t.(IPrimaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionContext) AllOperator() []IOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperatorContext); ok {
			len++
		}
	}

	tst := make([]IOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperatorContext); ok {
			tst[i] = t.(IOperatorContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Operator(i int) IOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, goParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.PrimaryExpression()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1920) != 0 {
		{
			p.SetState(81)
			p.Operator()
		}
		{
			p.SetState(82)
			p.PrimaryExpression()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUM() antlr.TerminalNode
	ID() antlr.TerminalNode
	PIZQ() antlr.TerminalNode
	Expression() IExpressionContext
	PDER() antlr.TerminalNode

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) NUM() antlr.TerminalNode {
	return s.GetToken(goParserNUM, 0)
}

func (s *PrimaryExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *PrimaryExpressionContext) PIZQ() antlr.TerminalNode {
	return s.GetToken(goParserPIZQ, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) PDER() antlr.TerminalNode {
	return s.GetToken(goParserPDER, 0)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, goParserRULE_primaryExpression)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserNUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Match(goParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserPIZQ:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(goParserPIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Expression()
		}
		{
			p.SetState(93)
			p.Match(goParserPDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUM() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) SUM() antlr.TerminalNode {
	return s.GetToken(goParserSUM, 0)
}

func (s *OperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(goParserSUB, 0)
}

func (s *OperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(goParserMUL, 0)
}

func (s *OperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(goParserDIV, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *goParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, goParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
