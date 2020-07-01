package main

import (
	"fmt"
	"github.com/hendratommy/initializer"
	"github.com/hendratommy/initializer/examples"
	"github.com/hendratommy/initializer/examples/bar"
)

func main() {
	prt := func(format string, v ...interface{}) string {
		return fmt.Sprintf(format, v)
	}
	examples.Setup(prt)

	initializer.Init()

	//fmt.Println(foo.Foo()) // this will cause nil pointer dereference error
	fmt.Println(bar.Bar())
}
