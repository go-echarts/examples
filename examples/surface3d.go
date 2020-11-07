package examples

import (
	"io"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var surfaceRangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genSurface3dData0() []opts.Chart3DData {
	data := make([][3]interface{}, 0)
	for i := -60; i < 60; i++ {
		y := float64(i) / 60
		for j := -60; j < 60; j++ {
			x := float64(j) / 60
			z := math.Sin(x*math.Pi) * math.Sin(y*math.Pi)
			data = append(data, [3]interface{}{x, y, z})
		}
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{
			Value: []interface{}{d[0], d[1], d[2]},
		})
	}
	return ret
}

func genSurface3dData1() []opts.Chart3DData {
	data := make([][3]interface{}, 0)
	for i := -30; i < 30; i++ {
		y := float64(i) / 10
		for j := -30; j < 30; j++ {
			x := float64(j) / 10
			z := math.Sin(x*x+y*y) * x / math.Pi
			data = append(data, [3]interface{}{x, y, z})
		}
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{
			Value: []interface{}{d[0], d[1], d[2]},
		})
	}
	return ret
}

func surface3DBase() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic surface3D example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        3,
			Min:        -3,
		}),
	)

	surface3d.AddSeries("surface3d", genSurface3dData0())
	return surface3d
}

func surface3DRose() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Rose style"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        3,
			Min:        -3,
		}),
	)

	surface3d.AddSeries("surface3d", genSurface3dData1())
	return surface3d
}

type Surface3dExamples struct{}

func (Surface3dExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		surface3DBase(),
		surface3DRose(),
	)

	f, err := os.Create("examples/html/surface3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
