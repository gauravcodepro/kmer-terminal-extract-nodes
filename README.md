# terminal-ends-graph-edges

- a go program for getting the terminal edges with the specific types.
-  A Go program for filtering of the kmers that are specific to the terminal ends. It will first create the kmers and then will give out those kmers for which the terminal ends are mapping to the specific for implementation into the graph edges.
-  Instead of storing the kmers into a map, I am directly using those using the concurrency primitves to get the faster edges for the graphs. 
-  You can chain this using the concurrency for extracting the complete graph edges and creating a map like
```
func gochain (input <- chan string, kmer <-chan int, terminaltypes <- chan string) {
 / put the function here.
}
```
-  Finished GO programming in a week.

Gaurav Sablok 
