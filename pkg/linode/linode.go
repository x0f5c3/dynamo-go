package linode

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"os"
	"strings"

	_ "embed"

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

var Debug = pterm.PrintDebugMessages

type Client struct {
	stopped chan struct{}
	cancel  context.CancelFunc
	linodego.Client
	context.Context
}

func (c *Client) UpdateDomainRecord(domID, id int, ip string) error {
	doms, err := c.ListDomainRecords(c.Context, domID, nil)
	if err != nil {
		return err
	}
	var found *linodego.DomainRecord
	for _, v := range doms {
		if v.ID == id {
			found = &v
		}
	}
	if found != nil {
		updOpts := found.GetUpdateOptions()
		updOpts.Target = ip
		updOpts.TTLSec = 300
		updated, err := c.Client.UpdateDomainRecord(c.Context, domID, found.ID, updOpts)
		if err != nil {
			return err
		}
		if updated.ID != found.ID || updated.Target != updOpts.Target {
			return fmt.Errorf("failed to update")
		}
		return nil
	}
	return fmt.Errorf("can't find domain record in %d domain with ID %d", domID, id)
}

func (c *Client) FindDomains(toFind ...string) ([]linodego.Domain, error) {
	domains, err := c.ListDomains(nil)
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

func (c *Client) ListDomains(opts *linodego.ListOptions) ([]linodego.Domain, error) {
	return c.Client.ListDomains(c.Context, opts)
}

func NewClientFromEnv() (*Client, error) {
	tok, ok := envMap["LINODE_TOKEN"]
	if !ok {
		tok = os.Getenv("LINODE_TOKEN")
		if tok == "" {
			return nil, fmt.Errorf("no linode token found in env")
		}
	}
	return NewClientToken(tok)
}

func NewClientToken(token string) (*Client, error) {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	ctx, cancel := context.WithCancel(context.Background())
	cl := &http.Client{
		Transport: &oauth2.Transport{Source: src},
	}
	if Debug {
		trace := &httptrace.ClientTrace{
			GotConn: func(info httptrace.GotConnInfo) {
				pterm.Debug.Printfln("Got conn: %v", info)
			},
			WroteHeaderField: func(key string, value []string) {
				pterm.Debug.Printfln("Writing %s header with value %v", key, value)
			},
		}
		ctx = httptrace.WithClientTrace(ctx, trace)
	}
	linClient := linodego.NewClient(cl)
	linClient.SetDebug(Debug)
	return &Client{make(chan struct{}, 1), cancel, linClient, ctx}, nil
}
