// go build -o issues && ./issues search golang
// go build -o issues && ./issues list yykhomenko test
// go build -o issues && ./issues get yykhomenko test 3
// go build -o issues && export GITHUB_USER=USER && export GITHUB_PASS=PASS && ./issues create yykhomenko test
// go build -o issues && export GITHUB_USER=USER && export GITHUB_PASS=PASS && ./issues update yykhomenko test 3
// go build -o issues && export GITHUB_USER=USER && export GITHUB_PASS=PASS && ./issues close yykhomenko test 3
package main

import (
	"fmt"
	"os"
)

var usage = "usage: issues search|list|create|read|update|close TERM|(OWNER REPO)|(OWNER REPO NUMBER)]"

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		searchIssues(args)
	} else {
		switch cmd {
		case "list":
			getIssues(args[0], args[1])
		case "get":
			getIssue(args[0], args[1], args[2])
		case "create":
			createIssue(args[0], args[1], args[2])
		case "update":
			updateIssue(args[0], args[1], args[2])
		case "close":
			closeIssue(args[0], args[1], args[2])
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %q, use: %s\n", cmd, usage)
		}
	}
}
