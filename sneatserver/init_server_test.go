package sneatserver

import (
	"github.com/strongo/delaying"
	"github.com/strongo/logus"
	"testing"
)

func TestInitServer(_ *testing.T) {
	logEntryHandler := logus.NewStandardGoLogger()
	InitServer([]string{"some-host.example.com"}, logEntryHandler, delaying.VoidWithLog)
}
