package patricia

type Node struct {
	rlink *Node
	llink *Node
	key   string
	skip  int
}

func NewNode(key string, llink, rlink *Node, skip int) *Node {
	return &Node{
		key:   key,
		llink: llink,
		rlink: rlink,
		skip:  skip,
	}
}

func (n *Node) AddRLink(rlink *Node) {
	n.rlink = rlink
}

func (n *Node) AddLLink(llink *Node) {
	n.llink = llink
}

// Adapted from Algorithms in C
func (n *Node) search(key string, w int) string {
	if n.skip <= w {
		return n.key
	}

	if key[n.skip] == '0' {
		if n.llink == nil {
			return n.key
		}

		return n.llink.search(key, n.skip)
	}

	if n.rlink == nil {
		return n.key
	}

	return n.rlink.search(key, n.skip)
}

func (n *Node) Search(key string) bool {
	t := n.search(key, -1)

	return t == key
}

func (n *Node) Insert(key string) {
}
