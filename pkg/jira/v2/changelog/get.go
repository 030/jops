package changelog

import (
	"github.com/030/jops/internal/pkg/httprequest"
)

func (j *Jira) Get() (string, error) {
	url := j.FQDN + "/rest/api/2/issue/" + j.TicketNumber + "?expand=changelog"
	p := httprequest.Params{Pass: j.Pass, User: j.User, URL: url}
	s, err := p.Action()
	if err != nil {
		return "", err
	}
	return s, nil
}
