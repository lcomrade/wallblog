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
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// Provides a dynamically generated sitemap
func SiteMapHand(rw http.ResponseWriter, req *http.Request) {
	// Get site URL
	siteURL := getSiteURL(req)

	// Get file list
	fileList := getAllFiles(Config.WebRoot)

	// Generate sitemap
	siteMap := ""

	for _, file := range fileList {
		pathURL := strings.TrimPrefix(file, Config.WebRoot)
		pathURL = filepath.Join("/", pathURL)

		// Skip hidden dirs and files
		if Config.SiteMap.SkipHidden {
			skip := false

			for _, part := range strings.Split(pathURL, "/") {
				if strings.HasPrefix(part, ".") {
					skip = true
					continue
				}
			}

			if skip == true {
				continue
			}
		}

		// Skip by extetion
		ext := filepath.Ext(pathURL)
		if ext != ".txt" && ext != ".html" && ext != ".htmlp" && ext != ".md" {
			continue
		}

		// Skip robots.txt
		if pathURL == "/robots.txt" {
			continue
		}

		// Skip config files
		skip := false

		for _, skipURL := range noAccessURLs {
			if pathURL == skipURL {
				skip = true
				break
			}
		}

		if skip == true {
			continue
		}

		// Add to site map
		siteMap = siteMap + "<url><loc>" + siteURL + pathURL + "</loc></url>\n"
	}

	// Prepate sitemap
	siteMap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
` + siteMap + `</urlset>
`

	// Send result
	rw.WriteHeader(200)
	rw.Header().Set("Content-type", "text/xml")
	io.WriteString(rw, siteMap)
}

// Gets list of all files recursively.
func getAllFiles(path string) []string {
	var result []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return result
	}

	for _, file := range files {
		// If directory
		if file.IsDir() {
			newPath := filepath.Join(Config.WebRoot, file.Name())
			result = append(result, getAllFiles(newPath)...)

			// If file
		} else {
			filePath := filepath.Join(path, file.Name())
			filePath, err := filepath.Abs(filePath)
			if err != nil {
				continue
			}

			result = append(result, filePath)
		}
	}

	return result
}
