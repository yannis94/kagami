package handlers

import (
	"net/http"

	"github.com/yannis94/kagami/store"
	"github.com/yannis94/kagami/views/pages"

	"github.com/labstack/echo/v4"
)

type Pages struct {
    db *store.Store
}

func NewPageHandler(db *store.Store) *Pages {
    return &Pages{ db: db }
}

func (h Pages) HandleGetIndex(c echo.Context) error {
    return render(c, pages.Index())
}

func (h Pages) HandleGetProjects(c echo.Context) error {
    return render(c, pages.Projects())
}

func (h Pages) HandleGetProject(c echo.Context) error {
    project, err := h.db.Project.GetByID(c.Param("id"))

    if project == nil {
        return render(c, pages.Error(http.StatusNotFound, "Project not found."))
    }
    if err != nil {
        return render(c, pages.Error(http.StatusInternalServerError, "Internal server error."))
    }
    return render(c, pages.Project(project))
}

func (h Pages) HandleGetAdminHome(c echo.Context) error {
    return render(c, pages.AdminHome())
}

func (h Pages) HandleGetAdminLogin(c echo.Context) error {
    return render(c, pages.AdminLogin())
}

func (h Pages) HandleGetAdminSkill(c echo.Context) error {
    skill, err := h.db.Skill.GetByID(c.Param("id"))

    if skill == nil {
        return render(c, pages.Error(http.StatusNotFound, "Project not found."))
    }
    if err != nil {
        return render(c, pages.Error(http.StatusInternalServerError, "Internal server error."))
    }

    return render(c, pages.AdminSkill(skill))
}

func (h Pages) HandleGetAdminProject(c echo.Context) error {
    project, err := h.db.Project.GetByID(c.Param("id"))

    if project == nil {
        return render(c, pages.Error(http.StatusNotFound, "Project not found."))
    }
    if err != nil {
        return render(c, pages.Error(http.StatusInternalServerError, "Internal server error."))
    }

    return render(c, pages.AdminProject(project))
}
