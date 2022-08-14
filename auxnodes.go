package getl

import "io"

func (p *NodeWrapper) Tee(s *NodeWrapper) *NodeWrapper {

	// file -----> writer1 <===(read)==== reader1
	//                                        | copy
	//                                        v
	//                                     writer ------> writer2 <===(read)=== reader2
	//                                            |
	//                                            | multiply
	//                                            |
	//                                            ------> writer3 <===(read)=== reader3
	//

	oRead, oWrite := io.Pipe()
	sRead, sWrite := io.Pipe()

	nw := new(NodeWrapper)

	writer := io.MultiWriter(oWrite, sWrite)
	go io.Copy(writer, p.Output)

	s.Name = p.Name
	s.Node = p.Node
	s.Output = oRead

	nw.Name = p.Name
	nw.Node = p.Node
	nw.Output = sRead

	// forward to next node
	return nw
}

func (p *NodeWrapper) Union(s *NodeWrapper) *NodeWrapper {

	// source1 -----> writer1 <===(read)==== reader1--------
	//														|
	//														| merger
	//														|
	// source2 -----> writer2 <===(read)==== reader2-------------------reader
	//																	| copy
	//																	v
	//																	fWrite <===(read)== fRead

	nw := new(NodeWrapper)
	fRead, fWrite := io.Pipe()
	reader := io.MultiReader(p.Output, s.Output)

	go io.Copy(fWrite, reader)
	nw.Output = fRead

	return nw

}

func TeeAdapter() *NodeWrapper {

	return new(NodeWrapper)
}
