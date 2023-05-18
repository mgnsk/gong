// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"os"
	"unicode/utf8"

	"github.com/mgnsk/balafon/internal/parser/token"
)

const (
	NoState    = -1
	NumStates  = 95
	NumSymbols = 113
)

type Lexer struct {
	src     []byte
	pos     int
	line    int
	column  int
	Context token.Context
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:     src,
		pos:     0,
		line:    1,
		column:  1,
		Context: nil,
	}
	return lexer
}

// SourceContext is a simple instance of a token.Context which
// contains the name of the source file.
type SourceContext struct {
	Filepath string
}

func (s *SourceContext) Source() string {
	return s.Filepath
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	lexer := NewLexer(src)
	lexer.Context = &SourceContext{Filepath: fpath}
	return lexer, nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = &token.Token{}
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		tok.Pos.Context = l.Context
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn
	tok.Pos.Context = l.Context

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: ';'
1: '\n'
2: ';'
3: '\n'
4: ':'
5: 'b'
6: 'a'
7: 'r'
8: ':'
9: 'e'
10: 'n'
11: 'd'
12: ':'
13: 'p'
14: 'l'
15: 'a'
16: 'y'
17: ':'
18: 'a'
19: 's'
20: 's'
21: 'i'
22: 'g'
23: 'n'
24: ':'
25: 't'
26: 'e'
27: 'm'
28: 'p'
29: 'o'
30: ':'
31: 't'
32: 'i'
33: 'm'
34: 'e'
35: 's'
36: 'i'
37: 'g'
38: ':'
39: 'v'
40: 'e'
41: 'l'
42: 'o'
43: 'c'
44: 'i'
45: 't'
46: 'y'
47: ':'
48: 'c'
49: 'h'
50: 'a'
51: 'n'
52: 'n'
53: 'e'
54: 'l'
55: ':'
56: 'p'
57: 'r'
58: 'o'
59: 'g'
60: 'r'
61: 'a'
62: 'm'
63: ':'
64: 'c'
65: 'o'
66: 'n'
67: 't'
68: 'r'
69: 'o'
70: 'l'
71: ':'
72: 's'
73: 't'
74: 'a'
75: 'r'
76: 't'
77: ':'
78: 's'
79: 't'
80: 'o'
81: 'p'
82: '['
83: ']'
84: '-'
85: '#'
86: '$'
87: '^'
88: ')'
89: '.'
90: '/'
91: '3'
92: '/'
93: '5'
94: '*'
95: '0'
96: ' '
97: '/'
98: '/'
99: '\n'
100: '/'
101: '*'
102: '*'
103: '*'
104: '/'
105: ' '
106: '\t'
107: '\r'
108: '1'-'9'
109: '0'-'9'
110: 'a'-'z'
111: 'A'-'Z'
112: .
*/
