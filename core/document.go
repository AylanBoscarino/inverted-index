package core

import (
    "strings"
)

type Token struct {
    value string
    ocurrencies string
    position int
}

type Document struct {
    tokens []Token
    content string
}

func NewToken(text string) Token {
    var t Token
    t.value = text
    return t
}

func NewDocument(text string) Document {
    var d Document
    d.content = text
    d.tokens = []Token{}
    fields := strings.Fields(text)
    for _, field := range fields {
        d.tokens = append(d.tokens, NewToken(strings.ToLower(field)))
    }
    return d
}
