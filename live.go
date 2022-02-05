package main

import (
	"bytes"
	"context"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/c-bata/go-prompt"
	"github.com/mgnsk/gong/internal/ast"
	"github.com/mgnsk/gong/internal/constants"
	"github.com/mgnsk/gong/internal/interpreter"
)

var it = interpreter.New()

type shell struct {
	parser  prompt.ConsoleParser
	writer  prompt.ConsoleWriter
	results chan<- result
}

const (
	// Keycode for Ctrl+D.
	eot = 4

	defaultReso = 16

	gridBG    = "🟦"
	beatBG    = "⭕"
	currentBG = "🔴"
)

func printGrid(buf *bytes.Buffer, reso, current uint) {
	for i := uint(0); i < reso; i++ {
		if i == current {
			buf.WriteString(currentBG)
		} else if i%4 == 0 {
			buf.WriteString(beatBG)
		} else {
			buf.WriteString(gridBG)
		}
	}
}

func (s *shell) runMetronomePrinter(ctx context.Context, reso uint) {
	tickDuration := time.Duration(float64(time.Minute) / float64(it.Tempo()) / float64(constants.TicksPerQuarter))
	resoDuration := time.Duration(must(it.Parse(fmt.Sprintf("x%d", reso))).(ast.NoteList)[0].Length() * uint32(tickDuration))

	ticker := time.NewTicker(resoDuration)
	defer ticker.Stop()

	var buf bytes.Buffer

	i := uint(0)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			printGrid(&buf, reso, i%reso)

			s.writer.WriteRawStr("\r" + buf.String())
			s.writer.Flush()
			buf.Reset()

			i++
		}
	}
}

func (s *shell) runLive(reso uint) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go s.runMetronomePrinter(ctx, reso)

	for {
		b, err := s.parser.Read()
		if err != nil {
			return err
		}

		r, _ := utf8.DecodeRune(b)
		if r == eot {
			return nil
		}

		messages, err := it.NoteOn(r)
		if err != nil {
			// Ignore errors.
			continue
		}

		s.results <- result{"", messages}
	}
}

func (s *shell) handleInputLine(input string) error {
	if input == "live" {
		s.writer.WriteRawStr("Entered live mode. Press Ctrl+D to exit.\n")
		s.writer.Flush()

		if err := s.runLive(defaultReso); err != nil {
			return err
		}
		return nil
	}

	messages, err := it.Eval(input)
	if err != nil {
		return err
	}

	s.results <- result{"", messages}

	return nil
}

func must(res interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return res
}