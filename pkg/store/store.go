package store

import "time"

// Store saves the current migration progress.
type Store interface {
	// GetLatest returns the last ID from store.
	FindLatestID() (int, error)

	// Init should prepare your store (e.g. create tables if not exists, create directory, ...).
	Init() error

	// LogUp will be called after an up migration successfully ran.
	LogUp(id int, name string, executed time.Time) error

	// LogDown will be called after a down migration successfully ran.
	LogDown(id int, name string, executed time.Time) error

	// Save will be called after an Up/Down action fails or after all migrations ran successfully.
	Save() error
}
