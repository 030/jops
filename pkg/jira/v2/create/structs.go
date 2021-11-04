package create

type Jira struct {
	User, Pass, FQDN, Priority, Project, Summary, Description string
	Labels                                                    []string
}

type Payload struct {
	Fields Fields `json:"fields"`
}
type Issuetype struct {
	ID string `json:"name"`
}
type Project struct {
	ID string `json:"key"`
}
type Priority struct {
	ID string `json:"name"`
}

type Fields struct {
	Description string    `json:"description"`
	Issuetype   Issuetype `json:"issuetype"`
	Labels      []string  `json:"labels"`
	Priority    Priority  `json:"priority"`
	Project     Project   `json:"project"`
	Summary     string    `json:"summary"`
}
