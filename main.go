package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type options struct {
	critPct float64 `short:"c" long:"critical-pct" description:"Committed_AS/CommitLimit * 100 %" default:"1" required:"true"`
	warnPct float64 `short:"w" long:"warning-pct" description:"Committed_AS/CommitLimit * 100 %" default:"1" required:"true"`
}

var opts options

func parseMemInfo(path string) (float64, error) {
	fp, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	var commitAs float64
	var commitLimit float64
	for scanner.Scan() {
		t := scanner.Text()
		switch strings.Fields(t)[0] {
		case "CommitLimit:":
			commitLimit, err = strconv.ParseFloat(strings.Fields(t)[1], 64)
			if err != nil {
				return 0, err
			}
		case "Committed_AS:":
			commitAs, err = strconv.ParseFloat(strings.Fields(t)[1], 64)
			if err != nil {
				return 0, err
			}
		}
	}
	return commitAs / commitLimit * 100, nil
}

func parseArgs(args []string) error {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chkSt := checkers.OK
	msg := "OK"
	result, err := parseMemInfo("/proc/meminfo")
	if err != nil {
		os.Exit(1)
	}

	if opts.warnPct < result {
		chkSt = checkers.WARNING
		msg = fmt.Sprintf("[WARN] VirtMem Usage:%f %s", result, "%")
	}

	if opts.critPct < result {
		chkSt = checkers.CRITICAL
		msg = fmt.Sprintf("[CRIT] VirtMem Usage:%f %s", result, "%")
	}

	checkers.NewChecker(chkSt, msg).Exit()
}
