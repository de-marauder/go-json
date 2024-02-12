package gojson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name   string
	expect any
	input  []token
}

var test1 = []testCase{
	{
		name: "Can create a token", expect: "value", input: []token{
			newToken(LEFT_BRACE, "{"),
			newToken(STRING, "key"),
			newToken(COLON, ":"),
			newToken(STRING, "value"),
			newToken(RIGHT_BRACE, "}"),
		},
	},
}

func TestParser_parseObjectTokens(t *testing.T) {
	tc := test1[0]
	t.Run(tc.name, func(t *testing.T) {
		assert := assert.New(t)

		parser := newParser(tc.input)
		j := parser.parse()

		got := j.(JsonObject)["key"]

		assert.Equal(tc.expect, got, fmt.Sprintf("Parse result incorrect. Expected %v got %v\n", tc.expect, got))
	})
}
