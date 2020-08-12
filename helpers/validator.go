package helpers

import "net/url"

// IsValidURL checks validity of URL
func IsValidURL(link string) bool {
	_, err := url.ParseRequestURI(link)

	if err != nil {
		return false
	}

	u, err := url.Parse(link)

	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
