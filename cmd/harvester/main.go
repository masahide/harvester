package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/masahide/harvester/pkg/harvester"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	showVer = flag.Bool("version", false, "Show version")
)

func newConfig() *harvester.Config {
	c := harvester.Config{
		Concurrency:   0,
		MaxAgentConns: 50,
		User:          os.Getenv("USER"),
		Hostsfile:     "",
		ShowHostName:  false,
		ColorMode:     true,
		IgnoreHostKey: false,
		Debug:         false,
		SortPrint:     true,
		Timeout:       15 * time.Second,
	}
	flag.IntVar(&c.Concurrency, "p", c.Concurrency, "concurrency (defalut \"0\" is unlimit)")
	flag.BoolVar(&c.Debug, "debug", c.Debug, "debug outputs")
	flag.DurationVar(&c.Timeout, "timeout", c.Timeout, "maximum amount of time for the TCP connection to establish.")
	flag.Parse()

	return &c
}

func checkFlag(w io.Writer) (ret int, exit bool) {
	flag.CommandLine.SetOutput(w)
	if *showVer {
		// nolint: errcheck
		fmt.Fprintf(w, "%v, commit %v, built at %v\n", version, commit, date)
		return 0, true
	}
	if flag.NArg() == 0 {
		flag.Usage()
		// nolint: errcheck
		fmt.Fprintf(w, "example:\n$ ./harvester -h <(echo host1 host2) ls -la /etc/\n")
		return 2, true
	}
	return 0, false
}

func main() {
	p := &harvester.Harvester{Config: newConfig()}
	if ret, exit := checkFlag(os.Stdout); exit {
		os.Exit(ret)
	}
	p.Init()
	os.Exit(p.Run())
}
