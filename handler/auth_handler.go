package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/johskw/mapic_api/domain"
	"github.com/labstack/echo"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) (err error) {
	auth := new(Auth)
	err = c.Bind(auth)
	user, err := domain.GetUserByEmail(auth.Email)
	if err != nil {
		return
	}
	if auth.Password != user.Password {
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 876000).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
		"user":  user,
	})
}

func Restricted(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	user, err := domain.GetUser(id)
	return c.String(http.StatusOK, "Welcome "+user.Name+"!")
}
