package store

import (
	"database/sql"
	"fmt"
	"time"
)

// SQLStore to store the current migration status in a database.
//
// Currently there is just sqlite support, but adding support for any
// other sql.DB compatible DB should not be a problem.
type SQLStore struct {
	db        *sql.DB
	TableName string

	logUpStmt   *sql.Stmt
	logDownStmt *sql.Stmt
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:        db,
		TableName: "mygrate",
	}
}

func (s *SQLStore) Init() error {
	_, err := s.db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS "%s" (
			"id"		INTEGER PRIMARY KEY,
			"name"		TEXT NOT NULL,
			"executed"	INTEGER NOT NULL
		)`, s.TableName))
	if err != nil {
		return err
	}

	s.logUpStmt, err = s.db.Prepare(
		fmt.Sprintf(`INSERT INTO "%s" ("id", "name", "executed") VALUES (?, ?, ?)`, s.TableName),
	)
	if err != nil {
		return err
	}

	s.logDownStmt, err = s.db.Prepare(
		fmt.Sprintf(`DELETE FROM "%s" WHERE id = ?`, s.TableName),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) FindLatestID() (int, error) {
	var id int
	row := s.db.QueryRow(fmt.Sprintf(`SELECT id FROM "%s" order by ID DESC LIMIT 1`, s.TableName))
	err := row.Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
		return 0, nil
	}
	return id, nil
}

func (s *SQLStore) LogUp(id int, name string, executed time.Time) error {
	res, err := s.logUpStmt.Exec(id, name, executed.UnixNano() / 1000000)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not inserted")
	}

	return nil
}

func (s *SQLStore) LogDown(id int, name string, executed time.Time) error {
	res, err := s.logDownStmt.Exec(id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not deleted")
	}

	return nil
}

func (s *SQLStore) Save() error {
	defer s.logUpStmt.Close()
	defer s.logDownStmt.Close()

	return nil
}
