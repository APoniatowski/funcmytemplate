package funcmytemplate

import (
	"text/template"

	"github.com/APoniatowski/funcmytemplate/examples"
)

var Add = template.FuncMap{
	"CurrentDate": examples.CurrentDate,
	// ... other functions
}
