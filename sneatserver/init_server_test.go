package sneatserver

import (
	"testing"

	"github.com/strongo/delaying"
	"github.com/strongo/logus"
)

func TestInitServer(_ *testing.T) {
	logEntryHandler := logus.NewStandardGoLogger()
	InitServer([]string{"some-host.example.com"}, logEntryHandler, delaying.VoidWithLog)
}
