package store

import "github.com/yannis94/kagami/models"

type ProjectStorage interface {
    Create(*models.Project) error
    Read(filter *models.Project) ([]*models.Project, error)
    Update(*models.Project) error
    Delete(filter *models.Project) error
}
