package flamewriter

// Record flame record with call stack info and statistical values
type Record struct {
	Name     string
	Value    int
	Children map[string]*Record
}

// NewRecord creates a new Record
func NewRecord(name string, value int) *Record {
	return &Record{
		Name:     name,
		Value:    value,
		Children: make(map[string]*Record, 0),
	}
}

// Add adds a callstack
func (r *Record) Add(stack []string, value int) {
	r.Value += value
	if len(stack) == 0 {
		return
	}
	if child, ok := r.Children[stack[0]]; ok {
		child.Add(stack[1:], value)
	} else {
		child = NewRecord(stack[0], 0)
		r.AddChild(child)
		child.Add(stack[1:], value)
	}
}

// AddChild adds a child record
func (r *Record) AddChild(child *Record) {
	r.Children[child.Name] = child
}

// ReduceRoot returns the only child if len(children) == 1, otherwise the root itself
func (r *Record) ReduceRoot() *Record {
	if len(r.Children) == 1 {
		for _, child := range r.Children {
			return child
		}
	}
	return r
}

// FixRootValue fix value using sum of children values
func (r *Record) FixRootValue() *Record {
	r.Value = 0
	for _, child := range r.Children {
		r.Value += child.Value
	}
	return r
}
