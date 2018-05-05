package main

import (
	"fmt"
)

type itself []string

func (h *itself) appendToItself(test string) {
	*h = append(*h, test)
}

func main() {
	h := itself{"1", "2"}
	//h := []string{}
	//h := make([]string, 0, 10)
	//logs.Info(cap(h))
	//h.appendToItself("3")
	test(&h)
	fmt.Println(h, "<- how do I make it [1,2,3]")
}
func test(h *itself) {
	h.appendToItself("44")
	fmt.Println(h)
}
