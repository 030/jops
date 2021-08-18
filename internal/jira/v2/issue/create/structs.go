package create

type Jira struct {
	User, Pass, FQDN, Project, Summary, Description string
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

type Fields struct {
	Description string    `json:"description"`
	Issuetype   Issuetype `json:"issuetype"`
	Project     Project   `json:"project"`
	Summary     string    `json:"summary"`
}
