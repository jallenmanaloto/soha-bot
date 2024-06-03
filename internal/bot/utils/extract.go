package utils

import "strings"

func ExtractTitle(param []string) string {
	p := param[2:]
	title := strings.Join(p, " ")
	title = strings.ToLower(title)
	return title
}

func ExtractUrl(param string) string {
	split := strings.SplitAfter(param, "(webp)/")
	url := split[1]
	return url
}
