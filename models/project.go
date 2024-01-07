package models

import "github.com/google/uuid"

type Project struct {
    ID string
    Name string
    Overview string
    Description string
    Keywords []string
}

func CreateProject(keywords []string, name, overview, description string) *Project {
    return &Project{
        ID: uuid.New().String(),
        Name: name,
        Overview: overview,
        Description: description,
        Keywords: keywords,
    }
}
