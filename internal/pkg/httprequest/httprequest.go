package httprequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Jira struct {
	Data                                      interface{}
	APIVersion, Method, FQDN, User, Pass, URI string
}

func (j *Jira) Construct() (*http.Request, error) {
	b, err := json.Marshal(j.Data)
	if err != nil {
		return nil, err
	}
	log.Info("---------------")
	log.Info(string(b))
	log.Info("---------------")
	body := bytes.NewReader(b)

	req, err := http.NewRequest(j.Method, j.FQDN+"/rest/api/"+j.APIVersion+"/issue"+j.URI, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (j *Jira) Initiate(req *http.Request) (string, error) {
	req.SetBasicAuth(j.User, j.Pass)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ticketNumber = ""
	if resp.StatusCode == http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		log.Info(bodyString)

		value := gjson.Get(bodyString, "key")
		ticketNumber = value.String()
	}

	return ticketNumber, nil
}

func (j *Jira) ConstructAndInitiate() (string, error) {
	h, err := j.Construct()
	if err != nil {
		return "", err
	}

	ticketNumber, err := j.Initiate(h)
	if err != nil {
		return "", err
	}
	return ticketNumber, nil
}
