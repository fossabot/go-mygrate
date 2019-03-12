package mygrate

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)

type mygration struct {
	Name     string    `json:"name"`
	Executed time.Time `json:"executed"`
}

type fileStore struct {
	Current    int                `json:"current"`
	Mygrations map[int]*mygration `json:"mygrations"`
}

func (f *fileStore) GetLatest() int {
	return f.Current
}

func (f *fileStore) Init() error {
	buf, err := ioutil.ReadFile(".mygrate")
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, f)
	if err != nil {
		return err
	}

	return nil
}

func (f *fileStore) LogUp(id int, name string, executed time.Time) error {
	f.Mygrations[id] = &mygration{
		Name:     name,
		Executed: executed,
	}
	f.Current = id
	return nil
}

func (f *fileStore) LogDown(id int, name string, executed time.Time) error {
	delete(f.Mygrations, id)

	var keys []int
	for k := range f.Mygrations {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	f.Current = keys[len(keys)-1]
	return nil
}

func (f *fileStore) Save() error {
	buf, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".mygrate", buf, 0644)
	if err != nil {
		return err
	}
	return nil
}
