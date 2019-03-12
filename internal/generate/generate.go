package generate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"text/template"
	"time"
)

type tplMygration struct {
	ID   int
	Name string
	Up   string
	Down string
}

type tplVarForMygrations struct {
	Time       time.Time
	Mygrations []*tplMygration
}

type tplVarForMygration struct {
	ID   int
	Name string
}

var (
	templateMygrationsGo = template.Must(
		template.New("mygrations.go").Parse(`// Code generated by mygrate; DO NOT EDIT.
// {{ .Time }}

// Package mygrations provides the migration files for the project
package mygrations

import (
	"github.com/stahlstift/go-mygrate/pkg/mygrate"
)

func Register(m *mygrate.Mygrate) {
{{- range .Mygrations}}
	m.Register({{ .ID }}, "{{ .Name }}", {{ .Up }}, {{ .Down }})
{{- end}}
}

`))

	templateMygrationFile = template.Must(
		template.New("mygration.go").Parse(`package mygrations

import (
	"errors"
)

func Mygration_{{ .ID }}_{{ .Name }}_Up(dep interface{}) error {
	return errors.New("need to implement")
}

func Mygration_{{ .ID }}_{{ .Name }}_Down(dep interface{}) error {
	return errors.New("need to implement")
}

`))
)

func Init(mygrationsPath string) error {
	err := ensureDir(mygrationsPath)
	if err != nil {
		return err
	}

	return nil
}

func GenerateMygration(mygrationsPath string, id int, name string) error {
	buf := new(bytes.Buffer)
	err := templateMygrationFile.Execute(buf, &tplVarForMygration{
		ID:   id,
		Name: name,
	})
	if err != nil {
		return err
	}

	filePath := filepath.Join(mygrationsPath, fmt.Sprintf("%d_%s.go", id, name))
	err = ioutil.WriteFile(filePath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func GenerateMygrations(mygrationsPath string) error {
	mygrationFiles := findMygrationsInDir(mygrationsPath)
	sort.Strings(mygrationFiles)

	var tplMygrations []*tplMygration
	for _, p := range mygrationFiles {
		id, name, err := parseIDAndName(p)
		if err != nil {
			continue
		}

		tplMygrations = append(tplMygrations, &tplMygration{
			ID:   id,
			Name: name,
			Up:   fmt.Sprintf("Mygration_%d_%s_Up", id, name),
			Down: fmt.Sprintf("Mygration_%d_%s_Down", id, name),
		})
	}

	buf := new(bytes.Buffer)
	err := templateMygrationsGo.Execute(buf, &tplVarForMygrations{
		Time:       time.Now(),
		Mygrations: tplMygrations,
	})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(mygrationsPath, "mygrations.go"), buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
