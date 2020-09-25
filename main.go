package main

import (
	"echo_twitter/lib/twitter"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

func main() {
	var e = echo.New()

	store := session.NewCookieStore([]byte("secret-key"))

	store.MaxAge(86400)
	e.Use(session.Sessions("ESESSION", store))

	e.GET("/oauth", twitter.Oauth)
	e.GET("/callback", twitter.Callback)
	e.GET("/post", twitter.PostTweet)

	e.Logger.Fatal(e.Start(":8080"))
}
