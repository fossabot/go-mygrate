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

	"github.com/demaggus83/go-mygrate/internal/generate"
	"github.com/demaggus83/go-mygrate/pkg/mygrate"
)

const (
	nameRegex = `^[a-zA-Z]\w+[a-zA-Z0-9]$`
)

var validName = regexp.MustCompile(nameRegex)

const usage = `Usage:
%s [command]
Commands:
	create name	creates a new migration
	init		init the mygration folder
	version		shows the current version
`

const createUsage = `Usage:
%s create name
Positional Arguments:
	name		the name for the migration
`

const version = "0.2.0"

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
		must(generate.Init(mygrate.MygrationsPath))
		must(generate.GenerateMygration(mygrate.MygrationsPath, int(time.Now().Unix()), name))
		must(generate.GenerateMygrations(mygrate.MygrationsPath))
		return
	case "init":
		must(generate.Init(mygrate.MygrationsPath))
		must(generate.GenerateMygrations(mygrate.MygrationsPath))
		return
	case "version":
		fmt.Printf(version)
	default:
		fmt.Printf(usage, os.Args[0])
	}
}
