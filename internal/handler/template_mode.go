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
	"net/http"
	"strings"
	"text/template"
	"time"
)

type templateData struct {
	Request templateDataRequest
	Time    templateDataTime
}

type templateDataRequest struct {
	ClientIP string
	Method   string
	URL      templateDataURL
	Header   templateDataHeader
}

type templateDataURL struct {
	Protocol string
	Path     string
	RawQuery string
	Full     string
}

type templateDataHeader struct {
	AcceptLanguage string
	Host           string
	Referer        string
	UserAgent      string
}

type templateDataTime struct {
	Unix    int64
	Year    int
	Month   int
	Day     int
	Weekday int
	Hour    int
	Minute  int
	Second  int
}

func useTemplate(text string, req *http.Request) string {
	// Get time
	timeNow := time.Now().UTC()

	// Prepare template data
	tmplData := templateData{
		Request: templateDataRequest{
			ClientIP: strings.Split(req.RemoteAddr, ":")[0],
			Method:   req.Method,
			URL: templateDataURL{
				Protocol: req.URL.Scheme,
				Path:     req.URL.Path,
				RawQuery: req.URL.RawQuery,
				Full:     getSiteURL(req) + req.URL.Path,
			},
			Header: templateDataHeader{
				AcceptLanguage: strings.Split(req.Header.Get("Accept-Language"), ";")[0],
				Host:           req.Host,
				Referer:        req.Header.Get("Referer"),
				UserAgent:      req.Header.Get("User-Agent"),
			},
		},
		Time: templateDataTime{
			Unix:    timeNow.Unix(),
			Year:    timeNow.Year(),
			Month:   int(timeNow.Month()),
			Day:     timeNow.Day(),
			Weekday: int(timeNow.Weekday()),
			Hour:    timeNow.Hour(),
			Minute:  timeNow.Minute(),
			Second:  timeNow.Second(),
		},
	}

	// Get scheme
	if tmplData.Request.URL.Protocol == "" {
		if req.URL.Port() == "443" {
			tmplData.Request.URL.Protocol = "https"

		} else {
			tmplData.Request.URL.Protocol = "http"
		}
	}

	// Get full URL
	if req.URL.RawQuery != "" {
		tmplData.Request.URL.Full = tmplData.Request.URL.Full + "?" + req.URL.RawQuery
	}

	// Get accept language
	for _, lang := range strings.Split(tmplData.Request.Header.AcceptLanguage, ",") {
		if len(lang) == 2 {
			tmplData.Request.Header.AcceptLanguage = lang
			break
		}
	}

	// New template
	tmpl, err := template.New("Template").Parse(text)
	if err != nil {
		return err.Error()
	}

	// Execute template
	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, tmplData)
	if err != nil {
		return err.Error()
	}

	text = buffer.String()

	// Escape
	text = strings.Replace(text, `\{`, "{", -1)
	text = strings.Replace(text, `\}`, "}", -1)

	return text
}
