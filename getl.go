package getl

import (
	"io"
)

// NodeWrapper is used to pass nodes around, making the chain of execution possible

type NodeWrapper struct {
	Node   Node
	Name   string
	Output *io.PipeReader
}

// Node definition

type Node interface {
	Execute(func(*io.PipeWriter, *io.PipeReader), *io.PipeWriter, *io.PipeReader, bool)
}
