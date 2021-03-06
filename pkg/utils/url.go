package utils

import (
	"net/url"
	"path"
	"strings"
)

// IsValidURL https://golangcode.com/how-to-check-if-a-string-is-a-url/
func IsValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func URLJoin(baseURL string, p ...string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	u.Path = path.Join(u.Path, path.Join(p...))
	s := u.String()
	s = strings.Replace(s, "%252F", "%2F", -1)
	return s
}
