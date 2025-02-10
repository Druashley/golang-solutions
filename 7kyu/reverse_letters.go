package sevenkyu

import "regexp"

func ReverseLetters(s string) string {

	reversedString := ""
	stringLength := len(s)
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z]+`)

	for i := stringLength - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}
	result := nonAlphanumericRegex.ReplaceAllString(reversedString, "")

	return result
}
