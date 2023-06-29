package radix

func ExampleTrie() *Node {
	// []: node
	// [x]: leaf
	// trie: [r] te   [] st   [x] er   [x]
	//                   t    [x] as   [x]
	//                   rre  []  no   [x]
	//                            moto [x]
	// Root
	tr := NewNode(nil, true)
	// root -> first
	firstTrNode := NewNode(nil, false)
	// first -> second
	secondTrNode := NewNode(nil, false)
	// first -> third
	thirdTrNode := NewNode(nil, true)
	// first -> fourth
	fourthTrNode := NewNode(nil, true)
	// second -> fifth
	fifthTrNode := NewNode(nil, true)
	// second -> six
	sixTrNode := NewNode(nil, true)
	// third -> seven
	sevenTrNode := NewNode(nil, true)
	// fourth -> eight
	eightTrNode := NewNode(nil, true)

	tr.Edges = []*Edge{
		{TargetNode: firstTrNode, Label: "te"},
	}

	firstTrNode.Edges = []*Edge{
		{TargetNode: secondTrNode, Label: "rre"},
		{TargetNode: thirdTrNode, Label: "st"},
		{TargetNode: fourthTrNode, Label: "t"},
	}

	secondTrNode.Edges = []*Edge{
		{TargetNode: fifthTrNode, Label: "no"},
		{TargetNode: sixTrNode, Label: "moto"},
	}

	thirdTrNode.Edges = []*Edge{
		{TargetNode: sevenTrNode, Label: "er"},
	}

	fourthTrNode.Edges = []*Edge{
		{TargetNode: eightTrNode, Label: "as"},
	}

	return tr
}
