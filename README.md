# kmer-terminal-extract-nodes

- go program making specific kmers and then extracting those with the specific terminal edges.
- for the hapmers and also to sort out and filter the kmers with the specific terminal edges and bubbles. 
- it will first create the kmers and then will give out those kmers for which the terminal ends are mapping to the specific for implementation into the graph edges.
- instead of storing the kmers into a map, using those using the concurrency primitves to get the faster edges for the graphs. 
- you can chain this using the concurrency for extracting the complete graph edges and creating a map like
```
func gochain (input <- chan string, kmer <-chan int, terminaltypes <- chan string) {
 / put the function here.
}
```
-  Finished GO programming in a week.

Gaurav Sablok 
