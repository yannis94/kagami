package store

import "github.com/yannis94/kagami/models"

type SkillStorage interface {
    Create(*models.Skill) error
    Read(filter *models.Skill) ([]*models.Skill, error)
    Update(*models.Skill) error
    Delete(filter *models.Skill) error
}
