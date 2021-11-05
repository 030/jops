package done

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/030/jops/internal/pkg/httprequest"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func (j *Jira) transitionDoneID() (string, error) {
	url := j.FQDN + "/rest/api/2/issue/" + j.TicketNumber + "/transitions"
	log.Info(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(j.User, j.Pass)
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
	log.Info(bodyString)

	value := gjson.Get(bodyString, "transitions.#(name==\"Done\").id")
	id := value.String()
	if id == "" {
		return "", fmt.Errorf("transition id Done should not be empty")
	}

	return id, nil
}

func (j *Jira) Done() error {
	tdi, err := j.transitionDoneID()
	if err != nil {
		return err
	}

	log.Infof("ticketnumber: '%s' and transitionID: '%s'", j.TicketNumber, tdi)
	add := Add{Body: j.Comment}
	comment := []Comment{{add}}
	update := Update{Comment: comment}
	fields := Fields{Resolution: Resolution{"Done"}}
	data := Payload{
		Fields:     fields,
		Transition: Transition{tdi},
		Update:     update,
	}

	uri := "/" + j.TicketNumber + "/transitions"
	htj := httprequest.Jira{URI: uri, APIVersion: "2", Data: data, Method: "POST", FQDN: j.FQDN, User: j.User, Pass: j.Pass}
	_, err = htj.ConstructAndInitiate()
	if err != nil {
		return err
	}

	return nil
}
