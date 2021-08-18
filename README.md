# JOps

Jira Operations

URL encoding, see https://www.w3schools.com/tags/ref_urlencode.ASP

```bash
= : %3D
space : %20
' : %27
```

## Example

```bash
package main

import (
  "github.com/030/jops/internal/jira"
  log "github.com/sirupsen/logrus"
)

func main() {
  creds := []string{"user", "pass"}
  j := jira.Server{Credentials: creds, FQDN: "fqdn", Project: "project-name"}
  b, err := j.JQL("jql")
  if err != nil {
    log.Fatal(err)
  }

  jira.ParseJSON(b)
}
```

## Sources

* https://developer.atlassian.com/cloud/jira/platform/rest/v2
