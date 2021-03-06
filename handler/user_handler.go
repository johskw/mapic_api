package handler

import (
	"net/http"
	"strconv"

	"github.com/johskw/mapic_api/domain"
	"github.com/labstack/echo"
)

func GetUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := domain.GetUser(id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &user)
}

func PostUser(c echo.Context) (err error) {
	user := new(domain.User)
	err = c.Bind(user)
	if err != nil {
		return
	}
	createdUser, err := user.Create()
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &createdUser)
}

func PutUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := domain.GetUser(id)
	if err != nil {
		return
	}
	newUser := new(domain.User)
	err = c.Bind(newUser)
	if err != nil {
		return
	}
	err = user.Update(*newUser)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &newUser)
}

func DeleteUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := domain.GetUser(id)
	if err != nil {
		return
	}
	err = user.Delete()
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &user)
}
