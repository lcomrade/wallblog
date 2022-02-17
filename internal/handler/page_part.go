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
	"io/ioutil"
	"os"
	"path/filepath"
)

func readFile(path string) (string, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read file
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(fileByte), nil
}

func pagePart(nameWithoutExt string) string {
	basePath := filepath.Join(ServerRoot, nameWithoutExt)

	// htmlp file
	page, err := readFile(basePath + ".htmlp")
	if err == nil {
		return page
	}

	// md file
	page, err = md2html.ConvertFile(basePath + ".md")
	if err == nil {
		return page
	}

	// Not found
	return ""
}
