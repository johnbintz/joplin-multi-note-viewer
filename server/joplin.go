package main

import (
	"encoding/json"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/url"
)

type joplinConfig struct {
	APIKey string `yaml:"api_key"`
	URL    string `yaml:"url"`
}

type config struct {
	JoplinConfig joplinConfig `yaml:"joplin"`
}

func (c *config) BaseURL() string {
	return c.JoplinConfig.URL
}

func (c *config) BaseParams() url.Values {
	values := url.Values{}
	values.Add("token", c.JoplinConfig.APIKey)

	return values
}

var appConfig = &config{}

func readConfig() error {
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, &appConfig)
	if err != nil {
		return err
	}

	return nil
}

type searchResponseItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type searchResponse struct {
	Items   []searchResponseItem `json:"items"`
	HasMore bool                 `json:"has_more"`
}

func (s *searchResponse) GetFrontendBody() ([]byte, error) {
	res, err := json.Marshal(s.Items)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type noteResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (s *noteResponse) GetFrontendBody() ([]byte, error) {
	res, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func performSearch(query string) (searchResponse, error) {
	params := appConfig.BaseParams()
	params.Add("query", query+"*")
	params.Add("type", "note")

	response := searchResponse{}

	url := appConfig.BaseURL() + "/search?" + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

func retrieveNote(ID string) (noteResponse, error) {
	params := appConfig.BaseParams()
	params.Add("fields", "title,body,id")

	response := noteResponse{}

	url := appConfig.BaseURL() + "/notes/" + ID + "?" + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}
