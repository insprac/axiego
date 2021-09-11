package axiego

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://axieinfinity.com/graphql-server-v2/graphql"

func Req(query string, vars map[string]interface{}, v interface{}) error {
	if vars == nil {
		vars = map[string]interface{}{}
	}

	postBody, err := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": vars,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))

	return json.Unmarshal(body, &v)
}
