package models

import "github.com/google/uuid"

type Project struct {
    ID string
    Name string
    Overview string
    Description string
    Repository string
    Keywords []string
}

func CreateProject(keywords []string, name, overview, description, repository string) *Project {
    return &Project{
        ID: uuid.New().String(),
        Name: name,
        Overview: overview,
        Description: description,
        Keywords: keywords,
        Repository: repository,
    }
}
