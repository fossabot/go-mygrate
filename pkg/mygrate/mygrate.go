// Package mygrate provides the programmatically interface for Mμgrate
package mygrate

import (
	"fmt"
	"sort"
	"time"

	"github.com/demaggus83/go-mygrate/internal/generate"
	"github.com/demaggus83/go-mygrate/pkg/store"
)

const MygrationsPath = "./mygrations"

// Generate is a helper func to be used bye go generate
func Generate() error {
	err := generate.Init(MygrationsPath)
	if err != nil {
		return err
	}

	err = generate.GenerateMygrations(MygrationsPath)
	if err != nil {
		return err
	}

	return nil
}

type mygrationFunc = func(deps interface{}) error

type mygrationsFuncs struct {
	name string
	up   mygrationFunc
	down mygrationFunc
}

type Mygrate struct {
	mygrations map[int]*mygrationsFuncs
	Store      store.Store
	dep        interface{}
}

func New(dep interface{}) (*Mygrate, error) {
	return NewWithStore(dep, store.NewFileStore())
}

func NewWithStore(dep interface{}, st store.Store) (*Mygrate, error) {
	if err := st.Init(); err != nil {
		return nil, err
	}

	return &Mygrate{
		mygrations: make(map[int]*mygrationsFuncs, 0),
		Store:      st,
		dep:        dep,
	}, nil
}

func (m *Mygrate) Register(id int, name string, up mygrationFunc, down mygrationFunc) {
	m.mygrations[id] = &mygrationsFuncs{
		name: name,
		up:   up,
		down: down,
	}
}

func (m *Mygrate) getKeysAsc() []int {
	var keys []int
	for k := range m.mygrations {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (m *Mygrate) getKeysDesc() []int {
	var keys []int
	for k := range m.mygrations {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func (m *Mygrate) up(id int, name string, fn mygrationFunc) error {
	err := fn(m.dep)
	if err != nil {
		_ = m.Store.Save()
		return fmt.Errorf("Migration_Up with id '%d' returned err '%s'", id, err)
	}

	err = m.Store.LogUp(id, name, time.Now())
	if err != nil {
		_ = m.Store.Save()
		return err
	}

	return nil
}

func (m *Mygrate) down(id int, name string, fn mygrationFunc) error {
	err := fn(m.dep)
	if err != nil {
		_ = m.Store.Save()
		return fmt.Errorf("Migration_Down with id '%d' returned err '%s'", id, err)
	}

	err = m.Store.LogDown(id, name, time.Now())
	if err != nil {
		_ = m.Store.Save()
		return err
	}

	return nil
}

func (m *Mygrate) redoLast(keys []int) error {
	id := keys[len(keys)-1]

	err := m.down(id, m.mygrations[id].name, m.mygrations[id].down)
	if err != nil {
		return err
	}

	err = m.up(id, m.mygrations[id].name, m.mygrations[id].up)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mygrate) Up(redoLast bool) error {
	keys := m.getKeysAsc()

	latestID, err := m.Store.FindLatestID()
	if err != nil {
		return err
	}

	changes := 0
	for _, id := range keys {
		if id <= latestID {
			continue
		}

		err := m.up(id, m.mygrations[id].name, m.mygrations[id].up)
		if err != nil {
			return err
		}

		changes++
	}

	if changes == 0 && redoLast && len(m.mygrations) >= 1 {
		err := m.redoLast(keys)
		if err != nil {
			return err
		}
	}

	err = m.Store.Save()
	if err != nil {
		return err
	}

	return nil
}

func (m *Mygrate) DownTo(targetID int) error {
	keys := m.getKeysDesc()

	for _, id := range keys {
		if id < targetID {
			break
		}

		err := m.down(id, m.mygrations[id].name, m.mygrations[id].down)
		if err != nil {
			return err
		}
	}

	err := m.Store.Save()
	if err != nil {
		return err
	}

	return nil
}
