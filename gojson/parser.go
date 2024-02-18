package gojson

import "fmt"

type Parser interface {
	parse() JsonValue
	parseFromToken(token token) JsonValue
	parseObject() JsonObject
	parseArray() JsonArray
	peek() token
	advance() token
	consume(tokenType tokenType, errorMessage string)
	consumeCommaUnless(tokenType tokenType, errorMessage string)
}

type parser struct {
	tokens  []token
	current int
}

func newParser(tokens []token) *parser {
	return &parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *parser) parse() JsonValue {
	token := p.advance()
	return p.parseFromToken(token)
}

func (p *parser) parseFromToken(token token) JsonValue {
	switch token.tokenType {
	case STRING, NUMBER, NULL, BOOLEAN, EOF:
		return token.value
	case LEFT_BRACE:
		return p.parseObject()
	case LEFT_BRACKET:
		return p.parseArray()
	default:
		logErrorAndFail(CustomError{fmt.Sprintf("Invalid token %v\n", token)})
	}
	return ""
}

func (p *parser) parseObject() JsonObject {
	o := make(map[string]JsonValue)

	key := p.advance()

	for key.tokenType != RIGHT_BRACE {
		if key.tokenType == EOF {
			logErrorAndFail(CustomError{fmt.Sprintln("Unterminate JSON object", key.value, p.tokens)})
		}
		if key.tokenType != STRING {
			logErrorAndFail(CustomError{fmt.Sprintf("JSON Keys must be type string got %v\n", key.value)})
		}

		// check colon
		p.consume(COLON, "JSON Object keya nd values must be separated by ':'")

		// check value
		value := p.advance()
		o[key.value.(string)] = p.parseFromToken(value)

		p.consumeCommaUnless(RIGHT_BRACE, fmt.Sprintln("Entries must be separated by ',' unless end of array '}' is reached"))
		key = p.advance()
	}
	return o
}

func (p *parser) parseArray() JsonArray {
	a := []JsonValue{}
	value := p.advance()
	for value.tokenType != RIGHT_BRACKET {

		if value.tokenType == EOF {
			logErrorAndFail(CustomError{"Unterminated JSON Array"})
		}

		a = append(a, p.parseFromToken(value))
		p.consumeCommaUnless(RIGHT_BRACKET, fmt.Sprintln("Entries must be separated by ',' unless end of array ']' is reached"))
		value = p.advance()
	}
	return a
}

func (p *parser) peek() token {
	return p.tokens[p.current]
}

func (p *parser) advance() token {
	t := p.tokens[p.current]
	if p.current < len(p.tokens) {
		p.current += 1
	}
	return t
}

func (p *parser) consume(tokenType tokenType, errorMessage string) {
	if p.peek().tokenType != tokenType {
		logErrorAndFail(CustomError{errorMessage})
	}
	p.advance()
}

func (p *parser) consumeCommaUnless(tokenType tokenType, errorMessage string) {

	if p.peek().tokenType == COMMA {
		p.advance()
	} else if p.peek().tokenType != tokenType {

		logErrorAndFail(CustomError{errorMessage})
	}
}
