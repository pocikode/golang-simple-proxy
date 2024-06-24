package main

import (
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

func main() {
	cmd := ParseCmd()

	// Setting up the logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	url, err := url.Parse(cmd.remote)
	if err != nil {
		logger.Fatalf("Failed to parse target URL: %v", err)
	}

	// Start server
	logger.Infof("Starting proxy server on %s", cmd.bind)
	h := &handle{remote: url, logger: logger}
	http.HandleFunc("/", h.ServeHTTP)

	if err := http.ListenAndServe(cmd.bind, nil); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
