package http

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Loginhandler struct{}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *Loginhandler) Login(echoCtx echo.Context) error {
	username := echoCtx.FormValue("username")
	password := echoCtx.FormValue("password")
	// Check in your db if the user exists or not
	if username == "jon" && password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256) // Set claims
		// This is the information which frontend can use
		// The backend can also decode the token and get admin etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Doe"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Minute * 45).Unix()
		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID          works too)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return echoCtx.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

func (h *Loginhandler) Private(echoCtx echo.Context) error {
	user := echoCtx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return echoCtx.String(http.StatusOK, "Welcome "+name+"!")
}
