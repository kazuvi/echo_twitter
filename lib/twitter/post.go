package twitter

import (
	"log"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/skratchdot/open-golang/open"
)

func PostTweet(c echo.Context) error {
	session := session.Default(c)
	token := session.Get("oauth_token").(string)
	secret := session.Get("oauth_secret").(string)

	api := anaconda.NewTwitterApi(token, secret)

	message := "Test"
	tweet, err := api.PostTweet(message, nil)
	if err != nil {
		log.Printf("Failed to Post Tweet: %v", err)
		return c.JSON(http.StatusAccepted, "redirect")
	}
	link := "https://twitter.com/" + tweet.User.IdStr + "/status/" + tweet.IdStr

	open.Start(link)

	return c.JSON(http.StatusOK, link)

}
