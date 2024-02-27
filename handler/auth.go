package handler

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
}

type Authentification struct {
	Session TokenDTO
}

func (a *Authentification) saveSession(tokenDTO TokenDTO) {
	fmt.Printf("Saved token %s", tokenDTO.Token)
	a.Session = tokenDTO
}

func (a *Authentification) getSession() TokenDTO {
	fmt.Printf("Read token %s", a.Session.Token)
	return a.Session
}

var Auth = Authentification{Session: TokenDTO{}}

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
		sessionId, err := generateSessionID(16)
		if err != nil {
			return TokenDTO{}, err
		}

		token := TokenDTO{
			Token: sessionId,
		}

		Auth.saveSession(token)
		return token, err
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

func generateSessionID(length int) (string, error) {
	// Calculate the number of bytes needed for the given length
	numBytes := length / 4 * 3
	if length%4 != 0 {
		numBytes++
	}

	// Generate random bytes
	bytes := make([]byte, numBytes)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	sessionID := base64.URLEncoding.EncodeToString(bytes)

	// Truncate or pad the session ID to the desired length
	sessionID = sessionID[:length]

	return sessionID, nil
}
