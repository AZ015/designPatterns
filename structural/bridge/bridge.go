package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct {
}

func (i *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)

	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass ioi.Writer")
	}
	fmt.Printf("%s\n", msg)

	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	err := n.Printer.PrintMessage(n.Msg)
	if err != nil {
		return errors.New(fmt.Sprintf("error in normal printer: %s", err))
	}

	return nil
}

type PactPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PactPrinter) Print() error {
	err := p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
	if err != nil {
		return errors.New(fmt.Sprintf("error in pact printer: %s", err))
	}

	return nil
}
