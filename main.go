package main

import (
	"io"
	"os"
	"strings"

	"github.com/sneat-co/sneat-go-core/emails/email2writer"
	"github.com/sneat-co/sneat-go-server/sneatgaeapp"
	"github.com/sneat-co/sneat-go-server/sneatserver"
	"github.com/strongo/delaying"
	"github.com/strongo/logus"
)

func main() { // TODO: document why we need this wrapper

	logEntryHandler := logus.NewStandardGoLogger()
	extraKnownHosts := strings.Split(os.Getenv("KNOWN_HOSTS"), ",")

	httpRouter := sneatserver.InitServer(extraKnownHosts, logEntryHandler, delaying.VoidWithLog)

	emailClient := email2writer.NewClient(func() (io.StringWriter, error) {
		return os.Stdout, nil
	})

	// TODO(refactoring): Should we combine sneatgaeapp.Start() with sneatserver.InitServer()?
	sneatgaeapp.Start(nil, nil, httpRouter, emailClient)
}
