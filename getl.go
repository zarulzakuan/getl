package getl

import (
	"io"
)

type NodeWrapper struct {
	Node   Node
	Name   string
	Output *io.PipeReader
}

type Node interface {
	Execute(func(*io.PipeWriter, *io.PipeReader), *io.PipeWriter, *io.PipeReader)
}
