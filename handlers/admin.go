package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yannis94/kagami/helpers"
	"github.com/yannis94/kagami/middlewares"
	"github.com/yannis94/kagami/models"
	"github.com/yannis94/kagami/store"
)

var customClaims struct {
    UserID string
    jwt.MapClaims
}

type AdminHandler struct {
    db *store.Store
}

func NewAdminHandler(db *store.Store) *AdminHandler {
    return &AdminHandler{ db: db }
}

func (h AdminHandler) CreateAdmin(username, pwd string) error {
    if username == "" || pwd == "" {
        return errors.New("You can't add a user without username and/or password.")
    }

    hashPwd, err := helpers.HashPassword(pwd)

    if err != nil {
        return err
    }

    user := models.NewUser(username, hashPwd)

    err = h.db.User.Create(user)

    return err
}

func (h AdminHandler) HandlePostLogin(c echo.Context) error {
    user := &models.User{
        Username: c.FormValue("username"),
        Password: c.FormValue("password"),
    }

    userFound, err := h.db.User.GetByUsername(user.Username)

    if err != nil {
        return c.String(http.StatusInternalServerError, "Database error")
    }

    if err := helpers.IsPasswordCorrect(userFound.Password, user.Password); err != nil {
        time.Sleep(time.Second * 3)
        return c.String(http.StatusUnauthorized, "Unauthorized.")
    }

    tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "uuid": userFound.ID,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(time.Minute * 25).Unix(),
    })
    tknString, err := tkn.SignedString([]byte(middlewares.JWT_Secret))

    cookie := new(http.Cookie)
    cookie.Name = middlewares.AuthCookieName
    cookie.Value = tknString
    cookie.HttpOnly = true
    cookie.Secure = true
    cookie.Path = "/"
    cookie.Expires = time.Now().Add(time.Minute * 25)
    c.SetCookie(cookie)

    c.Response().Header().Add("HX-Redirect", "/yayadmin")
    return c.String(http.StatusAccepted, "Welcome")
}

func (h AdminHandler) HandlePostSkill(c echo.Context) error {
    skill := &models.Skill{
        ID: uuid.New().String(),
        Name: c.FormValue("name"),
        Category: c.FormValue("category"),
    }
    level, err := strconv.Atoi(c.FormValue("level"))

    if err != nil {
        return c.String(400, err.Error())
    }

    skill.Level = level
    logo, err := c.FormFile("logo")

    if err != nil {
        return c.String(500, err.Error())
    }


    src, err := logo.Open()

    if err != nil {
        return c.String(500, err.Error())
    }

    logoFilePath := path.Join(helpers.UploadImgPath, logo.Filename)
    dst, err := os.Create(logoFilePath)

    if err != nil {
        return c.String(500, err.Error())
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return c.String(500, err.Error())
    }

    skill.Logo = path.Join("/static/images/upload", logo.Filename)
    err = h.db.Skill.Create(skill)

    if err != nil {
        return c.String(400, err.Error())
    }

    c.Response().Header().Add("HX-Refresh", "true")
    return c.String(200, "Skill added !")
}

func (h AdminHandler) HandlePostProject(c echo.Context) error {
    project := &models.Project{
        ID: uuid.New().String(),
        Name: c.FormValue("name"),
        Overview: c.FormValue("overview"),
        Description: c.FormValue("description"),
        Repository: c.FormValue("repository"),
        Keywords: strings.Split(c.FormValue("keywords"), ","),
    }

    err := h.db.Project.Create(project)

    if err != nil {
        return c.String(500, err.Error())
    }

    c.Response().Header().Add("HX-Refresh", "true")
    return c.String(200, "Project added !")
}

func (h AdminHandler) HandlePutProject(c echo.Context) error {
    project := &models.Project{
        ID: c.Param("id"),
        Name: c.FormValue("name"),
        Overview: c.FormValue("overview"),
        Description: c.FormValue("description"),
        Repository: c.FormValue("repository"),
        Keywords: strings.Split(c.FormValue("keywords"), ","),
    }
    err := h.db.Project.Update(project)

    if err != nil {
        return c.String(500, err.Error())
    }

    c.Response().Header().Add("HX-Refresh", "true")
    return c.String(200, "Project updated !")
}

func (h AdminHandler) HandlePutSkill(c echo.Context) error {
    skill := &models.Skill{
        ID: c.Param("id"),
        Name: c.FormValue("name"),
        Category: c.FormValue("category"),
    }
    level, err := strconv.Atoi(c.FormValue("level"))

    if err != nil {
        return c.String(400, err.Error())
    }

    skill.Level = level
    logo, err := c.FormFile("logo")

    if err != nil {
        if err != http.ErrMissingFile {
            fmt.Println(err)
            return c.String(500, err.Error())
        } 
    } else {
        src, err := logo.Open()
        
        if err != nil {
            return c.String(500, err.Error())
        }

        tempSkill, err := h.db.Skill.GetByID(skill.ID)

        if err != nil {
            return c.String(500, err.Error())
        }
        if tempSkill == nil {
            return c.String(400, "Bad request")
        }

        oldLogoPath := path.Join(helpers.UploadImgPath, path.Base(tempSkill.Logo))
        helpers.DeleteFile(oldLogoPath)
        logoFilePath := path.Join(helpers.UploadImgPath, logo.Filename)
        dst, err := os.Create(logoFilePath)

        if err != nil {
            return c.String(500, err.Error())
        }
        defer dst.Close()

        if _, err := io.Copy(dst, src); err != nil {
            return c.String(500, err.Error())
        }

        skill.Logo = path.Join("/static/images/upload", logo.Filename)
    }

    err = h.db.Skill.Update(skill)

    if err != nil {
        return c.String(400, err.Error())
    }

    c.Response().Header().Add("HX-Refresh", "true")
    return c.String(200, "Skill added !")
}

func (h AdminHandler) HandleDeleteSkill(c echo.Context) error {
    skillID := c.Param("id")

    err := h.db.Skill.Delete(skillID)

    if err != nil {
        return c.String(500, err.Error())
    }

    c.Response().Header().Add("HX-Refresh", "true")
    return c.String(200, "Skill deleted !")
}

func (h AdminHandler) HandleDeleteProject(c echo.Context) error {
    projectID := c.Param("id")

    err := h.db.Project.Delete(projectID)

    if err != nil {
        return c.String(500, err.Error())
    }

    c.Response().Header().Add("HX-Refresh", "true")
    return c.String(200, "Skill deleted !")
}
