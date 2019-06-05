package flamewriter

import "io"

// HTMLWriter writes the flame in html format, with d3-flame-graph support
type HTMLWriter struct {
	w io.Writer
}

// NewHTMLWriter creates a new HTMLWriter
func NewHTMLWriter(w io.Writer) *HTMLWriter {
	return &HTMLWriter{w: w}
}

const (
	tpHTMLPrefix = `<head>
  <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/gh/spiermar/d3-flame-graph@2.0.3/dist/d3-flamegraph.css">
</head>
<body>
  <div id="chart"></div>
  <script type="text/javascript" src="https://d3js.org/d3.v4.min.js"></script>
  <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/d3-tip/0.9.1/d3-tip.min.js"></script>
  <script type="text/javascript" src="https://cdn.jsdelivr.net/gh/spiermar/d3-flame-graph@2.0.3/dist/d3-flamegraph.min.js"></script>
  <script type="text/javascript">
  var flamegraph = d3.flamegraph()
    .width(960);

  {
    var data = `
	tpHTMLSuffix = `;
    d3.select("#chart")
      .datum(data)
      .call(flamegraph);
  };
  </script>
</body>`
)

// Write writes the flame into html
func (w *HTMLWriter) Write(root *Record) error {
	if _, err := w.w.Write([]byte(tpHTMLPrefix)); err != nil {
		return err
	}
	jsonWriter := NewJSONWriter(w.w)
	if err := jsonWriter.Write(root); err != nil {
		return err
	}
	if _, err := w.w.Write([]byte(tpHTMLSuffix)); err != nil {
		return err
	}
	return nil
}
