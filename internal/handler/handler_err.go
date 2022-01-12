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
	var errName string
	var errDesc string

	// File not exist
	if os.IsNotExist(err) {
		rw.WriteHeader(404)
		errName = "404 Not found"
		errDesc = "File not exist"

		// Permission is denied
	} else if os.IsPermission(err) {
		rw.WriteHeader(500)
		errName = "500 Internal Server Error"
		errDesc = "Permission denied"

		// File read timeout
	} else if os.IsTimeout(err) {
		rw.WriteHeader(500)
		errName = "500 Internal Server Error"
		errDesc = "File read timeout"

		// Other errors
	} else {
		rw.WriteHeader(500)
		errName = "500 Internal Server Error"
		errDesc = "Unknown"
	}

	// PAGE
	page := `
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset='utf-8'>
	</head>
	<body>
		<h1>` + errName + `</h1>
		<hr />
		<p>` + errDesc + `</p>
	</body>
</html>
`

	// Write resposnse body
	rw.Header().Set("Content-type", "text/html")
	io.WriteString(rw, page)
}
