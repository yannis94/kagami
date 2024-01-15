package store

import (
	"database/sql"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yannis94/kagami/helpers"
	"github.com/yannis94/kagami/models"
)

type SQLiteSkill struct {
    dbPath string
    db *sql.DB
}

func NewSQLiteSkill(dbPath string) *SQLiteSkill {
    return &SQLiteSkill{dbPath: dbPath}
}

func (s *SQLiteSkill) connect() error {
    db, err := sql.Open("sqlite3", s.dbPath)
    if err != nil {
        return err
    }

    s.db = db

    return nil
}

func (s *SQLiteSkill) disconnect() error {
    return s.db.Close()
}

func (s *SQLiteSkill) Create(skill *models.Skill) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    insertQuery := "INSERT INTO skills (id, name, category, logo, level) VALUES (?, ?, ?, ?, ?);"
    _, err = s.db.Exec(insertQuery, skill.ID, skill.Name, skill.Category, skill.Logo, skill.Level)
    return err
}

func (s *SQLiteSkill) GetAll() ([]*models.Skill, error) {
    var skills []*models.Skill
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM skills;"

    rows, err := s.db.Query(query)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        var skill models.Skill

        err := rows.Scan(&skill.ID, &skill.Name, &skill.Category, &skill.Logo, &skill.Level)

        if err != nil {
            return nil, err
        }

        skills = append(skills, &skill)
    }

    return skills, nil
}

func (s *SQLiteSkill) GetByID(id string) (*models.Skill, error) {
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM skills WHERE id = ?;"
    row := s.db.QueryRow(query, id)

    var skill models.Skill

    err = row.Scan(&skill.ID, &skill.Name, &skill.Category, &skill.Logo, &skill.Level)

    if err != nil {
        return nil, err
    }

    return &skill, nil
}

func (s *SQLiteSkill) GetByCategory(category string) ([]*models.Skill, error) {
    var skills []*models.Skill
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM skills WHERE category = ?;"

    rows, err := s.db.Query(query, category)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        var skill models.Skill

        err := rows.Scan(&skill.ID, &skill.Name, &skill.Category, &skill.Logo, &skill.Level)

        if err != nil {
            return nil, err
        }

        skills = append(skills, &skill)
    }

    return skills, nil
}

func (s *SQLiteSkill) Update(skill *models.Skill) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    updateQuery := "UPDATE skills SET name = ?, category = ?, logo = ?, level = ? WHERE id = ?;"
    _, err = s.db.Exec(updateQuery, skill.Name, skill.Category, skill.Logo, skill.Level, skill.ID)
    return err
}

func (s *SQLiteSkill) Delete(id string) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    query := "SELECT * FROM skills WHERE id = ?;"
    row := s.db.QueryRow(query, id)

    var skill models.Skill

    err = row.Scan(&skill.ID, &skill.Name, &skill.Category, &skill.Logo, &skill.Level)

    if err != nil {
        return err
    }

    if skill.Logo != "" {
        fileName := path.Base(skill.Logo)
        logoPath := path.Join(helpers.UploadImgPath, fileName)
        err = helpers.DeleteFile(logoPath)

        if err != nil {
            return err
        }
    }

    deleteQuery := "DELETE FROM skills WHERE id = ?;"
    _, err = s.db.Exec(deleteQuery, id)

    return err
}

func (s *SQLiteSkill) GetCategories() ([]string, error) {
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()
    query := "SELECT DISTINCT category FROM skills;"

    rows, err := s.db.Query(query)

    if err != nil {
        return nil, err
    }

    var categories []string

    for rows.Next() {
        var category string

        err := rows.Scan(&category)

        if err != nil {
            return nil, err
        }

        categories = append(categories, category)
    }

    return categories, nil
}
