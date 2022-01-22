package main

import "fmt"

type Token struct {
    value string
    ocurrencies string
    position int
}

type Document struct {
    tokens []Token
    content string
}

type InvertedIndex struct {
    m map[string][]Document
}

func NewInvertedIndex() InvertedIndex {
    return InvertedIndex{make(map[string][]Document)}
}

func (i *InvertedIndex) indexDocumentWithToken(token string, d Document) {
    docs, isSet:= i.m[token]
    if isSet == false {
        i.m[token] = []Document{}
    }
    i.m[token] = append(docs, d)
}

func (i *InvertedIndex) insertDocument(d Document) {
    for _, token := range d.tokens {
        i.indexDocumentWithToken(token.value, d)
    }
}

func main() {
    i := NewInvertedIndex()
    d := Document{
        content: "a fenda que abunda força",
        tokens: []Token{
            Token{value: "a"},
            Token{value: "fenda"},
            Token{value: "que"},
            Token{value: "abunda"},
            Token{value: "força"},
        },
    }
    i.insertDocument(d)
    fmt.Println(i.m)
}
