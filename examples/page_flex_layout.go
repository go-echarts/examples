package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

type PageFlexLayoutExamples struct{}

func (PageFlexLayoutExamples) Examples() {
	page := genPages()
	page.SetLayout(components.PageFlexLayout)
	f, err := os.Create("examples/html/page_flex_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
