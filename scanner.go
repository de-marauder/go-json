package json

import (
	"fmt"
	"strconv"
)

type Scanner interface {
	scan() []token
	isAtEnd() bool
	scanToken()
	advance() byte
	addString()
	addNumber()
	addKeyword()
	peek() byte
	peekNext() byte
}

type scanner struct {
	tokens  []token
	source  string
	current int
	start   int
	line    int
}

func newScanner(input string) *scanner {
	return &scanner{
		tokens:  []token{},
		source:  input,
		current: 0,
		start:   0,
		line:    1,
	}
}

func (s *scanner) scan() []token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, newToken(EOF, nil))
	return s.tokens
}
func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) scanToken() {
	b := s.advance()
	c := string(b)
	switch c {
	case "{":
		s.tokens = append(s.tokens, newToken(LEFT_BRACE, c))
	case "}":
		s.tokens = append(s.tokens, newToken(RIGHT_BRACE, c))
	case "[":
		s.tokens = append(s.tokens, newToken(LEFT_BRACKET, c))
	case "]":
		s.tokens = append(s.tokens, newToken(RIGHT_BRACKET, c))
	case ",":
		s.tokens = append(s.tokens, newToken(COMMA, c))
	case ":":
		s.tokens = append(s.tokens, newToken(COLON, c))
	case "\n":
		s.line += 1
	case "\"":
		s.addString()
	case "-":
		if isDigit(s.peek()) {
			s.advance()
			s.addNumber()
		} else {
			logErrorAndFail(CustomError{"Invalid syntax. '-' must be followed by digit"})
		}
	case " ":
		break
	default:
		if isDigit(b) {
			s.addNumber()
		} else if isAlpha(b) {
			s.addKeyword()
		} else {
			logErrorAndFail(CustomError{fmt.Sprintf("Unexpected Token %v at line %v", c, s.line)})
		}

	}
}

func (s *scanner) advance() byte {
	if !s.isAtEnd() {
		c := s.source[s.current]
		s.current += 1
		return c
	}
	return s.source[len(s.source)-1]
}

func (s *scanner) peek() byte {
	if s.isAtEnd() {
		return byte(0)
	}
	return s.source[s.current]
}

func (s *scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return byte(0)
	}
	return s.source[s.current+1]
}

func (s *scanner) addString() {
	s.advance()
	for string(s.peek()) != "\"" {
		if s.isAtEnd() {
			logErrorAndFail(CustomError{fmt.Sprint("Unterminated string at line ", (s.line))})
		}

		s.advance()
	}

	s.advance()
	s.tokens = append(s.tokens, newToken(STRING, string(s.source[s.start+1:s.current-1])))
}

func (s *scanner) addNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if string(s.peek()) == "." {
		if !isDigit(s.peekNext()) {
			logErrorAndFail(CustomError{fmt.Sprint("Invalid token. Digit must follow '.' on line ", s.line)})
		}

		// consume '.'
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}

		s.advance()
		f64, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
		logErrorAndFail(err)
		s.tokens = append(s.tokens, newToken(NUMBER, f64))

	} else {
		i, err := strconv.Atoi(s.source[s.start:s.current])
		logErrorAndFail(err)
		s.tokens = append(s.tokens, newToken(NUMBER, i))
		// s.advance()
	}
}

func (s *scanner) addKeyword() {
	for isAlpha(s.peek()) {
		s.advance()
	}
	keyword := s.source[s.start:s.current]
	switch keyword {
	case "true":
		s.tokens = append(s.tokens, newToken(BOOLEAN, true))
	case "false":
		s.tokens = append(s.tokens, newToken(BOOLEAN, false))
	case "null":
		s.tokens = append(s.tokens, newToken(NULL, nil))
	default:
		s.tokens = append(s.tokens, newToken(STRING, keyword))
	}

}
