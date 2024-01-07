package models

import "github.com/google/uuid"

type Skill struct {
    ID string
    Name string
    Category string
    Logo string
    Level int
}

func CreateSkill(name string, category string, logoPath string, level int) *Skill {
    return &Skill{
        ID: uuid.New().String(),
        Name: name,
        Category: category,
        Logo: logoPath,
        Level: level,
    }
}
