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
	head := ""

	// Head: Page title
	if Config.Page.AutoTitle.Enable {
		title := getHtmlHeader(pageArticle)
		// If auto generated title is not empty
		if title != "" {
			head = head + "\n<title>" + Config.Page.AutoTitle.Prefix + title + Config.Page.AutoTitle.Sufix + "</title>"

			// If auto generated title is empty
		} else {
			// If default title exists
			if Config.Page.AutoTitle.Default != "" {
				head = head + "\n<title>" + Config.Page.AutoTitle.Default + "</title>"
			}
		}
	}

	// Body: Page header
	pageHeader := pagePart("header")

	// Body: Page footer
	pageFooter := pagePart("footer")

	// PAGE
	page := `
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>` + head + `
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
