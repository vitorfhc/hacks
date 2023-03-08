package rxss

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

// Get sends a GET request to the specified URL and returns the response.
// If the response status code is >= 400, an error is returned.
func Get(url string) (*req.Response, error) {
	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}
	if resp.Response.StatusCode >= 400 {
		return nil, fmt.Errorf("error sending request: %s", resp.Response.Status)
	}
	return resp, nil
}

// RandomString returns a random string of the specified length.
// The string is composed of lowercase and uppercase letters.
func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// CaseInsensitiveContains returns true if the string s contains the string substr,
// ignoring case.
// Example: CaseInsensitiveContains("Hello World", "hello") returns true.
func CaseInsensitiveContains(s string, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}

// FindBodyReflections finds reflected strings in the response body.
func FindBodyReflections(resp *req.Response, reflections []string) (bool, error) {
	found := false

	bodyString := make([]byte, 1024*1024)
	_, err := resp.Body.Read(bodyString)
	if err != nil {
		return false, err
	}

	for _, reflection := range reflections {
		found = found || CaseInsensitiveContains(string(bodyString), reflection)
		if found {
			break
		}
	}

	return found, nil
}

// FindHeadersReflections finds reflected strings in the response headers.
func FindHeadersReflections(resp *req.Response, reflections []string) (bool, error) {
	found := false

	for _, reflection := range reflections {
		headers := resp.Response.Header
		for headerKey, headerValue := range headers {
			reflected := CaseInsensitiveContains(headerKey, reflection)
			logrus.Debugf("Header key: %s, reflection: %s, reflected: %t", headerKey, reflection, reflected)
			if reflected {
				found = true
				break
			}
			for _, value := range headerValue {
				reflected = CaseInsensitiveContains(value, reflection)
				if reflected {
					found = true
					break
				}
			}
		}
		if found {
			break
		}
	}

	return found, nil
}

func FindAnyReflections(resp *req.Response, reflections []string) (bool, error) {
	found, err := FindHeadersReflections(resp, reflections)
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	found, err = FindBodyReflections(resp, reflections)
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	return false, nil
}
