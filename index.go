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
    m map[string][]*Document
}

func NewInvertedIndex() InvertedIndex {
    var i InvertedIndex
    i.m = make(map[string][]*Document)
    return i
    //return InvertedIndex{m:= make(map[string][]Document, d:= make([]Document))}
}

func (i *InvertedIndex) indexDocumentWithToken(token string, d *Document) {
    docs, isSet:= i.m[token]
    if isSet == false {
        i.m[token] = []*Document{}
    }
    i.m[token] = append(docs, d)
}

func (i *InvertedIndex) insertDocument(d Document) {
    for _, token := range d.tokens {
        i.indexDocumentWithToken(token.value, &d)
    }
}

func (i *InvertedIndex) findByToken(t string) {
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
    i.findByToken("abunda")
}
