package getl

import (
	"fmt"
	"io"
)

// SinkNode is the wrapper for the user defined node definition to be run on the pipeline
// Name - Name of node
// Runner - pointer to the user defined function to dump the processed data
type SinkNode struct {
	Name   string
	Runner func(*io.PipeWriter, *io.PipeReader)
}

// Sink receives the user defined node definition
func (p *NodeWrapper) Sink(s *SinkNode) *NodeWrapper {

	nw := new(NodeWrapper)

	nw.Node = s
	nw.Name = s.Name

	r, w := io.Pipe()
	nw.Output = r

	s.Execute(s.Runner, w, p.Output, true)
	return nw
}

// Execute runs the user defined function
func (n *SinkNode) Execute(f func(*io.PipeWriter, *io.PipeReader), pipeWriter *io.PipeWriter, input *io.PipeReader, closeWhenDone bool) {
	fmt.Println("Dump sink: " + n.Name)
	go func() {
		if closeWhenDone {
			defer pipeWriter.Close()
		}
		f(pipeWriter, input)
	}()
}
