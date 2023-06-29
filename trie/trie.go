// Following wikipedia description of a trie
// adapted to an alphabet of size 2. For binary representation.
// https://en.wikipedia.org/wiki/Trie
// This implementation is not efficient for this particular case
// would be better to implement a Patrician tree wich
// are Radix trees with radix equal 2.
package trie

// Node implementation of a trie with an alphabet of
// size 2, which correspond to 0 and 1
// we are not going to include IsTerminal neither Value
// IsTerminal is used for deletetion and we are not implementing it
// And value is not needed we just want to make sure we find the key
type Node struct {
    // each node can have a children for every letter of the alphabet
    // in our case we are using an alphabet of size 2
    // cause we want to look binary strings
	Children   [2]*Node
}

func New() *Node {
	return &Node{}
}

// Lookup operation of trie
// Taken from wikipedia
// for 0 ≤ i < key.length do
//
//	    if x.Children[key[i]] = nil then
//	        return false
//	    end if
//	    x := x.Children[key[i]]
//	repeat
//	return x.Value
//
// Example:
// suppose key = 101
// we will iterate 3 times
// Suppose our trie will have height 4
// Iteration 1 Example:
//
//	index = key[0] - '0' => '1' - '0' => 1
//	n.Children[1] =>  won't be nil, it will contain the sub tree in
//	the right side of the root node
//	n = n.Children[1] right sub tree
//
// Iteration 2:
//
//	n = n.Children[1].Children[0]
//
// Iteration 3:
//
//	n = n.Children[1].Children[0].Children[1]
//
// Given that we didn't found a nil node we can assure
// that the value exists in the trie
func (n *Node) Lookup(key string) bool {
	for i := 0; i < len(key); i++ {
		index := key[i] - '0'
		if n.Children[index] == nil {
            // if we cannot continue exploring the trie
            // then this key is not on it
			return false
		}

		// in next iteration we will look from
		// where we left so keep constructing the trie according to the key
		n = n.Children[index]
	}

    // given that we could explore all the chars
    // in the key, that means the key can be represented as a path in the trie
	return true
}

// Insert operation of trie
// Pseudocode taken from Wikipedia
// Trie-Insert(x, key, value)
//
//	for 0 ≤ i < key.length do
//	    if x.Children[key[i]] = nil then
//	        x.Children[key[i]] := Node()
//	    end if
//	    x := x.Children[key[i]]
//	repeat
//	x.Value := value
//	x.Is-Terminal := True
func (n *Node) Insert(key string) {
	for i := 0; i < len(key); i++ {
        // remember each node could have a children for each letter
        // of the alphabet. In this case our alphabet is composed
        // by '0' and '1'. In the following operation I'm substracting
        // the run in key[i] to the ascii value of '0' which is 48
        // for example if we have key[i] = '1' then this is
        // '1' - '0' => 49 - 48. Keep in mind that '1' in Golang
        // is not a string but a rune, which is a 32-bit integer
		index := key[i] - '0'
		if n.Children[index] == nil {
			// create a new node in this index
			n.Children[index] = New()
		}
		// in next iteration we will look from
		// where we left so keep constructing the trie according to the key
		n = n.Children[index]
	}
}
