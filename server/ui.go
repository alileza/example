package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func uiFileServerMux(staticFilePath string) http.Handler {
	return http.FileServer(http.Dir(staticFilePath))
}

func uiProxyMux(uiProxyURL string) http.Handler {
	origin, err := url.Parse(uiProxyURL)
	if err != nil {
		panic("Failed to parse UIProxyURL: " + err.Error())
	}

	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", origin.Host)
			req.URL.Scheme = "http"
			req.URL.Host = origin.Host
		},
	}
}
