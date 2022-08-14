package getl

import (
	"fmt"
	"io"
)

type SinkNode struct {
	Name   string
	Runner func(*io.PipeWriter, *io.PipeReader)
}

func (p *NodeWrapper) Sink(s *SinkNode) *NodeWrapper {

	nw := new(NodeWrapper)

	nw.Node = s
	nw.Name = s.Name

	r, w := io.Pipe()
	nw.Output = r

	s.Execute(s.Runner, w, p.Output)
	return nw
}

func (n *SinkNode) Execute(f func(*io.PipeWriter, *io.PipeReader), pipeWriter *io.PipeWriter, input *io.PipeReader) {
	fmt.Println("Dump sink: " + n.Name)
	go f(pipeWriter, input)
}
