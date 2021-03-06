// Deps prints all packages transitive dependencies.
// go run deps.go os
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	p := ancestors(packages(os.Args[1:]))
	fmt.Println(strings.Join(p, "\n"))
}

func packages(names []string) []string {
	args := []string{"list", `-f={{.ImportPath}}`}
	for _, name := range names {
		args = append(args, name)
	}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	return strings.Split(string(out), " ")
}

func ancestors(names []string) (pkgs []string) {
	seen := make(map[string]bool)
	args := []string{"list", `-f={{join .Imports " "}} {{join .Deps " "}}`}
	for _, name := range names {
		args = append(args, strings.TrimSpace(name))
	}

	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	names = strings.Split(strings.TrimSpace(string(out)), " ")
	for _, name := range names {
		seen[name] = true
	}

	for pkg := range seen {
		pkgs = append(pkgs, pkg)
	}

	sort.Strings(pkgs)

	return
}
