package adapter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}

	returnedMsg := adapter.PrintStored()
	require.Equal(t, returnedMsg, "Legacy Printer: Adapter: Hello World!")

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	require.Equal(t, returnedMsg, "Hello World!")
}
