package flamewriter

import (
	"bytes"
	"testing"
)

func makeRecord() *Record {
	root := NewRecord("root", 0)
	root.Add([]string{"one"}, 1)
	root.Add([]string{"one", "two"}, 2)
	root.Add([]string{"one", "two", "three"}, 3)
	return root
}

func TestJSONWriter(t *testing.T) {
	buf := bytes.NewBufferString("")
	writer := NewJSONWriter(buf)
	if err := writer.Write(makeRecord().ReduceRoot()); err != nil {
		t.Error(err)
		return
	}
}
