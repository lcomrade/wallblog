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
	"path/filepath"
)

// Markdown page handler
func mdHand(rw http.ResponseWriter, path string) {
	// Set response header
	rw.Header().Set("Content-type", "text/html")

	// PAGE ARTICLE
	pageArticle, err := md2html.ConvertFile(path)
	if err != nil {
		errWrite(err, rw)
		return
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
