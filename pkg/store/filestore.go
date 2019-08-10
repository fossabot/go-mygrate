package store

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

type FileStore struct {
	Current    int                `json:"current"`
	Mygrations map[int]*mygration `json:"mygrations"`
}

func NewFileStore() *FileStore {
	return &FileStore{
		Mygrations: make(map[int]*mygration),
	}
}

func (f *FileStore) Init() error {
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

func (f *FileStore) FindLatestID() (int, error) {
	return f.Current, nil
}

func (f *FileStore) LogUp(id int, name string, executed time.Time) error {
	f.Mygrations[id] = &mygration{
		Name:     name,
		Executed: executed,
	}
	f.Current = id
	return nil
}

func (f *FileStore) LogDown(id int, name string, executed time.Time) error {
	delete(f.Mygrations, id)

	var keys []int
	for k := range f.Mygrations {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	f.Current = keys[len(keys)-1]
	return nil
}

func (f *FileStore) Save() error {
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
