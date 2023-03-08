package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vitorfhc/hacks/queryxss/rxss"
)

func logNonFatalError(err error) {
	if logrus.GetLevel() < logrus.InfoLevel {
		return
	}

	if err != nil {
		logrus.Error(err)
	}
}

func main() {
	verbose := flag.Bool("v", false, "enables verbose output")
	debug := flag.Bool("d", false, "enables debug output")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug output enabled\n")
	} else if *verbose {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Debug("Verbose output enabled\n")
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 0, 64*1024)
	scanner.Buffer(buffer, 1024*1024)

	for scanner.Scan() {
		rxssScanner, err := rxss.NewScanner(scanner.Text())
		if err != nil {
			logNonFatalError(err)
			continue
		}
		found, err := rxssScanner.Scan()
		if err != nil {
			logNonFatalError(err)
			continue
		}
		if found {
			fmt.Println(scanner.Text())
		}
	}
}
