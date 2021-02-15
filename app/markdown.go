
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

https://github.com/psychedelicnekopunch/go-sample

[go-sample](https://github.com/psychedelicnekopunch/go-sample)
`

	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.HrefTargetBlank,
	})
	// output := blackfriday.Run([]byte(s), blackfriday.WithExtensions(blackfriday.HardLineBreak), blackfriday.WithRenderer(renderer))
	output := blackfriday.Run([]byte(s), blackfriday.WithExtensions(blackfriday.HardLineBreak + blackfriday.Autolink), blackfriday.WithRenderer(renderer))
	html := bluemonday.UGCPolicy().SanitizeBytes(output)

	fmt.Print(template.HTML(string(html)))
}
