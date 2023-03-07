package main

import (
	"bufio"
	"flag"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
)

func hasQuery(urlString string) (bool, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return false, err
	}
	if u.RawQuery != "" {
		return true, nil
	}
	return false, nil
}

func main() {
	verbose := flag.Bool("v", false, "enables verbose output")

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 0, 64*1024)
	scanner.Buffer(buffer, 1024*1024)

	for scanner.Scan() {
		params, err := hasQuery(scanner.Text())
		if err != nil && *verbose {
			logrus.Errorf("Error parsing URL: %s", err)
		}

		if params {
			logrus.Println(scanner.Text())
		}
	}
}
