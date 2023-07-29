// Code generated by gocc; DO NOT EDIT.

package parser

import (
    "github.com/mgnsk/balafon/internal/ast"
    "github.com/mgnsk/balafon/internal/parser/token"
)

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : SourceFile	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SourceFile : RepeatTerminator DeclList	<< X[1], nil >>`,
		Id:         "SourceFile",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `RepeatTerminator : empty	<< ast.RepeatTerminator(nil), nil >>`,
		Id:         "RepeatTerminator",
		NTType:     2,
		Index:      2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.RepeatTerminator(nil), nil
		},
	},
	ProdTabEntry{
		String: `RepeatTerminator : terminator RepeatTerminator	<< ast.NewRepeatTerminator(string(X[0].(*token.Token).Lit), X[1].(ast.RepeatTerminator)...), nil >>`,
		Id:         "RepeatTerminator",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewRepeatTerminator(string(X[0].(*token.Token).Lit), X[1].(ast.RepeatTerminator)...), nil
		},
	},
	ProdTabEntry{
		String: `DeclList : Decl terminator RepeatTerminator DeclList	<< ast.NewNodeList(X[0].(ast.Node), append([]ast.Node{ast.RepeatTerminator{string(X[1].(*token.Token).Lit)}, X[2].(ast.RepeatTerminator)}, X[3].(ast.NodeList)...)...), nil >>`,
		Id:         "DeclList",
		NTType:     3,
		Index:      4,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNodeList(X[0].(ast.Node), append([]ast.Node{ast.RepeatTerminator{string(X[1].(*token.Token).Lit)}, X[2].(ast.RepeatTerminator)}, X[3].(ast.NodeList)...)...), nil
		},
	},
	ProdTabEntry{
		String: `DeclList : Decl RepeatTerminator	<< ast.NewNodeList(X[0].(ast.Node), X[1].(ast.RepeatTerminator)), nil >>`,
		Id:         "DeclList",
		NTType:     3,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNodeList(X[0].(ast.Node), X[1].(ast.RepeatTerminator)), nil
		},
	},
	ProdTabEntry{
		String: `Decl : Bar	<<  >>`,
		Id:         "Decl",
		NTType:     4,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Decl : Command	<<  >>`,
		Id:         "Decl",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Decl : NoteList	<<  >>`,
		Id:         "Decl",
		NTType:     4,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Decl : Comment	<<  >>`,
		Id:         "Decl",
		NTType:     4,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bar : cmdBar RepeatTerminator DeclList cmdEnd	<< ast.NewBar(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit[len(":bar "):]), X[2].(ast.NodeList)), nil >>`,
		Id:         "Bar",
		NTType:     5,
		Index:      10,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewBar(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit[len(":bar "):]), X[2].(ast.NodeList)), nil
		},
	},
	ProdTabEntry{
		String: `NoteList : NoteObject	<< ast.NewNodeList(X[0].(ast.Node)), nil >>`,
		Id:         "NoteList",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNodeList(X[0].(ast.Node)), nil
		},
	},
	ProdTabEntry{
		String: `NoteList : NoteObject NoteList	<< ast.NewNodeList(X[0].(ast.Node), X[1].(ast.NodeList)...), nil >>`,
		Id:         "NoteList",
		NTType:     6,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNodeList(X[0].(ast.Node), X[1].(ast.NodeList)...), nil
		},
	},
	ProdTabEntry{
		String: `NoteObject : NoteSymbol PropertyList	<< ast.NewNote(X[0].(*token.Token).Pos, []rune(string(X[0].(*token.Token).Lit))[0], X[1].(ast.PropertyList)), nil >>`,
		Id:         "NoteObject",
		NTType:     7,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNote(X[0].(*token.Token).Pos, []rune(string(X[0].(*token.Token).Lit))[0], X[1].(ast.PropertyList)), nil
		},
	},
	ProdTabEntry{
		String: `NoteObject : NoteGroup	<<  >>`,
		Id:         "NoteObject",
		NTType:     7,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NoteGroup : bracketBegin NoteList bracketEnd PropertyList	<< ast.NewNoteGroup(X[1].(ast.NodeList), X[3].(ast.PropertyList)) >>`,
		Id:         "NoteGroup",
		NTType:     8,
		Index:      15,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNoteGroup(X[1].(ast.NodeList), X[3].(ast.PropertyList))
		},
	},
	ProdTabEntry{
		String: `NoteSymbol : symbol	<<  >>`,
		Id:         "NoteSymbol",
		NTType:     9,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NoteSymbol : rest	<<  >>`,
		Id:         "NoteSymbol",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PropertyList : empty	<< ast.PropertyList(nil), nil >>`,
		Id:         "PropertyList",
		NTType:     10,
		Index:      18,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.PropertyList(nil), nil
		},
	},
	ProdTabEntry{
		String: `PropertyList : Property PropertyList	<< ast.NewPropertyList(X[0].(*token.Token), X[1]) >>`,
		Id:         "PropertyList",
		NTType:     10,
		Index:      19,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewPropertyList(X[0].(*token.Token), X[1])
		},
	},
	ProdTabEntry{
		String: `Property : propSharp	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propFlat	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propStaccato	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propAccent	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propMarcato	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propGhost	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : uint	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propDot	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propTuplet	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propLetRing	<<  >>`,
		Id:         "Property",
		NTType:     11,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Command : cmdAssign symbol uint	<< ast.NewCmdAssign(X[0].(*token.Token).Pos, []rune(string(X[1].(*token.Token).Lit))[0], ast.Must(X[2].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdAssign(X[0].(*token.Token).Pos, []rune(string(X[1].(*token.Token).Lit))[0], ast.Must(X[2].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdPlay	<< ast.NewCmdPlay(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit[len(":play "):])) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdPlay(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit[len(":play "):]))
		},
	},
	ProdTabEntry{
		String: `Command : cmdTempo uint	<< ast.NewCmdTempo(ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdTempo(ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdTimesig uint uint	<< ast.NewCmdTimeSig(ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdTimeSig(ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdVelocity uint	<< ast.NewCmdVelocity(ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      34,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdVelocity(ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdChannel uint	<< ast.NewCmdChannel(ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      35,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdChannel(ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdProgram uint	<< ast.NewCmdProgram(ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      36,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdProgram(ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdControl uint uint	<< ast.NewCmdControl(ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     12,
		Index:      37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdControl(ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdStart	<< ast.CmdStart{}, nil >>`,
		Id:         "Command",
		NTType:     12,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.CmdStart{}, nil
		},
	},
	ProdTabEntry{
		String: `Command : cmdStop	<< ast.CmdStop{}, nil >>`,
		Id:         "Command",
		NTType:     12,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.CmdStop{}, nil
		},
	},
	ProdTabEntry{
		String: `Comment : blockComment	<< ast.NewBlockComment(string(X[0].(*token.Token).Lit)), nil >>`,
		Id:         "Comment",
		NTType:     13,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewBlockComment(string(X[0].(*token.Token).Lit)), nil
		},
	},
}
