package httprequest

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Params struct {
	Pass, URL, User string
}

func (p *Params) Action() (string, error) {
	log.Info(p.URL)
	req, err := http.NewRequest("GET", p.URL, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(p.User, p.Pass)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	log.Info(resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}
