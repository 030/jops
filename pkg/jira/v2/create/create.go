package create

import (
	"fmt"
	"strings"

	"github.com/030/jops/internal/pkg/httprequest"
	log "github.com/sirupsen/logrus"
)

func (j *Jira) Create() (string, error) {
	multiLineIssueFixInDescription := strings.ReplaceAll(j.Description, "\\n", "\n")
	fields := Fields{
		Description: multiLineIssueFixInDescription,
		Summary:     j.Summary,
		Issuetype:   Issuetype{"Story"},
		Labels:      j.Labels,
		Project:     Project{j.Project},
		Priority:    Priority{j.Priority},
	}
	data := Payload{
		Fields: fields,
	}

	htj := httprequest.Jira{APIVersion: "2", Data: data, Method: "POST", FQDN: j.FQDN, User: j.User, Pass: j.Pass}
	ticketNumber, err := htj.ConstructAndInitiate()
	if err != nil {
		return "", err
	}
	if ticketNumber == "" {
		return "", fmt.Errorf("ticketNumber should not be empty")
	}
	log.Infof("Ticket: '%s' has been created", ticketNumber)

	return ticketNumber, nil
}
