package main

import (
	"github.com/projectdiscovery/gologger"
	"ktbs.dev/mubeng/internal/runner"
)

// hard coded to run on Windows
func main() {
	opt := runner.Options()
	opt.Address = "127.0.0.1:23528"
	opt.File = "./proxies.txt" // US IPv4

	if err := runner.New(opt); err != nil {
		gologger.Fatal().Msgf("Error! %s", err)
	}
}
