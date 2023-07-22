package test

import (
	"github.com/devhindo/langs.md/output"
)

func Test() {
	dic := make(map[string]int)
	dic["C"] = 100
	dic["python"] = 200
	output.ConstructJSONFile(dic)
}