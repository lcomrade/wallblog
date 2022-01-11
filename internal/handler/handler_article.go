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
	"github.com/lcomrade/md2html"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Article page handler
func Article(rw http.ResponseWriter, req *http.Request) {
	// File or dir? Find path.
	path := filepath.Join(ServerRoot, "article", req.URL.Path)

	file, err := os.Stat(path)
	if err == nil && file.IsDir() {
		path = filepath.Join(path, "index.md")
	}

	// Set response header
	rw.Header().Set("Content-type", "text/html")

	// PAGE ARTICLE
	pageArticle, err := md2html.ConvertFile(path)
	if err != nil {
		// File not exist
		if os.IsNotExist(err) {
			rw.WriteHeader(404)
			pageArticle = `<h1>404 Not found</h1>`

			// Permission is denied
		} else if os.IsPermission(err) {
			rw.WriteHeader(500)
			pageArticle = `<h1>500 Internal Server Error</h1>`

			// File read timeout
		} else if os.IsTimeout(err) {
			rw.WriteHeader(500)
			pageArticle = `<h1>500 Internal Server Error</h1>`

			// Other errors
		} else {
			rw.WriteHeader(500)
			pageArticle = `<h1>500 Internal Server Error</h1>`
		}
	}

	// PAGE HEADER
	pageHeader, _ := md2html.ConvertFile(filepath.Join(ServerRoot, "header.md"))

	// PAGE FOOTER
	pageFooter, _ := md2html.ConvertFile(filepath.Join(ServerRoot, "footer.md"))

	// PAGE
	page := `
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>
		<link rel='stylesheet' type='text/css' href='/style.css'>
	</head>
	<body>
		<header>` + pageHeader + `</header>
		<article>` + pageArticle + `</article>
		<footer>` + pageFooter + `</footer>
	</body>
</html>
`

	// Write resposnse body
	io.WriteString(rw, page)
}
