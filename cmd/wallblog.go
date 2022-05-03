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

package main

import (
	"flag"
	"os"
)

func printHelp() {
	println("Usage:", os.Args[0], "[OPTION]...")
	println("Lightweight blogging engine with markdown support.")
	println("")
	println("-config   path to config file")
	println("-help     display this help and exit")
}

func main() {
	// Read cmd args
	flag.Usage = printHelp

	flagConfig := flag.String("config", "/etc/wallblog/config.json", "")
	flagHelp := flag.Bool("help", false, "")

	flag.Parse()

	// -help flag
	if *flagHelp == true {
		printHelp()
		os.Exit(0)
	}

	// Run WEB server
	err := runServer(*flagConfig)
	if err != nil {
		panic(err)
	}
}
