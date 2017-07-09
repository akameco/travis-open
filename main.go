package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	gopen "github.com/petermbenjamin/go-open"
	"github.com/tj/docopt"
)

// Version is the package version
var Version = "0.0.1"

// Usage is the package usage infomation
const Usage = `
  Usage:
    travis-open
    travis-open -h | --help
    travis-oepn --version

  Example:
    $ travis-open 
`

const baseURL = "https://travis-ci.org/"

func main() {
	_, err := docopt.Parse(Usage, nil, true, Version, false)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	repoURL := out.String()
	re := strings.NewReplacer("ssh://git@github.com/", "", "git@github.com:", "", ".git", "")
	repoName := strings.TrimSpace(re.Replace(repoURL))

	if repoName != "" {
		gopen.Open(baseURL + repoName)
	}
}
