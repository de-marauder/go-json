package gojson

type JsonValue any
type JsonObject map[string]JsonValue
type JsonArray []JsonValue

func MustParse(json string) JsonValue {
	var scanner Scanner = newScanner(json)
	tokens := scanner.scan()
	var parser Parser = newParser(tokens)
	return parser.parse()
}

// func Stringify(object JsonValue) (string, error) {
// 	switch object.(type) {
// 	case int, string:
// 		return fmt.Sprint(object), nil
// 	case []any:
// 		return fmt.Sprint(object), nil
// 	}
// 	return "", NewError(fmt.Sprint("Invalid Type ", reflect.TypeOf(object), ". Failed to stringify"))
// }
