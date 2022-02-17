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
	"os"
)

func errWrite(err error, rw http.ResponseWriter) {
	var httpCode int
	var errName string
	var errDesc string
	var customPageFile string

	// File not exist
	if os.IsNotExist(err) {
		httpCode = 404
		errName = "404 Not found"
		errDesc = "File not exist"
		customPageFile = "404"

		// Permission is denied
	} else if os.IsPermission(err) {
		httpCode = 500
		errName = "500 Internal Server Error"
		errDesc = "Permission denied"
		customPageFile = "500_permission_denied"

		// File read timeout
	} else if os.IsTimeout(err) {
		httpCode = 500
		errName = "500 Internal Server Error"
		errDesc = "File read timeout"
		customPageFile = "500_file_read_timeout"

		// Other errors
	} else {
		httpCode = 500
		errName = "500 Internal Server Error"
		errDesc = "Unknown"
		customPageFile = "500_unknown"
	}

	// Custom page
	pageBody := pagePart(customPageFile)

	// Default page
	if pageBody == "" {
		pageBody = `
<h1>` + errName + `</h1>
<hr />
<p>` + errDesc + `</p>
`
	}

	// Build page
	page := `
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>
		<link rel='stylesheet' type='text/css' href='/error.css'>
	</head>
	<body>
		` + pageBody + `
	</body>
</html>
`

	// Write resposnse body
	rw.WriteHeader(httpCode)
	rw.Header().Set("Content-type", "text/html")
	io.WriteString(rw, page)
}
