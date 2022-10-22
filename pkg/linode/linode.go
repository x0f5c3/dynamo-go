package linode

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

var Debug = false

type Client struct {
	token string
	linodego.Client
}

func NewClientFromEnv() (*Client, error) {
	tok, ok := os.LookupEnv("LINODE_TOKEN")
	if !ok {
		return nil, fmt.Errorf("no linode token found in env")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tok})
	cl := &http.Client{
		Transport: &oauth2.Transport{Source: tokenSource},
	}
	linClient := linodego.NewClient(cl)
	if Debug {
		linClient.SetDebug(Debug)
	}
	return &Client{tok, linClient}, nil
}
