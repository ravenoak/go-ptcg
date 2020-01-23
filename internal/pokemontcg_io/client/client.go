package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const defaultAPIURL string = "https://api.pokemontcg.io/"

type (
	Client struct {
		config *Config
	}

	Set struct {
		Name          string `json:"name"`
		PTCGOCode     string `json:"ptcgoCode"`
		Series        string `json:"series"`
		TotalCards    int    `json:"totalCards"`
		StandardLegal bool   `json:"standardLegal"`
		ExpandedLegal bool   `json:"expandedLegal"`
		Code          string
		ReleaseDate   string
		SymbolUrl     url.URL
		LogoUrl       url.URL
		UpdatedAt     time.Time
	}
	SetRequest struct {
		UpdatedSince string
		Page         string
		PageSize     int
		Set
	}
	SetsResponse struct {
		Sets []Set `json:"sets"`
	}
	SetResponse struct {
		Set Set `json:"set"`
	}
)

func NewClient(config *Config) *Client {
	if config != nil {
		return &Client{config: config}
	} else {
		return &Client{config: &DefaultConfig}
	}
}

func (c *Client) Sets() []Set {
	v := 1

	resp, err := http.Get(fmt.Sprintf("%sv%d/sets", c.config.apiURL, v))
	if err != nil {
		log.Fatalln(err)
	}

	var result SetsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)

	return result.Sets
}
