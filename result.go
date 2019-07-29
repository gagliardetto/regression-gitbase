package gitbase

import (
	"fmt"

	regression "github.com/src-d/regression-core"
)

// Comparison struct holds the percentage difference between two results.
type Comparison struct {
	regression.Comparison

	Rows float64
}

// Result holds the resources and number of rows from a version test.
type Result struct {
	*regression.Result
	Query
	Rows int64
}

func NewResult() *Result {
	return &Result{Result: new(regression.Result)}
}

// ComparePrint shows the difference between two results and returns if
// it is within margin.
func (r *Result) ComparePrint(q *Result, allowance float64) bool {
	ok := r.Result.ComparePrint(q.Result, allowance)

	c := Comparison{
		Rows: regression.Percent(r.Rows, q.Rows),
	}
	ok = ok && (c.Rows <= allowance)

	fmt.Printf(regression.CompareFormat,
		"Rows",
		r.Rows,
		q.Rows,
		c.Rows,
		allowance > c.Rows,
	)

	return ok
}
