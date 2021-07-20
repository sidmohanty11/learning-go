package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFuncPrinter struct{}

func (m MultiFuncPrinter) Print(d Document) {

}
func (m MultiFuncPrinter) Fax(d Document) {

}
func (m MultiFuncPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct{} // cant scan or fax

// interface segregation principle...

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Fax interface {
	Fax(d Document)
}

type PhotoCopier interface {
	Print(d Document)
	Scan(d Document)
}

// the bigger picture
type MultiFuncDevice interface {
	// interfaces segregated ...
	Printer
	Scanner
	// or Fax
}
