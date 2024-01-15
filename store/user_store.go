package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yannis94/kagami/models"
)

type SQLiteUser struct {
    dbPath string
    db *sql.DB
}

func NewSQLiteUser(dbPath string) *SQLiteUser {
    return &SQLiteUser{ dbPath: dbPath }
}

func (s *SQLiteUser) connect() error {
    db, err := sql.Open("sqlite3", s.dbPath)
    if err != nil {
        return err
    }

    s.db = db

    return nil
}

func (s *SQLiteUser) disconnect() error {
    return s.db.Close()
}

func (s *SQLiteUser) Create(user *models.User) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    createQuery := "INSERT INTO users (id, username, password) VALUES (?, ?, ?);"
    _, err = s.db.Exec(createQuery, user.ID, user.Username, user.Password)

    return err
}

func (s *SQLiteUser) GetByUsername(username string) (*models.User, error) {
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM users WHERE username = ?;"
    row := s.db.QueryRow(query, username)

    var userFound models.User

    err = row.Scan(&userFound.ID, &userFound.Username, &userFound.Password)

    if err != nil {
        return nil, err
    }

    return &userFound, nil
}
