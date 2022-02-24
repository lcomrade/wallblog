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
	"net/http"
)

func pageWrite(pageArticle string, rw http.ResponseWriter) {
	// PAGE HEADER
	pageHeader := pagePart("header")

	// PAGE FOOTER
	pageFooter := pagePart("footer")

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
	rw.WriteHeader(200)
	rw.Header().Set("Content-type", "text/html")
	io.WriteString(rw, page)
}