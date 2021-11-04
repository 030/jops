package create

import "github.com/030/jops/internal/pkg/httprequest"

func (j *Jira) Create() error {
	fields := Fields{
		Description: j.Description,
		Summary:     j.Summary,
		Issuetype:   Issuetype{"Story"},
		Project:     Project{j.Project},
		Priority:    Priority{j.Priority},
	}
	data := Payload{
		Fields: fields,
	}

	htj := httprequest.Jira{APIVersion: "2", Data: data, Method: "POST", FQDN: j.FQDN, User: j.User, Pass: j.Pass}
	if err := htj.ConstructAndInitiate(); err != nil {
		return err
	}

	return nil
}
