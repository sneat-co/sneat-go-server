package sneatserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sneat-co/sneat-go-backend/src/sneatgae/sneatgaeapp"
	"github.com/sneat-co/sneat-go-core/security"
	"github.com/sneat-co/sneat-go-server/firebase4sneat"
	"github.com/strongo/delaying"
	"github.com/strongo/logus"
	"net/http"
)

func InitServer(
	extraKnownHosts []string,
	logEntryHandler logus.LogEntryHandler,
	registerDelayedFunc delaying.RegisterDelayedFunc,
) *httprouter.Router {
	if logEntryHandler == nil {
		panic("logEntryHandler is required")
	}
	logus.AddLogEntryHandler(logEntryHandler)

	if len(extraKnownHosts) > 0 {
		security.AddKnownHosts(extraKnownHosts...)
	}

	firebase4sneat.InitFirebase()

	httpRouter := sneatgaeapp.CreateHttpRouter()

	ServeStaticFiles(httpRouter)

	delaying.Init(registerDelayedFunc)

	return httpRouter
}

func ServeStaticFiles(httpRouter *httprouter.Router) {

	staticDir := http.Dir("static")
	httpRouter.ServeFiles("/static/*filepath", staticDir)

	fileServer := http.FileServer(staticDir)

	for _, file := range []struct {
		path        string
		contentType string
	}{
		{path: "robots.txt", contentType: "text/plain"},
		{path: "no-robots.txt", contentType: "text/plain"},
		{path: "favicon.ico", contentType: "image/x-icon"},
	} {
		httpRouter.GET("/"+file.path, func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			w.Header().Set("Content-Type", file.contentType)
			w.Header().Set("Cache-Control", "public, max-age=3600") // 1 hour (3600 seconds)
			fileServer.ServeHTTP(w, req)
		})
	}
}
