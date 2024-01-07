package store

type Store struct {
    skill SkillStorage
    project ProjectStorage
}

func NewStorage(skillDB SkillStorage, projectDB ProjectStorage) *Store {
    return &Store{ skill: skillDB, project: projectDB }
}
