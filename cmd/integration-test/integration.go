package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"

	"github.com/projectdiscovery/notify/internal/testutils"
)

var (
	providerConfig = flag.String("provider-config", "", "provider config to use for testing")
	debug          = os.Getenv("DEBUG") == "true"
	isDependabot   = os.Getenv("DEPENDABOT") == "true"
	errored        = false
	success        = aurora.Green("[✓]").String()
	failed         = aurora.Red("[✘]").String()
	testCases      = map[string]testutils.TestCase{
		"discord": &discord{},
		"slack":   &slack{},
		"custom":  &custom{},
		//		"telegram": &telegram{},
		//		"teams":    &teams{},
		//		"smtp":     &smtp{},
		//		"pushover": &pushover{},
		"gotify": &gotify{},
	}
)

func main() {
	flag.Parse()

	for name, test := range testCases {
		// run only gotify test for dependabot
		if isDependabot && name != "gotify" {
			continue
		}
		fmt.Printf("Running test cases for \"%s\"\n", aurora.Blue(name))
		err := test.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s Test \"%s\" failed: %s\n", failed, name, err)
			errored = true
		} else {
			fmt.Printf("%s Test \"%s\" passed!\n", success, name)
		}
	}
	if errored {
		os.Exit(1)
	}
}
