package getl

import (
	"fmt"
	"io"
)

// TransformNode is the wrapper for the user defined node definition to be run on the pipeline
// Name - Name of node
// Runner - pointer to the user defined function to process the data

type TransformNode struct {
	Name   string
	Runner func(*io.PipeWriter, *io.PipeReader)
}

// Transform receives the user defined node definition

func (p *NodeWrapper) Transform(s *TransformNode) *NodeWrapper {

	nw := new(NodeWrapper)

	nw.Node = s
	nw.Name = s.Name

	r, w := io.Pipe()
	nw.Output = r

	s.Execute(s.Runner, w, p.Output, true)

	return nw
}

// Execute runs the user defined function

func (n *TransformNode) Execute(f func(*io.PipeWriter, *io.PipeReader), pipeWriter *io.PipeWriter, input *io.PipeReader, closeWhenDone bool) {
	fmt.Println("Transform data: " + n.Name)
	go func() {
		if closeWhenDone {
			pipeWriter.Close()
		}
		f(pipeWriter, input)
	}()
}
