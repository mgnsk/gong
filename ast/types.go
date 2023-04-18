package ast

import (
	"fmt"

	"github.com/mgnsk/balafon/internal/parser/token"
)

// Property types.
var (
	// typeLetter    = mustGetType("letter")
	// typeStringLit = mustGetType("stringLit")
	typeSharp   = mustGetType("sharp")
	typeFlat    = mustGetType("flat")
	typeAccent  = mustGetType("accent")
	typeGhost   = mustGetType("ghost")
	typeUint    = mustGetType("uint")
	typeDot     = mustGetType("dot")
	typeTuplet  = mustGetType("tuplet")
	typeLetRing = mustGetType("letRing")
)

func mustGetType(tok string) token.Type {
	t := token.TokMap.Type(tok)
	if t == token.INVALID {
		panic(fmt.Sprintf("invalid token %s", tok))
	}
	return t
}
