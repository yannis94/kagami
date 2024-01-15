package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/yannis94/kagami/models"
	"github.com/yannis94/kagami/store"
	"github.com/yannis94/kagami/views/components/cards"
	"github.com/yannis94/kagami/views/components/forms"
	"github.com/yannis94/kagami/views/components/popup"
)

type Hypermedia struct {
    db *store.Store
}

func NewHypermediaHandler(db *store.Store) *Hypermedia {
    return &Hypermedia{ db: db }
}

func (h Hypermedia) HandleGetContact(c echo.Context) error {
    return render(c, popup.Contact())
}

func (h Hypermedia) HandleGetSkillCategories(c echo.Context) error {
    categories, err := h.db.Skill.GetCategories()

    if err != nil {
        return c.String(500, err.Error())
    }

    return render(c, forms.SkillCategories(categories))
}

func (h Hypermedia) HandleGetSkills(c echo.Context) error {
    var (
        skills []*models.Skill
        err error
        category string = c.FormValue("category")
    )

    if category == "" {
        skills, err = h.db.Skill.GetAll()
    } else {
        skills, err = h.db.Skill.GetByCategory(category)
    }

    if err != nil {
        return c.String(500, err.Error())
    }

    return render(c, cards.Skill(skills))
}

func (h Hypermedia) HandleGetProjects(c echo.Context) error {
    var (
        projects []*models.Project
        err error
    )

    if c.FormValue("keyword") != "" {
        projects, err = h.db.Project.GetByKeyword(c.FormValue("keyword"))
    } else {
        projects, err = h.db.Project.GetAll()
    }

    if err != nil {
        return c.String(500, err.Error())
    }

    return render(c, cards.Project(projects))
}

func (h Hypermedia) HandleGetProjectKeywords(c echo.Context) error {
    keywords := h.db.Project.GetKeywords()

    return render(c, forms.ProjectKeywords(keywords))
}

func (h Hypermedia) HandleGetAdminSkills(c echo.Context) error {
    var (
        skills []*models.Skill
        err error
    )

    skills, err = h.db.Skill.GetAll()

    if err != nil {
        return c.String(500, err.Error())
    }

    return render(c, cards.AdminSkill(skills))
}

func (h Hypermedia) HandleGetAdminProjects(c echo.Context) error {
    projects, err := h.db.Project.GetAll()

    if err != nil {
        return c.String(500, err.Error())
    }

    return render(c, cards.AdminProject(projects))
}
