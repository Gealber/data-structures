package radix

import (
	"fmt"
	"testing"
)

func Test_findNextEdgeIndex(t *testing.T) {
	t.Run("find next edge index", func(t *testing.T) {
	})
}

type testCase struct {
	name        string
	key         string
	expectedIdx int
	expectedJ   int
}

func Test_commonPrefixEdge(t *testing.T) {

	tr := ExampleTrie()
	for _, tc := range tcs() {
		t.Run(tc.name, func(t *testing.T) {
			idx, j := tr.commonPrefixEdge(tc.key)
			if idx != tc.expectedIdx || j != tc.expectedJ {
				msg := fmt.Sprintf("invalid result expected idx: %d got idx: %d expected j: %d got j: %d", tc.expectedIdx, idx, tc.expectedJ, j)
				t.Fatal(msg)
			}
		})
	}
}

func tcs() []testCase {
	return []testCase{
		{name: "searching for t", key: "t", expectedIdx: 0, expectedJ: 1},
		{name: "searching for te", key: "te", expectedIdx: 0, expectedJ: 2},
		{name: "searching for testy", key: "testy", expectedIdx: 0, expectedJ: 2},
	}
}
