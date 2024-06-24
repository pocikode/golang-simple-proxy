package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/sirupsen/logrus"
)

type handle struct {
	remote *url.URL
	logger *logrus.Logger
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log incoming request
	h.logger.WithFields(logrus.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
	}).Info("Incoming request")

	// Proxy handler
	proxy := httputil.NewSingleHostReverseProxy(h.remote)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = h.remote.Scheme
		req.URL.Host = h.remote.Host
		req.Host = h.remote.Host
		req.URL.Path = r.URL.Path
		req.URL.RawQuery = r.URL.RawQuery
	}

	proxy.ServeHTTP(w, r)
}
