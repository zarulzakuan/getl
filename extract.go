package getl

import (
	"fmt"
	"io"
)

type SourceNode struct {
	Name   string
	Runner func(*io.PipeWriter, *io.PipeReader)
}

func (schedule *Scheduler) Source(s *SourceNode) *NodeWrapper {

	nw := new(NodeWrapper)

	nw.Node = s
	nw.Name = s.Name

	r, w := io.Pipe()

	nw.Output = r

	schedule.SchedAt.Do(func() { s.Execute(s.Runner, w, nil) })
	schedule.SchedAt.StartAsync()

	return nw
}

func (n *SourceNode) Execute(f func(*io.PipeWriter, *io.PipeReader), pipeWriter *io.PipeWriter, input *io.PipeReader) {
	fmt.Println("Get source: " + n.Name)
	go f(pipeWriter, input)
}
