package treesort

type Tree struct {
	Value       int
	Left, Right *Tree
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elementes of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.Left)
		values = append(values, t.Value)
		values = appendValues(values, t.Right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &Tree{Value: value}.
		t = new(Tree)
		t.Value = value
		return t
	}
	if value < t.Value {
		t.Left = add(t.Left, value)
	} else {
		t.Right = add(t.Right, value)
	}
	return t
}
