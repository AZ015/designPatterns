package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct {
}

func (m *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s", s)
	println(newMsg)

	return newMsg
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		newMsg := fmt.Sprintf("Adapter: %s", p.Msg)

		return p.OldPrinter.Print(newMsg)
	}

	return p.Msg
}
