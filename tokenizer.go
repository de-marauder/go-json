package json

type tokenType int

const (
	STRING  tokenType = iota // 0
	NUMBER                   // 1
	BOOLEAN                  // 2
	NULL                     // 3

	RIGHT_BRACE   // 4
	LEFT_BRACE    // 5
	RIGHT_BRACKET // 6
	LEFT_BRACKET  // 7
	COMMA         // 8
	COLON         // 9
	EOF           // 10
	ESCAPE        // 11
)

type token struct {
	tokenType tokenType
	value     any
}

func newToken(tokenType tokenType, value any) token {
	return token{
		tokenType: tokenType,
		value:     value,
	}
}
