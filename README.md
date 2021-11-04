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

## Close

```bash
./jops done -t some-ticket -c test-ticket
```

## Create

### Prio 4

```bash
./jops create -d hello -s world -p "Prio 4"
```

## Comment

### Get all Ticket comments

```bash
./jops comment --ticketNumber=some-ticket --all
```

### Add a comment to an existing ticket

```bash
./jops comment --ticketNumber=some-ticket --add --message="hello world"
```

## Sources

* <https://developer.atlassian.com/cloud/jira/platform/rest/v2/>
