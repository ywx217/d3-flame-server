package flamewriter

// Record flame record with call stack info and statistical values
type Record struct {
	name     string
	value    int
	children map[string]*Record
}

// NewRecord creates a new Record
func NewRecord(name string, value int) *Record {
	return &Record{
		name:     name,
		value:    value,
		children: make(map[string]*Record, 0),
	}
}

// Add adds a callstack
func (r *Record) Add(stack []string, value int) {
	r.value += value
	if len(stack) == 0 {
		return
	}
	if child, ok := r.children[stack[0]]; ok {
		child.Add(stack[1:], value)
	} else {
		child = NewRecord(stack[0], 0)
		r.children[stack[0]] = child
		child.Add(stack[1:], value)
	}
}

// ReduceRoot returns the only child if len(children) == 1, otherwise the root itself
func (r *Record) ReduceRoot() *Record {
	if len(r.children) == 1 {
		for _, child := range r.children {
			return child
		}
	}
	return r
}
