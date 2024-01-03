package rxss

import (
	"net/http"
	"net/url"
)

type Scanner struct {
	url       string
	parsedUrl *url.URL
}

func NewScanner(urlRaw string) (*Scanner, error) {
	parsedUrl, err := url.Parse(urlRaw)
	if err != nil {
		return nil, err
	}

	scanner := &Scanner{
		url:       urlRaw,
		parsedUrl: parsedUrl,
	}

	return scanner, nil
}

func (s *Scanner) Scan() (bool, error) {
	found, err := s.ScanCurrentKeyValues()
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	found, err = s.ScanRandomKeyAndValue()
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	found, err = s.ScanHeaders()
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	return false, nil
}

// ScanRandomKeyAndValue adds a random key and value to the query of the URL,
// sends a GET request and checks if the key and value are reflected in the response.
func (s *Scanner) ScanRandomKeyAndValue() (bool, error) {
	randomKey := RandomString(10)
	randomValue := RandomString(10)

	parsedUrlCopy := *s.parsedUrl
	q := parsedUrlCopy.Query()
	q.Set(randomKey, randomValue)
	parsedUrlCopy.RawQuery = q.Encode()
	resp, err := http.Get(parsedUrlCopy.String())
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	found, err := FindAnyReflections(resp, []string{randomKey, randomValue})
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}

	return false, nil
}

// ScanCurrentKeyValues finds all reflected keys or values in the body or headers.
// It does not add any new keys or values to the query, instead it uses the ones that are already there.
func (s *Scanner) ScanCurrentKeyValues() (bool, error) {
	parsedUrlCopy := *s.parsedUrl
	q := parsedUrlCopy.Query()
	resp, err := http.Get(s.url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	for key, values := range q {
		for _, value := range values {
			found, err := FindAnyReflections(resp, []string{key, value})
			if err != nil {
				return false, err
			}
			if found {
				return true, nil
			}
		}
	}

	return false, nil
}

// ScanHeaders checks if any of the headers are reflected in the response body.
func (s *Scanner) ScanHeaders() (bool, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			found, err := FindAnyReflections(resp, []string{key, value})
			if err != nil {
				return false, err
			}
			if found {
				return true, nil
			}
		}
	}

	return false, nil
}
