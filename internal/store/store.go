package store

import (
	"database/sql"

	"github.com/myntdeveloper/pseudo-cli/internal/models"
	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func New(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS pseudonyms (id INTEGER PRIMARY KEY, name TEXT UNIQUE, command TEXT, description TEXT, tag TEXT)`)
	return &Store{db: db}, err
}

func (s *Store) Save(name string, command string, description string, tag string) error {
	_, err := s.db.Exec(
		"INSERT INTO pseudonyms (name, command, description, tag) VALUES ($1, $2, $3, $4)",
		name, command, description, tag,
	)
	return err
}

func (s *Store) List(tag string) ([]models.Pseudonym, error) {
	var rows *sql.Rows
	var err error
	if tag == "" {
		rows, err = s.db.Query("SELECT name, command, description, tag FROM pseudonyms")
	} else {
		rows, err = s.db.Query("SELECT name, command, description, tag FROM pseudonyms WHERE tag = $1", tag)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pseudonyms []models.Pseudonym
	for rows.Next() {
		var p models.Pseudonym
		if err := rows.Scan(&p.Name, &p.Command, &p.Description, &p.Tag); err != nil {
			return nil, &ErrNotFoundByTag{Tag: tag}
		}
		pseudonyms = append(pseudonyms, p)
	}
	return pseudonyms, nil
}

func (s *Store) GetByName(name string) (models.Pseudonym, error) {
	row := s.db.QueryRow("SELECT name, command, description, tag FROM pseudonyms WHERE name = $1", name)

	var p models.Pseudonym
	if err := row.Scan(&p.Name, &p.Command, &p.Description, &p.Tag); err != nil {
		return models.Pseudonym{}, &ErrNotFoundByName{Name: name}
	}
	return p, nil

}

func (s *Store) Delete(name string) error {
	_, err := s.db.Exec("DELETE FROM pseudonyms WHERE name = $1", name)
	return err
}
