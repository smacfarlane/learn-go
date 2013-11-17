package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    var count = make(map[string]int)

    for _, field := range(strings.Fields(s)) {
        count[field] += 1
    }

    return count 
}

func main() {
    wc.Test(WordCount)
}
