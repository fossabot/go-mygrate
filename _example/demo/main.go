package main

import (
	"fmt"

	"github.com/stahlstift/go-mygrate/pkg/mygrate"
	"github.com/stahlstift/mygrate-test/mygrations"
)

//go:generate mygrate generate

func migrate() error {
	myg, err := mygrate.New(&mygrations.Deps{})
	if err != nil {
		return err
	}

	mygrations.Register(myg)

	err = myg.Up(false)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := migrate()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello Migrations!")
}
