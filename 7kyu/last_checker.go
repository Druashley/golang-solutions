package sevenkyu

import (
	"strings"
)

func LastChecker(str, ending string) bool {

	return strings.HasSuffix(str, ending)

}
