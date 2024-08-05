// Author Gaurav Sablok 
// Universitat Potsdam 
// Date 2024-8-5
/*
 A Go program for filtering of the kmers that are specific to the terminal ends. It will first create the kmers and then will give 
 out those kmers for which the terminal ends are mapping to the specific for implementation into the graph edges. You can change the value of 2 to any kmer specific. 
 I am building a go package as a part of go learning which will have all these including a desktop application, written in GO.
*/

package main 
import (
    "fmt"
    "stringconv"
)
func typeA () {
    msg := "ATATATAGATAGCGC"
    for i := 0; i <= len(msg)-2; i++ {
           if string(msg[i:i+2][len(msg[i:i+2])-1:]) == "A"{
           fmt.Println("The kmers with the typeA are:")
           fmt.Println(msg[i:i+2])
            
        }
    }
}
func typeG () {
    msg := "ATATATAGATAGCGC"
    for i := 0; i <= len(msg)-2; i++ {
           if string(msg[i:i+2][len(msg[i:i+2])-1:]) == "T"{
           fmt.Println("The kmers with the typeT are:")
           fmt.Println(msg[i:i+2])
            
        }
    }
}
func typeC () {
    msg := "ATATATAGATAGCGC"
    for i := 0; i <= len(msg)-2; i++ {
           if string(msg[i:i+2][len(msg[i:i+2])-1:]) == "G"{
           fmt.Println("The kmers with the typeG are:")
           fmt.Println(msg[i:i+2])
            
        }
    }
}
func typeA () {
    msg := "ATATATAGATAGCGC"
    for i := 0; i <= len(msg)-2; i++ {
           if string(msg[i:i+2][len(msg[i:i+2])-1:]) == "C"{
           fmt.Println("The kmers with the typeC are:")
           fmt.Println(msg[i:i+2])
            
        }
    }
}        
