package examples

type Printer = func(string, ...interface{}) string

var printer Printer

func GetPrinter() Printer {
	return printer
}

func Setup(p Printer) {
	printer = p
}
