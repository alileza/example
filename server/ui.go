package server

import (
	"net/http"
)

func uiMux(staticFilePath string) http.Handler {
	return http.FileServer(http.Dir(staticFilePath))
}
