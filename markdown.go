
package main


import (
	"fmt"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)


func main() {

	s := `
# title

* list1
* list2


## title2

this is markdown.
hello, world.
`

	output := blackfriday.Run([]byte(s), blackfriday.WithExtensions(blackfriday.HardLineBreak))
	html := bluemonday.UGCPolicy().SanitizeBytes(output)

	fmt.Print(template.HTML(string(html)))
}
