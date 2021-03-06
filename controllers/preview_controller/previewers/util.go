package previewers

import (
	"regexp"
	"strings"
)

func summarize(text string, maxWords int, maxLength int) string {
	// Normalize the whitespace to be something useful (crush it to one giant line)
	surroundingWhitespace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	interiorWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	newlines := regexp.MustCompile(`[\r\n]`)
	text = surroundingWhitespace.ReplaceAllString(text, "")
	text = interiorWhitespace.ReplaceAllString(text, " ")
	text = newlines.ReplaceAllString(text, " ")

	words := strings.Split(text, " ")
	result := text
	if len(words) >= maxWords {
		result = strings.Join(words[:maxWords], " ")
	}

	if len(result) > maxLength {
		// First try trimming off the last word
		words = strings.Split(result, " ")
		newResult := ""
		for _, word := range words {
			if len(newResult+" "+word) > maxLength {
				break
			}
			newResult = newResult + " " + word
		}
		result = newResult
	}

	if len(result) > maxLength {
		// It's still too long, just trim the thing and add an ellipsis
		result = result[:maxLength] + "..."
	}

	return result
}
