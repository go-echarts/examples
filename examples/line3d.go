package examples

import (
	"io"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genLine3dData() []opts.Chart3DData {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{(1 + 0.25*math.Cos(75*t)) * math.Cos(t), (1 + 0.25*math.Cos(75*t)) * math.Sin(t), t + 2.0*math.Sin(75.0*t)},
		)
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{Value: []interface{}{d[0], d[1], d[2]}})
	}
	return ret
}

func line3DBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line3d example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),
	)

	line3d.AddSeries("line3D", genLine3dData())
	return line3d
}

func line3DAutoRotate() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "auto rotating"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),

		charts.WithGrid3DOpts(opts.Grid3D{
			ViewControl: &opts.ViewControl{
				AutoRotate: true,
			},
		}),
	)

	line3d.AddSeries("line3D", genLine3dData())
	return line3d
}

type Line3dExamples struct{}

func (Line3dExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		line3DBase(),
		line3DAutoRotate(),
	)

	f, err := os.Create("examples/html/line3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
