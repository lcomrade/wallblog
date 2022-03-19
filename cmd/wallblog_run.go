// Copyright (C) 2022 Leonid Maslakov.

// This file is part of wallblog.

// wallblog is free software: you can redistribute it
// and/or modify it under the terms of the
// GNU Affero Public License as published by the
// Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.

// wallblog is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero Public License for more details.

// You should have received a copy of the GNU Affero Public License along with wallblog.
// If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/lcomrade/wallblog/internal/cfg"
	"github.com/lcomrade/wallblog/internal/handler"
	"net/http"
)

// Run HTTP server
func serveHTTP(port string, errs chan<- error) {
	errs <- http.ListenAndServe(port, nil)
}

// Run HTTPS server
func serveHTTPS(port string, cert string, key string, errs chan<- error) {
	errs <- http.ListenAndServeTLS(port, cert, key, nil)
}

// Read config and init http server
func runServer(configFile string) error {
	// Read config
	config, err := cfg.Read(configFile)
	if err != nil {
		return err
	}

	// Set handler config
	handler.Config = config

	// Add handlers
	http.HandleFunc("/", handler.Hand)

	if config.SiteMap.Enable {
		http.HandleFunc(config.SiteMap.URL, handler.SiteMapHand)
	}

	// Print init info
	println("path: config file:", configFile)
	println("path: web root:", config.WebRoot)
	println("config: site map: enable:", config.SiteMap.Enable)

	// Run server
	errs := make(chan error, 1)

	if config.HTTP.Enable {
		println("run: http:", config.HTTP.Port)
		serveHTTP(config.HTTP.Port, errs)
	}

	if config.HTTPS.Enable {
		println("run: https:", config.HTTP.Port)
		serveHTTPS(config.HTTPS.Port, config.HTTPS.Cert, config.HTTPS.Key, errs)
	}

	// If serve error
	return <-errs
}
