package bridge

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	require.NoError(t, err)
}
