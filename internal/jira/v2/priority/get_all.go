package priority

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Jira struct {
	User, Pass, FQDN string
}

func (j *Jira) GetAll() error {
	url := j.FQDN + "/rest/api/2/priority"
	log.Info(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(j.User, j.Pass)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Info(resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	log.Info(bodyString)
	return nil
}
