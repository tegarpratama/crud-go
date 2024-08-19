package helper

import "github.com/microcosm-cc/bluemonday"

func SanitizeInput(input string) string {
	p := bluemonday.UGCPolicy()
	return p.Sanitize(input)
}