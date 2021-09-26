package scanner

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/mgnsk/gong/internal/ast"
	"github.com/mgnsk/gong/internal/constants"
	"github.com/mgnsk/gong/internal/lexer"
	"github.com/mgnsk/gong/internal/parser"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
)

// Message is a tempo or a MIDI message.
type Message struct {
	Tempo uint32
	Tick  uint64
	Msg   midi.Message
}

// Scanner scans messages from raw text input.
type Scanner struct {
	scanner  *bufio.Scanner
	parser   *parser.Parser
	err      error
	messages []Message

	notes       map[string]uint8
	bars        map[string][]*ast.Track
	barBuffer   []*ast.Track
	currentBar  string
	currentTick uint64
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (s *Scanner) Err() error {
	if s.err != nil {
		return s.err
	}
	return s.scanner.Err()
}

// Messages returns the currently accumulated messages.
func (s *Scanner) Messages() []Message {
	return s.messages
}

func (s *Scanner) Suggest() []string {
	// Suggest assigned notes at any time.
	var sug []string
	for note := range s.notes {
		sug = append(sug, note)
	}
	if s.currentBar != "" {
		// Suggest ending the current bar if we're in the middle of a bar.
		sug = append(sug, "end")
	} else {
		// Suggest commands.
		sug = append(sug, "bar", "tempo")
		// Suggest playing a bar.
		for name := range s.bars {
			sug = append(sug, "play "+name)
		}
	}
	return sug
}

// Scan the next batch of messages.
func (s *Scanner) Scan() bool {
	s.messages = nil

	for s.scanner.Scan() {
		if len(s.scanner.Bytes()) == 0 {
			continue
		}

		s.parser.Reset()

		res, err := s.parser.Parse(lexer.NewLexer(s.scanner.Bytes()))
		if err != nil {
			s.err = err
			return false
		}

		switch r := res.(type) {
		case *ast.Assignment:
			switch r.Name {
			case "tempo":
				s.messages = []Message{{
					Tempo: r.Value,
				}}
				return true

			default:
				if len(r.Name) != 1 {
					s.err = fmt.Errorf("invalid assignment to '%s', must be either 'tempo' or a single character note", r.Name)
					return false
				}
				// TODO out of range test uint8
				s.notes[r.Name] = uint8(r.Value)
			}

		case *ast.Track:
			if s.currentBar != "" {
				s.barBuffer = append(s.barBuffer, r)
			} else {
				messages, err := s.parseBar(r)
				if err != nil {
					s.err = err
					return false
				}
				s.messages = messages
				return true
			}

		case *ast.Command:
			switch r.Name {
			case "bar": // Begin a bar.
				if s.currentBar != "" {
					s.err = errors.New("cannot begin a bar: already in a bar")
					return false
				}
				if _, ok := s.bars[r.Arg]; ok {
					s.err = fmt.Errorf("bar '%s' already defined", r.Arg)
					return false
				}
				s.currentBar = r.Arg

			case "end": // End the current bar.
				if s.currentBar == "" {
					s.err = errors.New("cannot end a bar: no active bar")
					return false
				}
				s.bars[s.currentBar] = s.barBuffer
				s.currentBar = ""
				s.barBuffer = nil

			case "play": // Play a bar.
				if s.currentBar != "" {
					s.err = errors.New("cannot play: current bar not ended")
					return false
				}
				bar, ok := s.bars[r.Arg]
				if !ok {
					s.err = fmt.Errorf("cannot play nonexistent bar '%s'", r.Arg)
					return false
				}
				messages, err := s.parseBar(bar...)
				if err != nil {
					s.err = err
					return false
				}
				s.messages = messages
				return true

			default:
				s.err = fmt.Errorf("invalid command '%s', must be either 'bar', 'end' or 'play'", r.Name)
				return false
			}

		default:
			panic("invalid token")
		}
	}

	return false
}

func (s *Scanner) parseBar(tracks ...*ast.Track) ([]Message, error) {
	var messages []Message

	for _, track := range tracks {
		var tick uint64
		for _, note := range track.Notes {
			length := s.noteLength(note)

			if note.Name != "-" {
				key, ok := s.notes[note.Name]
				if !ok {
					return nil, fmt.Errorf("key '%s' undefined", note.Name)
				}

				messages = append(messages, Message{
					Tick: s.currentTick + tick,
					// TODO velocity and channel
					Msg: channel.Channel0.NoteOn(key, 50),
				})

				messages = append(messages, Message{
					Tick: s.currentTick + tick + uint64(length),
					Msg:  channel.Channel0.NoteOff(key),
				})
			}

			tick += uint64(length)
		}
	}

	// Sort the messages so that every note is off before on.
	sort.Slice(messages, func(i, j int) bool {
		if messages[i].Tick < messages[j].Tick {
			return true
		} else if messages[i].Tick == messages[j].Tick {
			if a, ok := messages[i].Msg.(channel.NoteOff); ok {
				if b, ok := messages[j].Msg.(channel.NoteOff); ok {
					// When both are NoteOff, sort by key.
					return a.Key() < b.Key()
				}
				// NoteOff before any other messages on the same tick.
				return true
			}
		}
		return false
	})

	s.currentTick = messages[len(messages)-1].Tick

	return messages, nil
}

func (s *Scanner) noteLength(note *ast.Note) uint16 {
	value := note.Value()
	length := 4 * constants.TicksPerQuarter / uint16(value)
	if note.IsDot() {
		length += (length / 2)
	}
	if division := note.Tuplet(); division > 0 {
		length = uint16(float64(length) * 2.0 / float64(division))
	}
	return length
}

// New creates a new Scanner instance.
func New(r io.Reader) *Scanner {
	return &Scanner{
		scanner: bufio.NewScanner(r),
		parser:  parser.NewParser(),
		notes:   make(map[string]uint8),
		bars:    make(map[string][]*ast.Track),
	}
}
