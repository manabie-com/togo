package utils

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	one string
}

type ans1 struct {
	one sql.NullString
}

func TestNullString(t *testing.T) {
	ast := assert.New(t)

	qs := []question1{
		question1{
			para1{
				"firstUser",
			},
			ans1{sql.NullString{
				String: "firstUser",
				Valid:  true,
			}},
		},
	}
	for _, q := range qs {
		a, p := q.ans1, q.para1
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, NullString(p.one), "input:%v", p)
	}
}
