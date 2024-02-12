package json_test

import (
	"fmt"
	"json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tesctCase1 []struct {
	name    string
	expect  any
	jsonStr string
}

var test1 = tesctCase1{
	{name: "Can parse single digit", expect: 1, jsonStr: "1"},
	{name: "Can parse multiple digits", expect: 1321287549, jsonStr: "1321287549"},
	{name: "Can parse decimal digits", expect: 13212.87549, jsonStr: "13212.87549"},
	{name: "Can parse single letter", expect: "a", jsonStr: "a"},
	{name: "Can parse multiple letters", expect: "abvuwjdskmd", jsonStr: "abvuwjdskmd"},
	{name: "Can parse complex string", expect: "vdsew'cdswc\n\rcawdwq:-2143vfssxw", jsonStr: "\"vdsew'cdswc\n\rcawdwq:-2143vfssxw\""},
	{name: "Can parse null", expect: nil, jsonStr: "null"},
	{name: "Can parse boolean true", expect: true, jsonStr: "true"},
	{name: "Can parse boolean false", expect: false, jsonStr: "false"},
}

func TestMustParse(t *testing.T) {
	for _, tc := range test1 {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			got := json.MustParse(tc.jsonStr)
			assert.Equal(tc.expect, got, fmt.Sprintf("Expected %v got %v", tc.expect, got))

		})
	}
}

var test2 = tesctCase1{
	{
		name: "can parse object with one key string and value string pair",
		expect: json.JsonObject{
			"name": "value",
		},
		jsonStr: "{\"name\": \"value\"}",
	},
	{
		name: "can parse complex object",
		expect: json.JsonObject{
			"name":   "value",
			"arrKey": json.JsonArray{"a", 3, "w"},
			"objKey": json.JsonObject{
				"nested key": "nested value",
			},
		},
		jsonStr: "{\"name\": \"value\", \"arrKey\": [\"a\",3,\"w\"], \"objKey\": {\"nested key\": \"nested value\"} }",
	},
}

func TestMustParse_simpleObject(t *testing.T) {
	tc := test2[0]
	t.Run(tc.name, func(t *testing.T) {
		assert := assert.New(t)
		got := json.MustParse(tc.jsonStr)
		assert.Equal(tc.expect.(json.JsonObject)["name"], got.(json.JsonObject)["name"], fmt.Sprintf("Expected %v got %v", tc.expect, got))
	})
}
func TestMustParse_complexObject(t *testing.T) {
	tc := test2[1]
	t.Run(tc.name, func(t *testing.T) {
		assert := assert.New(t)
		got := json.MustParse(tc.jsonStr).(json.JsonObject)
		assert.Equal(tc.expect.(json.JsonObject)["name"], got["name"], fmt.Sprintf("Expected %v got %v", tc.expect, got))
		assert.Equal(tc.expect.(json.JsonObject)["arrKey"].(json.JsonArray)[0], got["arrKey"].(json.JsonArray)[0], fmt.Sprintf("Expected %v got %v", tc.expect, got))
		assert.Equal(tc.expect.(json.JsonObject)["objKey"], got["objKey"], fmt.Sprintf("Expected %v got %v", tc.expect, got))
	})
}

var test3 = tesctCase1{
	{
		name:    "can parse empty array",
		expect:  []int{},
		jsonStr: "[]",
	},
	{
		name:    "can parse array with single integer",
		expect:  []int{1},
		jsonStr: "[1]",
	},
	{
		name:    "can parse array with multiple integers",
		expect:  []int{1, 2, 3},
		jsonStr: "[1,2,3]",
	},
	{
		name: "can parse  complex array ",
		expect: json.JsonArray{1, json.JsonObject{
			"name":   "value",
			"arrKey": json.JsonArray{"a", 3, "w"},
			"objKey": json.JsonObject{
				"nested key": "nested value",
			},
		}, 3},
		jsonStr: "[1, {\"name\": \"value\", \"arrKey\": [\"a\",3,\"w\"], \"objKey\": {\"nested key\": \"nested value\"} }, 3]",
	},
}

func TestMustParse_emptyArray(t *testing.T) {
	tc := test3[0]
	t.Run(tc.name, func(t *testing.T) {
		assert := assert.New(t)

		got := json.MustParse(tc.jsonStr).(json.JsonArray)

		assert.Equal(0, len(got), fmt.Sprintf("Expected %v got %v", tc.expect, got))
	})
}

func TestMustParse_array(t *testing.T) {
	for _, tc := range test3[1:3] {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			got := json.MustParse(tc.jsonStr).(json.JsonArray)

			assert.Equal(tc.expect.([]int)[0], got[0], fmt.Sprintf("Expected %v got %v", tc.expect, got))
			assert.Equal(len(tc.expect.([]int)), len(got), fmt.Sprintf("Expected %v got %v", tc.expect, got))
			assert.Equal(
				tc.expect.([]int)[len(tc.expect.([]int))-1],
				got[len(got)-1],
				fmt.Sprintf("Expected %v got %v", tc.expect, got))
		})
	}
}
func TestMustParse_complexArray(t *testing.T) {
	tc := test3[3]
	t.Run(tc.name, func(t *testing.T) {
		assert := assert.New(t)

		got := json.MustParse(tc.jsonStr).(json.JsonArray)[1].(json.JsonObject)

		assert.Equal(tc.expect.(json.JsonArray)[1].(json.JsonObject)["name"], got["name"], fmt.Sprintf("Expected %v got %v", tc.expect, got))
		assert.Equal(tc.expect.(json.JsonArray)[1].(json.JsonObject)["arrKey"].(json.JsonArray)[0], got["arrKey"].(json.JsonArray)[0], fmt.Sprintf("Expected %v got %v", tc.expect, got))
		assert.Equal(tc.expect.(json.JsonArray)[1].(json.JsonObject)["objKey"], got["objKey"], fmt.Sprintf("Expected %v got %v", tc.expect, got))
	})
}
