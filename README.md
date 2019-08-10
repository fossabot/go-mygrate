# go-mygrate

## Mμgrate

Mμgrate is a micro (μ) migration framework. It has no external deps, use go stdlib features and is production ready. A mμgration is just a timestamp, a name and a pair of go functions executed in the same order.

Mμgrate is for anyone looking for a simple, small and handy solution to migrate things in a project up and/or down.

**Features**

- ship your migrations compiled in your binary
- migrate programmatically
- no deps on some outdated or abandoned ORMs to migrate your database
- use the deps and driver that you are already using in your project
- migrate whatever you want! A migration is just a pair of functions which getting called in same order!
- ships with a FileStore (json based) and a SQLStore - but it's easy to implement your own custom store
- mygrate supports go´s generate to register new migrations files automatically
- the cli can create new mygration files for you

### Installation

#### CLI

This will install the mygrate command

```bash
go get github.com/demaggus83/go-mygrate/cmd/mygrate
```

#### GO

Add mygrate as a dependency to your project

```bash
go get github.com/demaggus83/go-mygrate
```

For a complete example see the example folder

```go

//go:generate go run mygrations/generate.go

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
```

### Usage

```bash
mygrate create name_your_migration
```

### Changelog

#### 0.2.2
+ fixing some little issues

#### 0.2.0
+ some cleanups and refactorings
+ added SQLStore
+ mygrate will now use the mygration.tpl   
	this allows to modify the template of a new migration which can be usefully for typecasting the dep to a specific factory or struct   
	If the file is not present mygration will generate a default mygration.tpl
+ mygrate will create now a mygrations/generate/mygrations.go file to be used with go:generate   
	this removes the requirement to install the cli tool to generate the mygrations.go file. (e.G. CI builds)

#### 0.1.0
+ init
