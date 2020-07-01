# Initializer

Deferred `init` function, call `init` when you ready, no hurry :)

```go
package bar

var barPrinter examples.Printer

func init() {
	initializer.Register("bar", Init)
}

func Init() error {
	barPrinter = examples.GetPrinter()
	return nil
}
```

```go
package main

examples.Setup(prt) // setup all dependencies

initializer.Init() // call Init to start initializing your packages

fmt.Println(bar.Bar())
```

See [example](https://github.com/hendratommy/initializer/tree/master/examples)