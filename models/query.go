package models

type Query struct {
	Status     string
	Expression string
}

func NewQuery(status, expression string) *Query {
	return &Query{
		Status:     status,
		Expression: expression,
	}
}

func (q *Query) IsEmpty() bool {
	return q.Status == "" &&
		q.Expression == ""
}
