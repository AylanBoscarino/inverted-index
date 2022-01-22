package main

import "fmt"

func main() {
    i := NewInvertedIndex()
    d := NewDocument("a fenda que abunda força")
    i.insertDocument(d)
    fmt.Println(i)
    i.findByToken("abunda")
}
