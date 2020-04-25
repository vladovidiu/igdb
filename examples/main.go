package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vladovidiu/igdb"
)

type endpoint string

const (
	rootURL string = "https://api-v3.igdb.com/"
)

func main() {
	games, err := getGames("games?fields=name")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, game := range games {
		fmt.Println(game.Name)
	}
}

func getGames(end endpoint) ([]igdb.Game, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", rootURL+string(end), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", "51c65c62a2e7420129ba4e521a576ab8")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var g []igdb.Game

	err = json.Unmarshal(body, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
