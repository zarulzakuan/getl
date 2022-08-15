package getl

import (
	"fmt"
	"io"
)

// SourceNode is the wrapper for the user defined node definition to be run on the pipeline
// Name - Name of node
// Runner - pointer to the user defined function to extract the data
type SourceNode struct {
	Name   string
	Runner func(*io.PipeWriter, *io.PipeReader)
}

// Source receives the user defined node definition
func (c *Scheduler) Source(s *SourceNode) *NodeWrapper {

	nw := new(NodeWrapper)

	nw.Node = s
	nw.Name = s.Name

	r, w := io.Pipe()

	nw.Output = r

	if c != nil {
		c.SchedAt.Do(func() { s.Execute(s.Runner, w, nil, false) })
		c.SchedAt.SingletonMode()
		c.SchedAt.StartAsync()
	} else {
		s.Execute(s.Runner, w, nil, true)
	}

	return nw
}

// Execute runs the user defined function
func (n *SourceNode) Execute(f func(*io.PipeWriter, *io.PipeReader), pipeWriter *io.PipeWriter, input *io.PipeReader, closeWhenDone bool) {
	fmt.Println("Get source: " + n.Name)
	go func() {
		if closeWhenDone {
			defer pipeWriter.Close()
		}
		f(pipeWriter, input)

	}()
}
