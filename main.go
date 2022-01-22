package main

import (
    "fmt"
    "github.com/AylanBoscarino/inverted-index/core"
)

func main() {
    i := core.NewInvertedIndex()

    d := core.NewDocument("a fenda que abunda força")
    i.InsertDocument(d)
    fmt.Println(i)
    i.FindByToken("abunda")
}
