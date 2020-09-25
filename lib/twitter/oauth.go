package twitter

import (
	"log"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

func Oauth(c echo.Context) error {
	config := GetConnect()
	rt, err := config.RequestTemporaryCredentials(nil, "http://localhost:8080/callback", nil)
	if err != nil {
		log.Printf("Failed to Connect: %v", err)
	}

	session := session.Default(c)
	session.Set("request_token", rt.Token)
	session.Set("request_token_secret", rt.Secret)
	session.Save()

	url := config.AuthorizationURL(rt, nil)

	return c.Redirect(http.StatusFound, url)
}
