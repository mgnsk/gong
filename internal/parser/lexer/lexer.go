// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/mgnsk/gong/internal/parser/token"
)

const (
	NoState    = -1
	NumStates  = 76
	NumSymbols = 84
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
	src, err := ioutil.ReadFile(fpath)
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
0: '-'
1: '"'
2: '"'
3: '#'
4: '$'
5: '^'
6: ')'
7: '.'
8: '/'
9: '*'
10: '['
11: ']'
12: 'a'
13: 's'
14: 's'
15: 'i'
16: 'g'
17: 'n'
18: 't'
19: 'e'
20: 'm'
21: 'p'
22: 'o'
23: 'c'
24: 'h'
25: 'a'
26: 'n'
27: 'n'
28: 'e'
29: 'l'
30: 'v'
31: 'e'
32: 'l'
33: 'o'
34: 'c'
35: 'i'
36: 't'
37: 'y'
38: 'p'
39: 'r'
40: 'o'
41: 'g'
42: 'r'
43: 'a'
44: 'm'
45: 'c'
46: 'o'
47: 'n'
48: 't'
49: 'r'
50: 'o'
51: 'l'
52: 'b'
53: 'a'
54: 'r'
55: 'e'
56: 'n'
57: 'd'
58: 'p'
59: 'l'
60: 'a'
61: 'y'
62: 's'
63: 't'
64: 'a'
65: 'r'
66: 't'
67: 's'
68: 't'
69: 'o'
70: 'p'
71: '0'
72: '/'
73: '/'
74: '\n'
75: ' '
76: '\t'
77: '\n'
78: '\r'
79: 'a'-'z'
80: 'A'-'Z'
81: '1'-'9'
82: '0'-'9'
83: .
*/
