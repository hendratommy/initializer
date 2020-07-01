package foo

import (
	"github.com/hendratommy/initializer/examples"
	"time"
)

var fooPrinter examples.Printer

func init() {
	fooPrinter = examples.GetPrinter()
}

func Foo() string {
	return fooPrinter("hello from foo, current time is %s", time.Now())
}
