package store

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yannis94/kagami/models"
)

type ProjectStorage interface {
    Create(*models.Project) error
    GetAll() ([]*models.Project, error)
    GetByKeyword(keyword string) ([]*models.Project, error)
    GetByID(id string) (*models.Project, error)
    Update(*models.Project) error
    Delete(id string) error
    GetKeywords() []string
}

type SkillStorage interface {
    Create(*models.Skill) error
    GetAll() ([]*models.Skill, error)
    GetByID(id string) (*models.Skill, error)
    GetByCategory(category string) ([]*models.Skill, error)
    GetCategories() ([]string, error)
    Update(*models.Skill) error
    Delete(id string) error
}

type UserStorage interface {
    Create(user *models.User) error
    GetByUsername(username string) (*models.User, error)
}

type Store struct {
    Skill SkillStorage
    Project ProjectStorage
    User UserStorage
    db *sql.DB
}


func NewSQLiteStorage() *Store {
    var dbPath = os.Getenv("DB_PATH")
    return &Store{ Skill: NewSQLiteSkill(dbPath), Project: NewSQLiteProject(dbPath), User: NewSQLiteUser(dbPath) }
}

func (s Store) Init() {
    fmt.Println("Database init...")
    db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))

    if err != nil {
        panic(err)
    }

    defer db.Close()

    err = createSkillsTable(db)

    if err != nil {
        panic(err)
    }
    fmt.Println("Skills table created.")

    err = createProjectsTable(db)

    if err != nil {
        panic(err)
    }
    fmt.Println("Projects table created.")

    err = createKeywordsTable(db)

    if err != nil {
        panic(err)
    }
    fmt.Println("Keywords table created.")

    err = createUserTable(db)

    if err != nil {
        panic(err)
    }
    fmt.Println("Users table created.")

    fmt.Println("Database init done.")
}

func (s *Store) Connect() error {
    db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))

    if err != nil {
        return err
    }

    s.db = db

    return nil
}

func createSkillsTable(db *sql.DB) error {
    createQuery := `
        CREATE TABLE IF NOT EXISTS skills (
            id text PRIMARY KEY,
            name TEXT,
            category TEXT,
            logo TEXT,
            level INTEGER
        );
    `

    _, err := db.Exec(createQuery)
    return err
}
func createProjectsTable(db *sql.DB) error {
    createQuery := `
        CREATE TABLE IF NOT EXISTS projects (
            id text PRIMARY KEY,
            name TEXT,
            overview TEXT,
            description TEXT,
            repository TEXT
        );
    `

    _, err := db.Exec(createQuery)
    return err
}
func createKeywordsTable(db *sql.DB) error {
    createQuery := `
        CREATE TABLE IF NOT EXISTS keywords (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            project_id TEXT,
            keyword TEXT,
            FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
        );
    `

    _, err := db.Exec(createQuery)
    return err
}
func createUserTable(db *sql.DB) error {
    createQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id TEXT,
            username TEXT,
            password TEXT
        );
    `

    _, err := db.Exec(createQuery)
    return err
}
