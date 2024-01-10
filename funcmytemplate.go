package funcmytemplate

import (
	"text/template"

	"github.com/APoniatowski/funcmytemplate/examples"
)

type funcmytemplate interface {
	Add() template.FuncMap
}

func Add() template.FuncMap {
	funcdTemplate := template.FuncMap{
		"CurrentDate": examples.CurrentDate,
		// more functions go here
	}

	return funcdTemplate
}
