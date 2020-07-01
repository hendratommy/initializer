package bar

import (
	"github.com/hendratommy/initializer"
	"github.com/hendratommy/initializer/examples"
	"time"
)

var barPrinter examples.Printer

func init() {
	initializer.Register("bar", Init)
}

func Init() error {
	barPrinter = examples.GetPrinter()
	return nil
}

func Bar() string {
	return barPrinter("hello from bar, current time is %s", time.Now())
}
