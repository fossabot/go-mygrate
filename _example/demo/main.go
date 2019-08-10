package main

import (
	"fmt"

	"github.com/demaggus83/go-mygrate/mygrations" // this is the generated file from mygration

	"github.com/demaggus83/go-mygrate/pkg/mygrate"
)

//go:generate go run mygrations/generate/mygrations.go

func migrate() error {
	someFactory := "this could be a factory"

	myg, err := mygrate.New(someFactory)
	if err != nil {
		return err
	}

	// Register your migrations with mygration
	mygrations.Register(myg)

	// Run your migrations
	// redoLast could be usefully in development
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
