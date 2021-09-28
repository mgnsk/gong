// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/mgnsk/gong/internal/token"
)

const (
	NoState    = -1
	NumStates  = 58
	NumSymbols = 54
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
0: '.'
1: '/'
2: '='
3: 'b'
4: 'a'
5: 'r'
6: 'e'
7: 'n'
8: 'd'
9: 'p'
10: 'l'
11: 'a'
12: 'y'
13: 't'
14: 'e'
15: 'm'
16: 'p'
17: 'o'
18: 'c'
19: 'h'
20: 'a'
21: 'n'
22: 'n'
23: 'e'
24: 'l'
25: 'p'
26: 'r'
27: 'o'
28: 'g'
29: 'r'
30: 'a'
31: 'm'
32: 'c'
33: 'o'
34: 'n'
35: 't'
36: 'r'
37: 'o'
38: 'l'
39: '-'
40: '0'
41: '"'
42: '"'
43: '#'
44: '\n'
45: ' '
46: '\t'
47: '\n'
48: '\r'
49: '0'-'9'
50: 'a'-'z'
51: 'A'-'Z'
52: '1'-'9'
53: .
*/
