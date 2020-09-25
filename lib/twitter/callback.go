package twitter

import (
	"log"
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

func Callback(c echo.Context) error {
	session := session.Default(c)
	secret := c.QueryParam("oauth_verifier")
	at, err := GetAccessToken(
		&oauth.Credentials{
			Token:  session.Get("request_token").(string),
			Secret: session.Get("request_token_secret").(string),
		},
		secret,
	)
	if err != nil {
		log.Printf("Failed to Get AccessToken: %v", err)
	}

	session.Set("oauth_token", at.Token)
	session.Set("oauth_secret", at.Secret)

	GetSelfData(at.Token, at.Secret)

	session.Save()

	return c.Redirect(http.StatusFound, "/post")
}
