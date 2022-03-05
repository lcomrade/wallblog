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

// Returns the content of the first HTML '<hX>' tag.
func getHtmlHeader(htmlText string) string {
	textRune := []rune(htmlText)
	textLen := len(textRune)

	var title string = ""

	var skip int = 0
	var hTagPart int = 0 // 1(<h1), 2(>), 3(title)
	var otherTagOpen bool = false

	for i := range textRune {
		// Skip
		if skip != 0 {
			skip = skip - 1
			continue
		}

		char := string(textRune[i])
		nextChar := " "
		nextNextChar := " "
		nextNextNextChar := " "
		nextNextNextNextChar := " "

		// Get next char
		if textLen > i+1 {
			nextChar = string(textRune[i+1])
		}

		// Get next next char
		if textLen > i+2 {
			nextNextChar = string(textRune[i+2])
		}

		// Get next next next char
		if textLen > i+3 {
			nextNextNextChar = string(textRune[i+3])
		}

		// Get next next next next char
		if textLen > i+4 {
			nextNextNextNextChar = string(textRune[i+4])
		}

		// <h1
		if hTagPart == 0 {
			if char+nextChar+nextNextChar == "<h1" {
				hTagPart = 1
				skip = 3
			}

			// >
		} else if hTagPart == 1 {
			if char == ">" {
				hTagPart = 2
			}

			// Read title and wait </h1>
		} else {
			// Exit
			if char+nextChar+nextNextChar+nextNextNextChar+nextNextNextNextChar == "</h1>" {
				return title
			}

			// Other tag open
			if char == "<" {
				otherTagOpen = true
				continue
			}

			// Other tag close
			if char == ">" {
				otherTagOpen = false
				continue
			}

			// If other tag
			if otherTagOpen == true {
				continue
			}

			// Read title
			title = title + string(char)
		}
	}

	return ""
}
