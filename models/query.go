package models

type Query struct {
	Status     string
	Expression string
}

func (q *Query) IsEmpty() bool {
	return q.Status == "" && q.Expression == ""
}
