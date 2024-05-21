package examples

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

// HeaderTpl copy from go-echarts/templates/header.go
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
	render.BaseRender
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

func renderSnippets() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "My renderSnippets",
	}))

	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())

	chartSnippet := bar.RenderSnippet()

	tmpl := "{{.Element  }} {{.Script}} {{.Option}}"
	t := template.New("snippet")
	t, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	data := struct {
		Element template.HTML
		Script  template.HTML
		Option  template.HTML
	}{
		Element: template.HTML(chartSnippet.Element),
		Script:  template.HTML(chartSnippet.Script),
		Option:  template.HTML(chartSnippet.Option),
	}

	fmt.Println("------ Render snippets output:")
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

type CustomizeExamples struct{}

func (CustomizeExamples) Examples() {
	// print it instead of generated files
	fmt.Println("Customer Render outputs:")
	renderSnippets()
	fmt.Println("------ Render customer render output:")
	bar := barCustomize()
	err := bar.Render(os.Stdout)
	if err != nil {
		panic(err)
	}
}
