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
	// Get host
	host := req.Host

	if Config.Overwrite.Host != "" {
		host = Config.Overwrite.Host
	}

	// Get protocol
	protocol := "http"

	if Config.HTTPS.Enable == true {
		protocol = "https"
	}
	
	if Config.Overwrite.Protocol != "" {
		protocol = Config.Overwrite.Protocol
	}

	// Get site URL
	siteURL := protocol + "://" + host

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

		// Skip config files
		skipURLs := []string{
			"/header.htmlp", "/header.md",
			"/footer.htmlp", "/footer.md",
			"/404.htmlp", "/404.md",
			"/500_permission_denied.htmlp", "/500_permission_denied.md",
			"/500_file_read_timeout.htmlp", "/500_file_read_timeout.md",
			"/500_unknown.htmlp", "/500_unknown.md",
			"/robots.txt",
		}

		skip := false

		for _, skipURL := range skipURLs {
			if pathURL == skipURL {
				skip = true
				continue
			}
		}

		if skip == true {
			continue
		}

		// Add to site map
		siteMap = siteMap + "\n<url><loc>" + siteURL + pathURL + "</loc></url>"
	}

	// Prepate sitemap
	siteMap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset>` + siteMap + `
</urlset>
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
