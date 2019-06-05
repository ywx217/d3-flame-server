package flamewriter

import (
	"bytes"
	"testing"
)

func TestHTMLWriter(t *testing.T) {
	buf := bytes.NewBufferString("")

	writer := NewHTMLWriter(buf)
	if err := writer.Write(makeRecord().ReduceRoot()); err != nil {
		t.Error(err)
		return
	}
}
