package middlewares

import (
	"fmt"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/yannis94/kagami/store"
)

var (
    AuthCookieName string = "kagami-access-token"
    JWT_Secret string = os.Getenv("JWT_SECRET")
)

var AuthMiddlewareConfig echojwt.Config = echojwt.Config{
    SigningKey: []byte(JWT_Secret),
    BeforeFunc: func(c echo.Context) {
        cookie, err := c.Cookie(AuthCookieName)

        if err != nil {
            fmt.Println("Cookie error:", err)
            return
        }

        fmt.Println(cookie.Value)

        c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie.Value))
    },
    ErrorHandler: func(c echo.Context, err error) error {
        fmt.Println("Middleware error:", err)
        c.Response().Header().Add("HX-Redirect", "/yayadmin/login")
        return err
    },
}

type AuthMiddleware struct {
    db *store.Store
}

func NewAuthMiddleware(db *store.Store) *AuthMiddleware {
    return &AuthMiddleware{ db: db }
}

