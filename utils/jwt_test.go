package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one string
	two string
}

type ans struct {
	one string
}

func Test_JWT(t *testing.T) {

	ast := assert.New(t)

	qs := []question{
		question{
			para{
				"firstUser",
				"wqGyEBBfPK9w3Lxw",
			},
			ans{""},
		},
	}

	for _, q := range qs {
		p := q.para
		fmt.Printf("~~%v~~\n", p)
		_, err := CreateToken(p.one, p.two)
		ast.NoError(err)
	}
}
