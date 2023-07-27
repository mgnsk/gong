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
		String: `RepeatTerminator : empty	<<  >>`,
		Id:         "RepeatTerminator",
		NTType:     2,
		Index:      2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `RepeatTerminator : terminator RepeatTerminator	<<  >>`,
		Id:         "RepeatTerminator",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DeclList : Decl terminator RepeatTerminator DeclList	<< ast.NewNodeList(X[0].(ast.Node), X[3].(ast.NodeList)), nil >>`,
		Id:         "DeclList",
		NTType:     3,
		Index:      4,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNodeList(X[0].(ast.Node), X[3].(ast.NodeList)), nil
		},
	},
	ProdTabEntry{
		String: `DeclList : Decl RepeatTerminator	<< ast.NewNodeList(X[0].(ast.Node), nil), nil >>`,
		Id:         "DeclList",
		NTType:     3,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNodeList(X[0].(ast.Node), nil), nil
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
		String: `NoteList : NoteObject	<< ast.NewNoteList(X[0].(ast.Node), nil), nil >>`,
		Id:         "NoteList",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNoteList(X[0].(ast.Node), nil), nil
		},
	},
	ProdTabEntry{
		String: `NoteList : NoteObject NoteList	<< ast.NewNoteList(X[0].(ast.Node), X[1].(ast.NoteList)), nil >>`,
		Id:         "NoteList",
		NTType:     6,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNoteList(X[0].(ast.Node), X[1].(ast.NoteList)), nil
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
		String: `NoteObject : bracketBegin NoteList bracketEnd PropertyList	<< ast.NewNoteListFromGroup(X[1].(ast.NoteList), X[3].(ast.PropertyList)) >>`,
		Id:         "NoteObject",
		NTType:     7,
		Index:      14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNoteListFromGroup(X[1].(ast.NoteList), X[3].(ast.PropertyList))
		},
	},
	ProdTabEntry{
		String: `NoteSymbol : symbol	<<  >>`,
		Id:         "NoteSymbol",
		NTType:     8,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NoteSymbol : rest	<<  >>`,
		Id:         "NoteSymbol",
		NTType:     8,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PropertyList : empty	<< ast.PropertyList(nil), nil >>`,
		Id:         "PropertyList",
		NTType:     9,
		Index:      17,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.PropertyList(nil), nil
		},
	},
	ProdTabEntry{
		String: `PropertyList : Property PropertyList	<< ast.NewPropertyList(X[0].(*token.Token), X[1]) >>`,
		Id:         "PropertyList",
		NTType:     9,
		Index:      18,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewPropertyList(X[0].(*token.Token), X[1])
		},
	},
	ProdTabEntry{
		String: `Property : propSharp	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propFlat	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propStaccato	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propAccent	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propMarcato	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propGhost	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : uint	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propDot	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propTuplet	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Property : propLetRing	<<  >>`,
		Id:         "Property",
		NTType:     10,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Command : cmdAssign symbol uint	<< ast.NewCmdAssign(X[0].(*token.Token).Pos, []rune(string(X[1].(*token.Token).Lit))[0], ast.Must(X[2].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdAssign(X[0].(*token.Token).Pos, []rune(string(X[1].(*token.Token).Lit))[0], ast.Must(X[2].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdPlay	<< ast.NewCmdPlay(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit[len(":play "):])) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdPlay(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit[len(":play "):]))
		},
	},
	ProdTabEntry{
		String: `Command : cmdTempo uint	<< ast.NewCmdTempo(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      31,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdTempo(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdTimesig uint uint	<< ast.NewCmdTimeSig(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdTimeSig(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdVelocity uint	<< ast.NewCmdVelocity(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      33,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdVelocity(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdChannel uint	<< ast.NewCmdChannel(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      34,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdChannel(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdProgram uint	<< ast.NewCmdProgram(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      35,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdProgram(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdControl uint uint	<< ast.NewCmdControl(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value())) >>`,
		Id:         "Command",
		NTType:     11,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCmdControl(X[0].(*token.Token).Pos, ast.Must(X[1].(*token.Token).Int64Value()), ast.Must(X[2].(*token.Token).Int64Value()))
		},
	},
	ProdTabEntry{
		String: `Command : cmdStart	<< ast.CmdStart{}, nil >>`,
		Id:         "Command",
		NTType:     11,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.CmdStart{}, nil
		},
	},
	ProdTabEntry{
		String: `Command : cmdStop	<< ast.CmdStop{}, nil >>`,
		Id:         "Command",
		NTType:     11,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.CmdStop{}, nil
		},
	},
	ProdTabEntry{
		String: `Comment : blockComment	<< ast.NewBlockComment(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit)), nil >>`,
		Id:         "Comment",
		NTType:     12,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewBlockComment(X[0].(*token.Token).Pos, string(X[0].(*token.Token).Lit)), nil
		},
	},
}
