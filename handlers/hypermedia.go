package handlers

import (
	"net/http"

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

func (h Hypermedia) HandleGetRelatedProjects(c echo.Context) error {
    currentProject, err := h.db.Project.GetByID(c.Param("id"))

    if err != nil || currentProject == nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    var results []*models.Project

    for _, keyword := range currentProject.Keywords {
        projects, err := h.db.Project.GetByKeyword(keyword)

        if err != nil {
            return c.String(http.StatusInternalServerError, err.Error())
        }

        results = append(results, projects...)
    }

    var matchs map[string]int = make(map[string]int)

    for _, result := range results {
        matchs[result.ID] += 1
    }

    orderedID := bubbleSort(matchs)

    var relatedProjects []*models.Project

    for _, id := range orderedID {
        if id == c.Param("id") {
            continue
        }
        for _, project := range results {
            if project.ID == id {
                relatedProjects = append(relatedProjects, project)
                break
            }
        }
    }

    return render(c, cards.RelatedProjects(relatedProjects))
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

func bubbleSort(matchs map[string]int) []string {
    n := len(matchs)
    res := make([]string, n)

    for key, _ := range matchs {
        res = append(res, key)
    }

    for i:=0; i<n-1; i++ {
        for j:=0; j<n-i-1; j++ {
            if res[j] < res[j+1] {
                res[j], res[j+1] = res[j+1], res[j]
            }
        }
    }

    return res
}
