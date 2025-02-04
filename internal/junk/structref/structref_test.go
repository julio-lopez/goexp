package structref

// objective: try out different struct and interface representations. This is
// to get insight into their space cost (memory size), as well as how much of it
// can be implemented without resorting to `unsafe.Pointer`.
//
// Additional insight will require implementing other basic `trie` structures
// first, and later implementing "packed array" representations.

import (
	"testing"
	"unsafe"
)

const (
	node4Max  = 4
	node16Max = 16
)

type nodeRef interface {
	getChildren() []nodeRef
	getChildrenCapacity() int
	getChildCount() int
	isTerminal() bool
}

type nodeBase struct {
	childCount uint8
}

type keyRef struct {
	r nodeRef
	k byte
	// Note: type and key could be packed into the same struct if the pointer
	// could be stored separately, either via unsafe.Pointer or using an
	// explicit "unsafe cast", is this possible ?
	// in that case the struct would be 8 + 1 + 1 = 10 -> 16 with 8-byte aligment
	// instead of 24 as in this struct.
	// Now if the key is stored separately in a byte array / slice then the
	// "amortized storage" is 16 + 1 byte = 17 bytes.
}

type (
	refArray4  [node4Max]nodeRef
	refArray16 [node16Max]nodeRef
)

type node4 struct {
	refs refArray4
	keys [node4Max]byte
	nodeBase
	// TODO: the set of keys present needs to be explicitly included in the node
	// it cannot be a bitmap for nodes smaller than node16, assuming a `byte`
	// alphabet, that is, 256 symbols in the alphabet.
}

func (n nodeBase) isTerminal() bool   { return false }
func (n nodeBase) getChildCount() int { return int(n.childCount) }

func (n *node4) getChildren() []nodeRef {
	return n.refs[:]
}

func (n *node4) getChildrenCapacity() int {
	return node4Max
}

type node16 struct {
	refs refArray16
	keys [node16Max]byte
	nodeBase
	// TODO: the set of keys present needs to be explicitly included in the node
	// it cannot be a bitmap for nodes smaller than node16, assuming a `byte`
	// alphabet, that is, 256 symbols in the alphabet.
}

func (n *node16) getChildren() []nodeRef {
	return n.refs[:]
}

func (n *node16) getChildrenCapacity() int {
	return node16Max
}

func TestTypes(t *testing.T) {
	var (
		ra4 refArray4
		n4  node4
		n16 node16
		nr  nodeRef
		nb  nodeBase
		kr  keyRef
	)

	t.Log("refArray4 len:", len(ra4))
	t.Log("refArray4 cap:", cap(ra4))
	t.Log("n4.getChildrenCapacity:", n4.getChildrenCapacity())

	t.Log("sizeof(nodeRef):", unsafe.Sizeof(nr))
	t.Log("sizeof(r4):", unsafe.Sizeof(ra4))
	t.Log("sizeof(nodeBase):", unsafe.Sizeof(nb))
	t.Log("sizeof(node4):", unsafe.Sizeof(n4))
	t.Log("sizeof(node16):", unsafe.Sizeof(n16))
	t.Log("sizeof(keyRef):", unsafe.Sizeof(kr))

	// check interface compatilibity
	nr = &n4
	nr = &n16
}
