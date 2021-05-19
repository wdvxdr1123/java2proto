package internal

import (
	"unicode/utf8"
)

type source struct {
	buf       []byte // source buf
	b, r, e   int    // begin read end
	line, col uint
	ch        rune
	chw       int
}

func (s *source) init(src []byte) {
	s.buf = src
	s.b, s.r, s.e = -1, 0, len(src)
	s.ch, s.chw = -1, 0
	s.nextCh()
}

func (s *source) start()          { s.b = s.r - s.chw }
func (s *source) stop()           { s.b = -1 }
func (s *source) segment() []byte { return s.buf[s.b : s.r-s.chw] }

// starting points for line and column numbers
const linebase = 1
const colbase = 1

// pos returns the (line, col) source position of s.ch.
func (s *source) pos() (line, col uint) {
	return linebase + s.line, colbase + s.col
}

func (s *source) rewind() {
	if s.b < 0 {
		panic("no active segment")
	}
	s.col -= uint(s.r - s.b)
	s.r = s.b
	s.nextCh()
}

func (s *source) nextCh() {
	s.col += uint(s.chw)
	if s.ch == '\n' {
		s.line++
		s.col = 0
	}

	if s.r == s.e { // EOF
		s.ch = -1
		s.chw = 0
		return
	}

	if s.ch = rune(s.buf[s.r]); s.ch < utf8.RuneSelf { // is a ASCII
		s.r++
		s.chw = 1
		return
	}

	s.ch, s.chw = utf8.DecodeRune(s.buf[s.r:s.e])
	s.r += s.chw
}
