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
	"bytes"
	"github.com/lcomrade/md2html/v2"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

type templateData struct {
	URL templateDataURL
}

type templateDataURL struct {
	Path string
	Full string
}

func useTemplate(text string, req *http.Request) string {
	// Prepare template data
	tmplData := templateData{
		URL: templateDataURL{
			Path: req.URL.Path,
			Full: getFullURL(req),
		},
	}

	// New template
	tmpl, err := template.New("Template").Parse(text)
	if err != nil {
		return text
	}

	// Execute template
	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, tmplData)
	if err != nil {
		return text
	}

	text = buffer.String()

	// Escape
	text = strings.Replace(text, `\{`, "{", -1)
	text = strings.Replace(text, `\}`, "}", -1)

	return text
}

func pagePart(nameWithoutExt string, req *http.Request) string {
	basePath := filepath.Join(Config.WebRoot, nameWithoutExt)

	// htmlp file
	page, err := readFile(basePath + ".htmlp")
	if err == nil {
		// Template mode
		if Config.Page.EnableTemplateMode == true {
			page = useTemplate(page, req)
		}

		// Return
		return page
	}

	// md file
	page, err = readFile(basePath + ".md")
	if err == nil {
		// Template mode
		if Config.Page.EnableTemplateMode == true {
			page = useTemplate(page, req)
		}

		// Return
		return md2html.Convert(page)
	}

	// Not found
	return ""
}
