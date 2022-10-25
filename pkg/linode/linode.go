package linode

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	_ "embed"

	"github.com/awnumar/memguard"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/linode/linodego"
	"github.com/pterm/pterm"
	"golang.org/x/oauth2"
)

//go:embed .env
var dotenv string

var envMap = func() map[string]string {
	v, err := godotenv.Unmarshal(dotenv)
	if err != nil {
		pterm.Fatal.Printfln("Failed to unmarshal the .env file: %s", err)
	}
	return v
}()

var Debug bool

type Client struct {
	token   *memguard.Enclave
	stopped chan struct{}
	cancel  context.CancelFunc
	linodego.Client
	context.Context
}

func (c *Client) FindDomains(toFind ...string) ([]linodego.Domain, error) {
	domains, err := c.ListDomains(c.Context, nil)
	if err != nil {
		return nil, err
	}
	var res []linodego.Domain

	if len(toFind) > 0 {
	OuterLoop:
		for _, v := range domains {
			for _, x := range toFind {
				if strings.Contains(v.Domain, x) {
					res = append(res, v)
					continue OuterLoop
				}
			}
		}
	}
	return domains, nil
}

func (c *Client) UpdateDNSRecord(domainID int, recordID int, record linodego.DomainRecordUpdateOptions) error {
	_, err := c.UpdateDomainRecord(c.Context, domainID, recordID, record)
	return err
}

func (c *Client) ListDomains(ctx context.Context, opts *linodego.ListOptions) ([]linodego.Domain, error) {
	return c.Client.ListDomains(ctx, opts)
}

func NewClientFromEnv() (*Client, error) {
	tok, ok := envMap["LINODE_TOKEN"]
	if !ok {
		return nil, fmt.Errorf("no linode token found in env")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tok})
	cl := &http.Client{
		Transport: &oauth2.Transport{Source: tokenSource},
	}
	linClient := linodego.NewClient(cl)
	linClient.SetDebug(Debug)
	ctx, cancel := context.WithCancel(context.Background())
	tokEnc := memguard.NewEnclave([]byte(tok))
	return &Client{tokEnc, make(chan struct{}, 1), cancel, linClient, ctx}, nil
}

func NewClientToken(token string) (*Client, error) {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	ctx, cancel := context.WithCancel(context.Background())
	cl := &http.Client{
		Transport: &oauth2.Transport{Source: src},
	}
	linClient := linodego.NewClient(cl)
	linClient.SetDebug(Debug)
	tokEnc := memguard.NewEnclave([]byte(token))
	return &Client{tokEnc, make(chan struct{}, 1), cancel, linClient, ctx}, nil
}
