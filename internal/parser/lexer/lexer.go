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
24: '['
25: ']'
26: '-'
27: '#'
28: '$'
29: '^'
30: ')'
31: '.'
32: '/'
33: '3'
34: '/'
35: '5'
36: '*'
37: ':'
38: 't'
39: 'e'
40: 'm'
41: 'p'
42: 'o'
43: ':'
44: 't'
45: 'i'
46: 'm'
47: 'e'
48: 's'
49: 'i'
50: 'g'
51: ':'
52: 'v'
53: 'e'
54: 'l'
55: 'o'
56: 'c'
57: 'i'
58: 't'
59: 'y'
60: ':'
61: 'c'
62: 'h'
63: 'a'
64: 'n'
65: 'n'
66: 'e'
67: 'l'
68: ':'
69: 'p'
70: 'r'
71: 'o'
72: 'g'
73: 'r'
74: 'a'
75: 'm'
76: ':'
77: 'c'
78: 'o'
79: 'n'
80: 't'
81: 'r'
82: 'o'
83: 'l'
84: ':'
85: 's'
86: 't'
87: 'a'
88: 'r'
89: 't'
90: ':'
91: 's'
92: 't'
93: 'o'
94: 'p'
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
