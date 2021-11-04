package done

type Jira struct {
	User, Pass, FQDN, TicketNumber, Comment string
}

type Payload struct {
	Update     Update     `json:"update"`
	Fields     Fields     `json:"fields"`
	Transition Transition `json:"transition"`
}

type Transition struct {
	ID string `json:"id"`
}
type Add struct {
	Body string `json:"body"`
}
type Resolution struct {
	Name string `json:"name"`
}
type Update struct {
	Comment []Comment `json:"comment"`
}
type Comment struct {
	Add Add `json:"add"`
}

type Fields struct {
	Resolution Resolution `json:"resolution"`
}
