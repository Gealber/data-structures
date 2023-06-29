package trie_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Gealber/data-structures/trie"
)

type testCase struct {
	name      string
	keyInsert string
	keyLookup string
	result    bool
}

func Test_InsertLookup(t *testing.T) {
	tcs := genTcs()

	for _, tc := range tcs {
		tr := trie.New()

		t.Run(tc.name, func(t *testing.T) {
			// insert keys on it
			tr.Insert(tc.keyInsert)
			// search for them
			if tr.Lookup(tc.keyLookup) != tc.result {
				t.Fatal("unable to find key just inserted")
			}
		})
	}
}

func genTcs() []testCase {
	tcs := make([]testCase, 0)

	simpleTc := testCase{
		name:      "simple insert and lookup, true",
		keyInsert: "01",
		keyLookup: "01",
		result:    true,
	}
	tcs = append(tcs, simpleTc)

	falseTc := testCase{
		name:      "simple false tc",
		keyInsert: "01",
		keyLookup: "011",
		result:    false,
	}
	tcs = append(tcs, falseTc)

	// generating 100 random tcs
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		n := rand.Int()
		binStrInsert := strconv.FormatInt(int64(n), 2)
		binStrLookup := binStrInsert

		if n > 1 && i%2 == 0 {
			// creating false cases
			binStrLookup = binStrInsert[1:]
		}

		tcName := fmt.Sprintf("random %t tc for insert key: %s lookup key: %s", binStrInsert == binStrLookup, binStrInsert, binStrLookup)

		tcs = append(tcs, testCase{
			name:      tcName,
			keyInsert: binStrInsert,
			keyLookup: binStrLookup,
			result:    binStrInsert == binStrLookup,
		})
	}

	return tcs
}
