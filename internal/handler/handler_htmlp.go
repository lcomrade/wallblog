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
	"net/http"
)

// HTMLP page handler
func htmlpHand(rw http.ResponseWriter, req *http.Request, path string) {
	// Read page article
	pageArticle, err := readFile(path)
	if err != nil {
		errWrite(err, rw)
		return
	}

	// Send page
	pageWrite(pageArticle, rw, req)
}
