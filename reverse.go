package main

import (
        "fmt"
        "flag"
        "os"
)

func main () {
fmt.Printf("hello world, \uACDF \n")
flag.Parse()
i := flag.Arg(0)
if i != "x" {
        os.Exit(0)        
 
} else {
fmt.Printf("Hmm Why did you give me that x ?\n") }
}

