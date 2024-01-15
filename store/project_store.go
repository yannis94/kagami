package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yannis94/kagami/models"
)

type SQLiteProject struct {
    dbPath string
    db *sql.DB
}

func NewSQLiteProject(dbPath string) *SQLiteProject {
    return &SQLiteProject{ dbPath: dbPath }
}

func (s *SQLiteProject) connect() error {
    db, err := sql.Open("sqlite3", s.dbPath)
    if err != nil {
        return err
    }

    s.db = db

    return nil
}

func (s *SQLiteProject) disconnect() error {
    return s.db.Close()
}


func (s *SQLiteProject) Create(project *models.Project) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    insertQuery := "INSERT INTO projects (id, name, overview, description, repository) VALUES (?, ?, ?, ?, ?);"
    _, err = s.db.Exec(insertQuery, project.ID, project.Name, project.Overview, project.Description, project.Repository)

    for _, keyword := range project.Keywords {
        insertKeywordQuery := "INSERT INTO keywords (project_id, keyword) VALUES (?, ?);"
        _, err := s.db.Exec(insertKeywordQuery, project.ID, keyword)

        if err != nil {
            return err
        }
    }

    return err
}

func (s *SQLiteProject) GetAll() ([]*models.Project, error) {
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM projects;"
    rows, err := s.db.Query(query)

    if err != nil {
        return nil, err
    }

    var projects []*models.Project

    for rows.Next() {
        var project models.Project

        err := rows.Scan(&project.ID, &project.Name, &project.Overview, &project.Description, &project.Repository)
        if err != nil {
            return nil, err
        }

        project.Keywords = s.getProjectKeywords(project.ID)

        projects = append(projects, &project)
    }

    return projects, nil
}

func (s *SQLiteProject) GetByID(id string) (*models.Project, error) {
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM projects WHERE id = ?;"
    row := s.db.QueryRow(query, id)

    var project models.Project

    err = row.Scan(&project.ID, &project.Name, &project.Overview, &project.Description, &project.Repository)

    if err != nil {
        return nil, err
    }

    project.Keywords = s.getProjectKeywords(project.ID)

    return &project, nil
}

func (s *SQLiteProject) GetByKeyword(keyword string) ([]*models.Project, error) {
    err := s.connect()
    if err != nil {
        return nil, err
    }

    defer s.disconnect()

    query := "SELECT * FROM projects JOIN keywords ON projects.id = keywords.project_id WHERE keywords.keyword = ?;"
    rows, err := s.db.Query(query, keyword)

    if err != nil {
        return nil, err
    }

    var projects []*models.Project

    for rows.Next() {
        var project models.Project
        var def string

        err := rows.Scan(&project.ID, &project.Name, &project.Overview, &project.Description, &project.Repository, &def, &def, &def)

        if err != nil {
            return nil, err
        }

        project.Keywords = s.getProjectKeywords(project.ID)

        projects = append(projects, &project)
    }

    return projects, nil
}

func (s *SQLiteProject) Update(project *models.Project) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    updateQuery := "UPDATE projects SET name = ?, overview = ?, description = ?, repository = ? WHERE id = ?;"
    _, err = s.db.Exec(updateQuery, project.Name, project.Overview, project.Description, project.Repository, project.ID)

    if err != nil {
        return err
    }

    deleteKeywordsQuery := "DELETE FROM keywords WHERE project_id = ?;"
    _, err = s.db.Exec(deleteKeywordsQuery, project.ID)

    if err != nil {
        return err
    }

    for _, keyword := range project.Keywords {
        insertKeywordQuery := "INSERT INTO keywords (project_id, keyword) VALUES (?, ?);"
        _, err := s.db.Exec(insertKeywordQuery, project.ID, keyword)

        if err != nil {
            return err
        }
    }

    return nil
}

func (s *SQLiteProject) Delete(id string) error {
    err := s.connect()
    if err != nil {
        return err
    }

    defer s.disconnect()

    deleteProjectQuery := "DELETE FROM projects WHERE id = ?;"
    _, err = s.db.Exec(deleteProjectQuery, id)

    return err
}

func (s *SQLiteProject) getProjectKeywords(projectID string) []string {
    err := s.connect()
    if err != nil {
        return nil
    }

    defer s.disconnect()
    query := "SELECT keyword FROM keywords WHERE project_id = ?;"
    rows, err := s.db.Query(query, projectID)

    if err != nil {
        return nil
    }

    var keywords []string

    for rows.Next() {
        var keyword string
        err := rows.Scan(&keyword)

        if err != nil {
            fmt.Println(err)
        }

        keywords = append(keywords, keyword)
    }

    return keywords
}

func (s *SQLiteProject) GetKeywords() []string {
    err := s.connect()
    if err != nil {
        return nil
    }

    defer s.disconnect()
    query := "SELECT DISTINCT keyword FROM keywords;"
    rows, err := s.db.Query(query)

    if err != nil {
        return []string{}
    }

    var keywords []string

    for rows.Next() {
        var keyword string
        rows.Scan(&keyword)

        keywords = append(keywords, keyword)
    }

    return keywords
}
