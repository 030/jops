package comment

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// curl --request GET \
//   --url 'https://your-domain.atlassian.net/rest/api/2/issue/{issueIdOrKey}/comment' \
//   --user 'email@example.com:<api_token>' \
//   --header 'Accept: application/json'

// curl --request GET \
//   --url 'https://your-domain.atlassian.net/rest/api/2/issue/{issueIdOrKey}/comment/{id}' \
//   --user 'email@example.com:<api_token>' \
//   --header 'Accept: application/json'

// type Jira struct {
// 	Data                                          interface{}
// 	APIVersion, HTTPMethod, FQDN, User, Pass, URI string
// }

func (j *Jira) GetAll() error {
	// req, err := http.NewRequest("GET", "https://your-domain.atlassian.net/rest/api/2/issue/{issueIdOrKey}/comment/{id}", nil)
	url := j.FQDN + "/rest/api/2/issue/" + j.TicketNumber + "/comment"
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
