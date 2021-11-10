package changelog

import (
	"github.com/030/jops/internal/pkg/httprequest"
	log "github.com/sirupsen/logrus"
)

func (j *Jira) Get() error {
	url := j.FQDN + "/rest/api/2/issue/" + j.TicketNumber + "?expand=changelog"
	p := httprequest.Params{Pass: j.Pass, User: j.User, URL: url}
	s, err := p.Action()
	if err != nil {
		return err
	}
	log.Info(s)
	return nil
}
