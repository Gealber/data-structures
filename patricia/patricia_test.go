package patricia_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Gealber/data-structures/patricia"
)

type testCase struct {
	name           string
	key            int64
	expectedResult bool
}

func Test_Search(t *testing.T) {
	tr := exampleTree()
	for _, tc := range basicTreeTcs() {
		t.Run(tc.name, func(t *testing.T) {
			if tr.STearch(chrToBinStr(tc.key)) != tc.expectedResult {
				t.Fatal(fmt.Sprintf("unexpected result for key %d, expected: %t", tc.key, tc.expectedResult))
			}
		})
	}
}

func chrToBinStr(chr int64) string {
	return strconv.FormatInt(chr-64, 2)
}

func exampleTree() *patricia.Node {
	// adding node with key A
	anode := patricia.NewNode(chrToBinStr('A'), nil, nil, 4)
	// adding node with key C
	cnode := patricia.NewNode(chrToBinStr('C'), anode, nil, 3)
	// adding node with key E
	enode := patricia.NewNode(chrToBinStr('E'), cnode, nil, 2)
	// adding node with key H
	hnode := patricia.NewNode(chrToBinStr('H'), enode, nil, 1)
	// adding node with key R
	rnode := patricia.NewNode(chrToBinStr('R'), nil, nil, 4)
	// adding node with key S, which is the root node
	snode := patricia.NewNode(chrToBinStr('S'), hnode, rnode, 0)
	// adding right node to rnode
	rnode.AddRLink(snode)

	return snode
}

func basicTreeTcs() []testCase {
	return []testCase{
		// true cases
		{name: "searching for key R", key: 'R', expectedResult: true},
		{name: "searching for key S", key: 'S', expectedResult: true},
		{name: "searching for key H", key: 'H', expectedResult: true},
		{name: "searching for key E", key: 'E', expectedResult: true},
		{name: "searching for key C", key: 'C', expectedResult: true},
		{name: "searching for key A", key: 'A', expectedResult: true},
		// false cases
		{name: "searching for key T", key: 'T', expectedResult: false},
		{name: "searching for key M", key: 'M', expectedResult: false},
		{name: "searching for key N", key: 'N', expectedResult: false},
		{name: "searching for key O", key: 'O', expectedResult: false},
		{name: "searching for key P", key: 'P', expectedResult: false},
		{name: "searching for key Q", key: 'Q', expectedResult: false},
		{name: "searching for key V", key: 'V', expectedResult: false},
		{name: "searching for key W", key: 'W', expectedResult: false},
		{name: "searching for key X", key: 'X', expectedResult: false},
	}
}
