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

package cfg

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	WebRoot string
	HTTP    ConfigHTTP
	HTTPS   ConfigHTTPS
	SiteMap ConfigSiteMap
}

type ConfigHTTP struct {
	Enable bool
	Port   string
}

type ConfigHTTPS struct {
	Enable bool
	Port   string
	Cert   string
	Key    string
}

type ConfigSiteMap struct {
	Enable     bool
	URL        string
	SkipHidden bool
}

func Read(path string) (Config, error) {
	// Default
	config := Config{
		WebRoot: "/var/lib/wallblog",
		HTTP: ConfigHTTP{
			Enable: true,
			Port:   ":80",
		},
		HTTPS: ConfigHTTPS{
			Enable: false,
			Port:   ":443",
			Cert:   "",
			Key:    "",
		},
		SiteMap: ConfigSiteMap{
			Enable:     true,
			URL:        "/sitemap.xml",
			SkipHidden: true,
		},
	}

	// Open file
	file, err := os.Open(path)
	if err != nil {
		// If config file not found
		if os.IsNotExist(err) {
			return config, nil

			// Else
		} else {
			return config, err
		}
	}
	defer file.Close()

	// Decode config
	parser := json.NewDecoder(file)
	err = parser.Decode(&config)
	if err != nil {
		return config, err
	}

	// Make web root path absolute
	config.WebRoot, err = filepath.Abs(config.WebRoot)
	if err != nil {
		return config, err
	}

	return config, nil
}
