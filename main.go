package main

import (
	"fmt"
	"leetcode/offer"
)

func main() {
	ins := offer.Constructor()
	ins.Push(2)
	ins.Push(0)
	ins.Push(3)
	ins.Push(0)
	fmt.Println(ins.Min())
	ins.Pop()
	fmt.Println(ins.Min())
	ins.Pop()
	fmt.Println(ins.Min())
	ins.Pop()
	fmt.Println(ins.Min())
}
