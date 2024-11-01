// Code generated from C:/Users/oviquez/Documents/00.Cursos/Compiladores (IC-5701)/2024/I Semestre/tests/GoExample/goScanner.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type goScanner struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GoScannerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goscannerLexerInit() {
	staticData := &GoScannerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"PyCOMA", "ASSIGN", "PIZQ", "PDER", "VIR", "DOSPUN", "SUM", "SUB", "MUL",
		"DIV", "COMA", "IF", "WHILE", "LET", "THEN", "ELSE", "IN", "DO", "BEGIN",
		"END", "CONST", "VAR", "ID", "NUM", "LETTER", "DIGIT", "LINE_COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 169, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 3, 22, 132, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 138,
		8, 22, 10, 22, 12, 22, 141, 9, 22, 1, 23, 4, 23, 144, 8, 23, 11, 23, 12,
		23, 145, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26,
		156, 8, 26, 10, 26, 12, 26, 159, 9, 26, 1, 26, 1, 26, 1, 27, 4, 27, 164,
		8, 27, 11, 27, 12, 27, 165, 1, 27, 1, 27, 0, 0, 28, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 0, 51, 0, 53, 25, 55, 26, 1, 0, 3, 2, 0, 65, 90, 97, 122,
		2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 173, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 59, 1, 0, 0,
		0, 5, 62, 1, 0, 0, 0, 7, 64, 1, 0, 0, 0, 9, 66, 1, 0, 0, 0, 11, 68, 1,
		0, 0, 0, 13, 70, 1, 0, 0, 0, 15, 72, 1, 0, 0, 0, 17, 74, 1, 0, 0, 0, 19,
		76, 1, 0, 0, 0, 21, 78, 1, 0, 0, 0, 23, 80, 1, 0, 0, 0, 25, 83, 1, 0, 0,
		0, 27, 89, 1, 0, 0, 0, 29, 93, 1, 0, 0, 0, 31, 98, 1, 0, 0, 0, 33, 103,
		1, 0, 0, 0, 35, 106, 1, 0, 0, 0, 37, 109, 1, 0, 0, 0, 39, 115, 1, 0, 0,
		0, 41, 119, 1, 0, 0, 0, 43, 125, 1, 0, 0, 0, 45, 131, 1, 0, 0, 0, 47, 143,
		1, 0, 0, 0, 49, 147, 1, 0, 0, 0, 51, 149, 1, 0, 0, 0, 53, 151, 1, 0, 0,
		0, 55, 163, 1, 0, 0, 0, 57, 58, 5, 59, 0, 0, 58, 2, 1, 0, 0, 0, 59, 60,
		5, 58, 0, 0, 60, 61, 5, 61, 0, 0, 61, 4, 1, 0, 0, 0, 62, 63, 5, 40, 0,
		0, 63, 6, 1, 0, 0, 0, 64, 65, 5, 41, 0, 0, 65, 8, 1, 0, 0, 0, 66, 67, 5,
		126, 0, 0, 67, 10, 1, 0, 0, 0, 68, 69, 5, 58, 0, 0, 69, 12, 1, 0, 0, 0,
		70, 71, 5, 43, 0, 0, 71, 14, 1, 0, 0, 0, 72, 73, 5, 45, 0, 0, 73, 16, 1,
		0, 0, 0, 74, 75, 5, 42, 0, 0, 75, 18, 1, 0, 0, 0, 76, 77, 5, 47, 0, 0,
		77, 20, 1, 0, 0, 0, 78, 79, 5, 44, 0, 0, 79, 22, 1, 0, 0, 0, 80, 81, 5,
		105, 0, 0, 81, 82, 5, 102, 0, 0, 82, 24, 1, 0, 0, 0, 83, 84, 5, 119, 0,
		0, 84, 85, 5, 104, 0, 0, 85, 86, 5, 105, 0, 0, 86, 87, 5, 108, 0, 0, 87,
		88, 5, 101, 0, 0, 88, 26, 1, 0, 0, 0, 89, 90, 5, 108, 0, 0, 90, 91, 5,
		101, 0, 0, 91, 92, 5, 116, 0, 0, 92, 28, 1, 0, 0, 0, 93, 94, 5, 116, 0,
		0, 94, 95, 5, 104, 0, 0, 95, 96, 5, 101, 0, 0, 96, 97, 5, 110, 0, 0, 97,
		30, 1, 0, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100, 5, 108, 0, 0, 100, 101,
		5, 115, 0, 0, 101, 102, 5, 101, 0, 0, 102, 32, 1, 0, 0, 0, 103, 104, 5,
		105, 0, 0, 104, 105, 5, 110, 0, 0, 105, 34, 1, 0, 0, 0, 106, 107, 5, 100,
		0, 0, 107, 108, 5, 111, 0, 0, 108, 36, 1, 0, 0, 0, 109, 110, 5, 98, 0,
		0, 110, 111, 5, 101, 0, 0, 111, 112, 5, 103, 0, 0, 112, 113, 5, 105, 0,
		0, 113, 114, 5, 110, 0, 0, 114, 38, 1, 0, 0, 0, 115, 116, 5, 101, 0, 0,
		116, 117, 5, 110, 0, 0, 117, 118, 5, 100, 0, 0, 118, 40, 1, 0, 0, 0, 119,
		120, 5, 99, 0, 0, 120, 121, 5, 111, 0, 0, 121, 122, 5, 110, 0, 0, 122,
		123, 5, 115, 0, 0, 123, 124, 5, 116, 0, 0, 124, 42, 1, 0, 0, 0, 125, 126,
		5, 118, 0, 0, 126, 127, 5, 97, 0, 0, 127, 128, 5, 114, 0, 0, 128, 44, 1,
		0, 0, 0, 129, 132, 5, 95, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0, 0,
		0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 139, 3, 49, 24, 0,
		134, 138, 3, 49, 24, 0, 135, 138, 3, 51, 25, 0, 136, 138, 5, 95, 0, 0,
		137, 134, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138,
		141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 46, 1,
		0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 144, 3, 51, 25, 0, 143, 142, 1, 0,
		0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0,
		146, 48, 1, 0, 0, 0, 147, 148, 7, 0, 0, 0, 148, 50, 1, 0, 0, 0, 149, 150,
		2, 48, 57, 0, 150, 52, 1, 0, 0, 0, 151, 152, 5, 47, 0, 0, 152, 153, 5,
		47, 0, 0, 153, 157, 1, 0, 0, 0, 154, 156, 8, 1, 0, 0, 155, 154, 1, 0, 0,
		0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 6, 26, 0, 0, 161, 54,
		1, 0, 0, 0, 162, 164, 7, 2, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0,
		167, 168, 6, 27, 0, 0, 168, 56, 1, 0, 0, 0, 7, 0, 131, 137, 139, 145, 157,
		165, 1, 6, 0, 0,
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

