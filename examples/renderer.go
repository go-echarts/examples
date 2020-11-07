package examples

import (
	"bytes"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

// copy from go-echarts/templates/header.go
// Now I want to customize my own Header (or tpls.BaseTpl / tpls.ChartTpl) template
var HeaderTpl = `
{{ define "header" }}
<head>
   <meta charset="utf-8">
   <title>{{ .PageTitle }} --> This is my own style template 🐶</title>
{{- range .JSAssets.Values }}
   <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
   <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ end }}
`

type myOwnRender struct {
	c      interface{}
	before []func()
}

func NewMyOwnRender(c interface{}, before ...func()) render.Renderer {
	return &myOwnRender{c: c, before: before}
}

func (r *myOwnRender) Render(w io.Writer) error {
	for _, fn := range r.before {
		fn()
	}

	contents := []string{HeaderTpl, tpls.BaseTpl, tpls.ChartTpl}
	tpl := render.MustTemplate("chart", contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "chart", r.c); err != nil {
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}

func barCustomize() *charts.Bar {
	bar := charts.NewBar()
	bar.Renderer = NewMyOwnRender(bar, bar.Validate)
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Bar-customize-template-example",
			Subtitle: "This is the subtitle.",
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

type CustomizeExamples struct{}

func (CustomizeExamples) Examples() {
	bar := barCustomize()
	f, err := os.Create("examples/html/customize.html")
	if err != nil {
		panic(err)
	}
	bar.Render(io.MultiWriter(f))
}
