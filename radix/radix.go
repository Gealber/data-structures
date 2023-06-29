// Implementation of a radix tree following
// Wikipedia description
// https://en.wikipedia.org/wiki/Radix_tree
package radix

import (
	"fmt"
)

type Edge struct {
	TargetNode *Node
	Label      string
}

type Node struct {
	Edges  []*Edge
	IsLeaf bool
}

func NewNode(edges []*Edge, isLeaf bool) *Node {
	if edges == nil {
		edges = make([]*Edge, 0)
	}

	return &Node{Edges: edges, IsLeaf: isLeaf}
}

func NewEdge(label string, targetNode *Node) *Edge {
	if label == "" {
		panic("invalid label")

	}
	return &Edge{
		TargetNode: targetNode,
		Label:      label,
	}
}

// Taken from wikipedia pseudocode but slightly modified
func (n *Node) LookUp(key string) bool {
	traverseNode := n
	elementsFound := 0

	// traverse until a leaf is found or it is not possible to continue
	for elementsFound < len(key) {
		// get the next edge index
		neIndex := traverseNode.findNextEdgeIndex(key, elementsFound)
		// no edge was found
		if neIndex == -1 {
			return false
		}

		// increment elements found
		elementsFound += len(traverseNode.Edges[neIndex].Label)
		// setting next node to explore
		traverseNode = traverseNode.Edges[neIndex].TargetNode
	}

	return (traverseNode.IsLeaf && elementsFound == len(key))
}

// Insert key in trie
// traverse the trie until you cannot progress more
// at this point we add a new outgoing labeled edge with all
// remaining elements in the input string, or there is already
// an outgoing edge sharing a prefix with the remaining input string
// we split it into two edges
func (n *Node) Insert(key string) {
	traverseNode := n
	elementsFound := 0

	if key == "" {
		return
	}

	for elementsFound < len(key) {
		currentKey := key[elementsFound:]
		// check for common prefix
		idx, j := traverseNode.commonPrefixEdge(currentKey)
		if idx == -1 || j == 0 {
			leaf := NewNode(nil, true)
			edge := NewEdge(currentKey, leaf)
			traverseNode.Edges = append(traverseNode.Edges, edge)

			return
		}

		if j <= len(currentKey) && j != len(traverseNode.Edges[idx].Label)-1 {
			// split edge
			traverseNode = traverseNode.split(idx, j, currentKey)
			elementsFound += j
		}
	}
}

// Print method is WRONG
func (n *Node) Print(currentWord string) {
	if n == nil {
		return
	}

	traverseNode := n

	for _, edge := range traverseNode.Edges {
		currentWord += edge.Label
		if edge.TargetNode != nil && edge.TargetNode.IsLeaf {
			fmt.Println(currentWord)
		}

		if edge.TargetNode != nil {
			edge.TargetNode.Print(currentWord)
			currentWord = currentWord[:len(currentWord)-len(edge.Label)]
		}
	}
}

// Find the next edge index where the label
// of the edge is a prefix of key[elementsFound:]
func (n *Node) findNextEdgeIndex(key string, elementsFound int) int {
	index := -1

	for i, e := range n.Edges {
		if len(e.Label) > (len(key) - elementsFound) {
			continue
		}

		// check if label is a prefix of key[elementsFound:]
		if e.Label == key[elementsFound:][:len(e.Label)] {
			return i
		}
	}

	return index
}

// testy with tester
// test with idx 0 j 4
func (n *Node) commonPrefixEdge(key string) (int, int) {
	for i, e := range n.Edges {
		j := longestCommonIndex(key, e.Label)
		if j != -1 {
			return i, j
		}
	}

	return -1, -1
}

// key: test label: tester common: test
// key: tet label: tester  common: te
// key: aet label: tester  common:
func longestCommonIndex(key, label string) int {
	min := len(label)
	if len(key) < min {
		min = len(key)
	}

	j := 0
	for {
		if j == min {
			return j
		}

		if key[j] != label[j] {
			if j == 0 {
				return -1
			}

			return j
		}

		j++
	}
}

// split edge
func (n *Node) split(idx, commonPrefixEnd int, key string) *Node {
	if key == "" {
		return n
	}

	if commonPrefixEnd == len(n.Edges[idx].Label) {
		return n.Edges[idx].TargetNode
	}

	// create a new edge with the common prefix
	newNode := NewNode(nil, commonPrefixEnd == len(key))
	newEdge := NewEdge(key[:commonPrefixEnd], newNode)
	// set label of edge to what is not a common prefix
	n.Edges[idx].Label = n.Edges[idx].Label[commonPrefixEnd:]
	newNode.Edges = append(newNode.Edges, n.Edges[idx])

	// associate new edge with its position
	n.Edges[idx] = newEdge

	return newNode
}
