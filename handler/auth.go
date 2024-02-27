package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
}

type User struct {
	Login    string `json:"login" form:"login" query:"login"`
	Password string `json:"password" form:"password" query:"password"`
}

type UserDTO struct {
	Login    string
	Password string
}

type TokenDTO struct {
	Token string `json:"token"`
}

func (h AuthHandler) CookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}

	if cookie.Value != "123" {
		return c.String(http.StatusForbidden, "not autorized")
	}

	return c.String(http.StatusOK, "read a cookie")
}

func checkUser(user UserDTO) (token TokenDTO, error error) {
	if user.Login == "admin" && user.Password == "123" {
		token := TokenDTO{
			Token: "123",
		}
		return token, nil
	}

	return TokenDTO{}, errors.New("wrong password")

}

func (h AuthHandler) LoginHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	// Load into separate struct for security
	user := UserDTO{
		Login:    u.Login,
		Password: u.Password,
	}

	token, error := checkUser(user)
	if error != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "The login/password is wrong"})
	}

	return c.JSON(http.StatusOK, token)
}
