package core

import (
    "fmt"
)

type InvertedIndex struct {
    m map[string][]*Document
}

func NewInvertedIndex() InvertedIndex {
    var i InvertedIndex
    i.m = make(map[string][]*Document)
    return i
}

func (i *InvertedIndex) indexDocumentWithToken(token string, d *Document) {
    docs, isSet:= i.m[token]
    if isSet == false {
        i.m[token] = []*Document{}
    }
    i.m[token] = append(docs, d)
}

func (i *InvertedIndex) InsertDocument(d Document) {
    for _, token := range d.tokens {
        i.indexDocumentWithToken(token.value, &d)
    }
}

func (i *InvertedIndex) FindByToken(t string) {
    docs, _ := i.m[t]
    if docs != nil {
        fmt.Println("Found document")
        for _, doc := range docs {
            fmt.Println(doc.content)
        }
    } else {
        fmt.Println("Document not found")
    }
}

