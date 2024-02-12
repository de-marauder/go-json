package json

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
	fmt.Println("Parsing from token => ", token)
	switch token.tokenType {
	case STRING, NUMBER, NULL, BOOLEAN:
		fmt.Println("Parsing from token regular type=> ", token)
		return token.value
	case LEFT_BRACE:
		fmt.Println("Parsing from token object type=> ", token)
		return p.parseObject()
	case LEFT_BRACKET:
		fmt.Println("Parsing from token array type=> ", token)
		return p.parseArray()
	default:
		fmt.Println("tokens = ", p.tokens)
		fmt.Println("current token = ", p.current, p.tokens[p.current])
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
	fmt.Println("previous value", fmt.Sprintln(value))
	fmt.Println("Initial value", fmt.Sprintln(p.peek().value))
	for value.tokenType != RIGHT_BRACKET {

		if value.tokenType == EOF {
			logErrorAndFail(CustomError{"Unterminated JSON Array"})
		}

		a = append(a, p.parseFromToken(value))
		fmt.Println("current => ", p.current)
		fmt.Println("tokens => ", p.tokens)
		fmt.Println("token => ", p.tokens[p.current])
		fmt.Println("tokenType => ", p.peek().tokenType)
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
