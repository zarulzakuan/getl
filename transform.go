package getl

import (
	"fmt"
	"io"
)

type TransformNode struct {
	Name   string
	Runner func(*io.PipeWriter, *io.PipeReader)
}

func (p *NodeWrapper) Transform(s *TransformNode) *NodeWrapper {

	nw := new(NodeWrapper)

	nw.Node = s
	nw.Name = s.Name

	r, w := io.Pipe()
	nw.Output = r

	s.Execute(s.Runner, w, p.Output)

	return nw
}

func (n *TransformNode) Execute(f func(*io.PipeWriter, *io.PipeReader), pipeWriter *io.PipeWriter, input *io.PipeReader) {
	fmt.Println("Transform data: " + n.Name)
	go f(pipeWriter, input)
}
