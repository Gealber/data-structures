package radix_test

import (
	"fmt"
	"testing"

	"ton/practice/datastructures/radix"
)

type testCase struct {
	name      string
	keyInsert string
	keyLookup string
	result    bool
}

func Test_InsertLookup(t *testing.T) {

	tr := radix.NewNode(nil, true)
	for _, tc := range tcs() {
		t.Run(tc.name, func(t *testing.T) {
			tr.Insert(tc.keyInsert)
			if tr.LookUp(tc.keyLookup) != tc.result {
				t.Fatal(fmt.Sprintf("for key inserted: '%s' and key searched: '%s' was expected result: %t", tc.keyInsert, tc.keyLookup, tc.result))
			}
		})
	}
}

func Test_Lookup(t *testing.T) {
	t.Run("look up", func(t *testing.T) {
		tr := radix.ExampleTrie()
		key := "tetas"

		if !tr.LookUp(key) {
			t.Fatal(fmt.Sprintf("not found key searched: '%s'", key))
		}
	})
}

func Test_Print(t *testing.T) {
	t.Run("printing radix trie", func(t *testing.T) {
		tr := radix.ExampleTrie()
		tr.Print("")
	})
}

func tcs() []testCase {
	return []testCase{
		{name: "insert lookup", keyInsert: "tester", keyLookup: "tester", result: true},
		{name: "insert lookup", keyInsert: "test", keyLookup: "test", result: true},
		{name: "insert lookup", keyInsert: "tet", keyLookup: "tet", result: true},
		{name: "insert lookup", keyInsert: "tetas", keyLookup: "tetas", result: true},
		{name: "insert lookup", keyInsert: "terreno", keyLookup: "terreno", result: true},
		{name: "insert lookup", keyInsert: "terremoto", keyLookup: "terremoto", result: true},
	}
}
