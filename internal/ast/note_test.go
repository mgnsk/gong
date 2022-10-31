package ast_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestValidInputs(t *testing.T) {
	type (
		match    types.GomegaMatcher
		testcase struct {
			input string
			match match
		}
	)

	for _, tc := range []testcase{
		{
			"k",
			ContainSubstring("k"),
		},
		{
			"kk",
			ContainSubstring("k k"),
		},
		{
			"k k8",
			ContainSubstring("k k8"),
		},
		{
			"kk8.", // Properties apply only to the previous note symbol.
			ContainSubstring("k k8."),
		},
		{
			"[kk.]8", // Group properties apply to all notes in the group.
			ContainSubstring("k8 k8."),
		},
		{
			"[k.].", // Group properties are appended.
			ContainSubstring("k.."),
		},
		{
			"[k]",
			ContainSubstring("k"),
		},
		{
			"[k][k].",
			ContainSubstring("k k."),
		},
		{
			"kk[kk]kk[kk]kk",
			ContainSubstring("k k k k k k k k k k"),
		},
		{
			"[[k]]8",
			ContainSubstring("k8"),
		},
		{
			"k8kk16kkkk16",
			ContainSubstring("k8 k k16 k k k k16"),
		},
		{
			"k8 [kk]16 [kkkk]32",
			ContainSubstring("k8 k16 k16 k32 k32 k32 k32"),
		},
		{
			"-", // Pause.
			ContainSubstring("-"),
		},
		{
			"-8", // 8th pause.
			ContainSubstring("-8"),
		},
		{
			"k/3.#8",
			ContainSubstring("k#8./3"),
		},
		{
			"[[[[[k]/3].]#]8]^^", // Testing the ordering of properties.
			ContainSubstring("k#^^8./3"),
		},
		{
			"[[[[[k*]/3].]$].8]))", // Testing the ordering of properties.
			ContainSubstring("k$))8../3*"),
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			g := NewGomegaWithT(t)

			res, err := parse(tc.input)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(res).To(tc.match)
		})
	}
}

func TestInvalidProperties(t *testing.T) {
	for _, input := range []string{
		"k#$", // Sharp flat note.
		"k$#",
		"k^)", // Accentuated ghost note.
		"k)^",
	} {
		t.Run(input, func(t *testing.T) {
			g := NewGomegaWithT(t)

			_, err := parse(input)
			g.Expect(err).To(HaveOccurred())
		})
	}
}

func TestInvalidNoteValue(t *testing.T) {
	for _, input := range []string{
		"k3",
		"k22",
		"k0",
		"k129",
	} {
		t.Run(input, func(t *testing.T) {
			g := NewGomegaWithT(t)

			_, err := parse(input)
			g.Expect(err).To(HaveOccurred())
		})
	}
}

func TestForbiddenDuplicateProperty(t *testing.T) {
	for _, input := range []string{
		// TODO: allow double sharp and flat?
		"k44", // TODO: redefine bnf 1 | 2 | 4 | 8 etc
		"k##",
		"k$$",
		"k/3/3",
		"k**",
	} {
		t.Run(input, func(t *testing.T) {
			g := NewGomegaWithT(t)

			_, err := parse(input)
			g.Expect(err).To(HaveOccurred())
		})
	}
}
