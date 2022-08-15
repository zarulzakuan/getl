# getl

getl, pronounce "getle" as in Go ETL is a framework for building pipeline for data integration
and data transformation. Similar to water pipeline, data is streamed from source to sink, and
transformed in between can be split into multiple flows and merged from multiple flows.
Nodes are the data processors. They can be responsible for data extracting; Source, data dumping; Sink
or data transformation; Transform. And 2 auxiliary nodes, Tee and Union

```go
 ___________                 _____________                 __________
|           |               |             |               |         |
|  Source   | ====(pipe)====|  Transform  | ====(pipe)====|  Sink   |
| (runner)  |               |  (runner)   |               | (runner)|
|___________|               |_____________|               |_________|
```

Data pipes are built automatically and will break if any of the pipe inlets is closed.
Each node (except the aux nodes) must have a runner, a user defined function to extract, process or write down the data.
These runners must satisfy the interface requirement which they must have a data reader and writer as parameters.

Example:

```go
func writeTerminal(writer *io.PipeWriter, input *io.PipeReader) {
   for {
       buff := make([]byte, 50)
       n, err := input.Read(buff)
       if n != 0 {
           println(string(buff[:n]))
       }
       if err != nil {
           break
       }
    }
}

sink := new(getl.SinkNode)
sink.Name = "Write to terminal"
sink.Runner = writeTerminal
```

Once we have defined all the nodes, we can use chain them together like the following examples:

Example 1:

```go
getl.RunNow("0 *\/1 * * *", time.Local, false).Source(source).Transform(filter).Sink(sink) // run every 1 hour and start immediately
```

Example 2 - Split into multiple data flows:

```go
ta := getl.TeeAdapter()
getl.RunAt(300, time.Local, true).Source(source1).Tee(ta).Transform(filter).Sink(sink) // run every 5 minutes (300s)
ta.Transform(sort).Sink(sink2)
```

Example 3 - Join multiple data flows:

```go
dataflow1 := getl.RunNow().Source(source1).Tee(ta).Transform(filter).Sink(sink) // run every 30 minutes
getl.RunNow().Source(source2).Union(dataflow1).Sink(sink)
```

---

