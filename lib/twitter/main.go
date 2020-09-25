package twitter

import (
	"fmt"
	"log"

	"github.com/ChimeraCoder/anaconda"

	"github.com/garyburd/go-oauth/oauth"

	"gopkg.in/ini.v1"
)

func GetConnect() *oauth.Client {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  cfg.Section("twitter").Key("consumer_token").String(),
			Secret: cfg.Section("twitter").Key("consumer_secret").String(),
		},
	}
}

func GetAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := GetConnect()
	at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

	return at, err
}

func GetSelfData(token, secret string) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to Get SelfData: %v", err)
	}
	anaconda.SetConsumerKey(cfg.Section("twitter").Key("consumer_token").String())
	anaconda.SetConsumerSecret(cfg.Section("twitter").Key("consumer_secret").String())
	api := anaconda.NewTwitterApi(token, secret)
	data, _ := api.GetSelf(nil)
	fmt.Printf("%v is authenticated", data)
}
