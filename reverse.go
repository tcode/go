package main

import (
        "fmt"
        "flag"
        "os"
        "./PrintHello"
)

func main () {
PrintHello.PrintHello()
flag.Parse()
i := flag.Arg(0)
if i == "x" {
        fmt.Printf("Hmm Why did you give me that x \u1011?\n")
        os.Exit(0)
} else if i == "c" { 
        fmt.Printf("Now you have given me the value c \uCCCC \n")
        var i string
        fmt.Scanf("%s", &i)
        if i == "s" {
                fmt.Printf("You have now chosen to start a little part of my program\n")
                fmt.Printf("\n\n")
                fmt.Printf("Now we can start a little adventure\n")
}
        os.Exit(0)
} else {
os.Exit(0) }
}

