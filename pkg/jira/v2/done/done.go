package done

import (
	"github.com/030/jops/internal/pkg/httprequest"
	log "github.com/sirupsen/logrus"
)

func (j *Jira) Done() error {
	log.Info(j.TicketNumber)
	add := Add{Body: j.Comment}
	comment := []Comment{{add}}
	update := Update{Comment: comment}
	fields := Fields{Resolution: Resolution{"Done"}}
	data := Payload{
		Fields:     fields,
		Transition: Transition{"21"},
		Update:     update,
	}

	uri := "/" + j.Project + "-" + j.TicketNumber + "/transitions"
	htj := httprequest.Jira{URI: uri, APIVersion: "2", Data: data, Method: "POST", FQDN: j.FQDN, User: j.User, Pass: j.Pass}
	if err := htj.ConstructAndInitiate(); err != nil {
		return err
	}

	return nil
}
