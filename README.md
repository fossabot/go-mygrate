# go-mygrate

## Mμgrate

Mμgrate is a micro (μ) migration framework. It has no external deps, use go stdlib features and is production ready. A mμgration is just a timestamp, a name and a pair of go functions executed in order.

Mμgrate is for anyone looking for a simple, small and handy solution to migrate things in your project up and down.

**Features**

- ship your migrations compiled in your binary
- migrate programmatically
- no need to install some outdated or abandoned ORM to migrate your database
- use the deps and driver that you are already using
- migrate whatever you want! A migration is just a pair of functions which getting called in order!
- ships with a json file based store, but you can implement your own Store  
  (sqlite will come soon)
- the cli works perfectly with go generate to register automatically new migrations files
- the cli creates new files for you

### Installation

#### CLI

This will install the mygrate command

```bash
go get github.com/stahlstift/go-migrate/cmd/mygrate
```

#### GO

Add mygrate as a dep to your project

```bash
go get github.com/stahlstift/go-migrate
```

For a complete example see the example folder

```go

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
```

### Usage

```bash
mygrate create name_your_migration
```
