package models

type Query struct {
	Status, Expression string
}

func (q *Query) IsEmpty() bool {
	return q.Status == "" && q.Expression == ""
}
