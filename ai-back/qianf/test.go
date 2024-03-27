package qianf

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
)

func Token(apiKey string, apiSecretKey string) (TokenService, error) {

	url := BASEURL + TOKENURL + "?grant_type=client_credentials&client_id=" + apiKey + "&client_secret=" + apiSecretKey
	payload := strings.NewReader(``)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Error().Str("location", "qianf/struct.go").Msg(err.Error())
		return TokenService{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Error().Str("location", "qianf/struct.go").Msg(err.Error())
		return TokenService{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return TokenService{}, err
	}

	var token TokenService
	err = json.Unmarshal(body, &token)

	return token, nil
}
