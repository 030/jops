package jira

import (
	"fmt"

	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

const (
	apiRestV2ProjectSearch = "/rest/api/2/search?jql=project%20%3D%20"
	and                    = "%20AND%20"
)

type Server struct {
	Credentials   []string
	FQDN, Project string
}

func (s *Server) JQL(query string) ([]byte, error) {
	url := s.FQDN + apiRestV2ProjectSearch + s.Project + and + query
	log.Debug(url)

	resp, err := grequests.Get(url, &grequests.RequestOptions{Auth: s.Credentials})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code not 200. Verify the url: '%s'", url)
	}

	return resp.Bytes(), nil
}

func ParseJSON(json []byte) {
	issues := gjson.GetBytes(json, "issues")
	for _, issue := range issues.Array() {
		log.Info(gjson.Get(issue.String(), "key").String())
		log.Info(gjson.Get(issue.String(), "fields.summary").String())
		log.Info(gjson.Get(issue.String(), "fields.description").String())
	}
}