// goScannerInit initializes any static state used to implement goScanner. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewgoScanner(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoScannerInit() {
	staticData := &GoScannerLexerStaticData
	staticData.once.Do(goscannerLexerInit)
}

// NewgoScanner produces a new lexer instance for the optional input antlr.CharStream.
func NewgoScanner(input antlr.CharStream) *goScanner {
	GoScannerInit()
	l := new(goScanner)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GoScannerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "goScanner.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// goScanner tokens.
const (
	goScannerPyCOMA       = 1
	goScannerASSIGN       = 2
	goScannerPIZQ         = 3
	goScannerPDER         = 4
	goScannerVIR          = 5
	goScannerDOSPUN       = 6
	goScannerSUM          = 7
	goScannerSUB          = 8
	goScannerMUL          = 9
	goScannerDIV          = 10
	goScannerCOMA         = 11
	goScannerIF           = 12
	goScannerWHILE        = 13
	goScannerLET          = 14
	goScannerTHEN         = 15
	goScannerELSE         = 16
	goScannerIN           = 17
	goScannerDO           = 18
	goScannerBEGIN        = 19
	goScannerEND          = 20
	goScannerCONST        = 21
	goScannerVAR          = 22
	goScannerID           = 23
	goScannerNUM          = 24
	goScannerLINE_COMMENT = 25
	goScannerWS           = 26
)
