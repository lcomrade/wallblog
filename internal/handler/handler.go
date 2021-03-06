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

package handler

import (
	"errors"
	"github.com/lcomrade/wallblog/internal/cfg"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var errHttpCode403 = errors.New("http code: 403 Forbidden")

var Config cfg.Config

var noAccessURLs = []string{
	"/article_start.htmlp", "article_start.md",
	"/article_end.htmlp", "article_end.md",
	"/header.htmlp", "/header.md",
	"/footer.htmlp", "/footer.md",
	"/403.htmlp", "/403.md",
	"/404.htmlp", "/404.md",
	"/500_permission_denied.htmlp", "/500_permission_denied.md",
	"/500_file_read_timeout.htmlp", "/500_file_read_timeout.md",
	"/500_unknown.htmlp", "/500_unknown.md",
}

func Hand(rw http.ResponseWriter, req *http.Request) {
	if req.Host != "" {
		req.URL.Host = req.Host
	}

	// Sanitize URL
	splitOrigURL := strings.Split(req.URL.Path, "/")
	req.URL.Path = "/"

	for _, part := range splitOrigURL {
		if part == "." {
			continue
		}

		if part == ".." {
			continue
		}

		if req.URL.Path != "/" {
			req.URL.Path = req.URL.Path + "/" + part

		} else {
			req.URL.Path = "/" + part
		}
	}

	// Find path
	url := req.URL.Path
	path := filepath.Join(Config.WebRoot, url)
	ext := filepath.Ext(path)

	// Skip URLs
	for _, skipURL := range noAccessURLs {
		if url == skipURL {
			errWrite(errHttpCode403, rw, req)
			return
		}
	}

	// Redirect
	if strings.HasSuffix(url, "/index.html") {
		urlRune := []rune(url)
		newURL := string(urlRune[0:len(urlRune)-11]) + "/"
		redirectWrite(newURL, rw, req)
		return
	}

	if strings.HasSuffix(url, "/index.htmlp") {
		urlRune := []rune(url)
		newURL := string(urlRune[0:len(urlRune)-12]) + "/"
		redirectWrite(newURL, rw, req)
		return
	}

	if strings.HasSuffix(url, "/index.md") {
		urlRune := []rune(url)
		newURL := string(urlRune[0:len(urlRune)-9]) + "/"
		redirectWrite(newURL, rw, req)
		return
	}

	// Get file info
	file, err := os.Stat(path)
	if err != nil {
		errWrite(err, rw, req)
		return
	}

	// Directory
	if file.IsDir() {
		dirHand(rw, req, path)
		return
	}

	// *.md file
	if ext == ".md" {
		mdHand(rw, req, path)
		return
	}

	// *.htmlp file
	if ext == ".htmlp" {
		htmlpHand(rw, req, path)
		return
	}

	// Other files
	fileHand(rw, req, path)
}
