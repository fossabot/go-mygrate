/*The CLI for Mμgrate.

Usage

	$ mygrate [command]
	Commands:
		create name     creates a new migration
		generate        generates the mygrations.go file
		version         shows the current version

	$ mygrate create name
	Positional Arguments:
		name		the name for the migration
*/
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/stahlstift/go-mygrate/internal/generate"
)

const (
	mygrationsPath = "./mygrations"
	nameRegex      = `^[a-zA-Z]\w+[a-zA-Z0-9]$`
)

var validName = regexp.MustCompile(nameRegex)

const usage = `Usage:
%s [command]
Commands:
	create name	creates a new migration
	generate	generates the mygrations.go
	version		shows the current version
`

const createUsage = `Usage:
%s create name
Positional Arguments:
	name		the name for the migration
`

const version = "0.1.0"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	argLen := len(os.Args)

	if argLen < 2 {
		fmt.Printf(usage, os.Args[0])
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "create":
		if argLen < 3 {
			fmt.Printf(createUsage, os.Args[0])
			return
		}
		name := os.Args[2]
		if !validName.MatchString(name) {
			fmt.Printf("Given name '%s' in illegal format. Format should be '%s'\n", name, nameRegex)
			os.Exit(1)
		}
		must(generate.Init(mygrationsPath))
		must(generate.GenerateMygration(mygrationsPath, int(time.Now().Unix()), name))
		must(generate.GenerateMygrations(mygrationsPath))
		return
	case "generate":
		must(generate.Init(mygrationsPath))
		must(generate.GenerateMygrations(mygrationsPath))
		return
	case "version":
		fmt.Printf(version)
	default:
		fmt.Printf(usage, os.Args[0])
	}
}
