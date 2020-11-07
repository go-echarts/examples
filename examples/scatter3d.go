package examples

import (
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	dataSeed = rand.NewSource(time.Now().UnixNano())

	scatter3DColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}
)

func genScatter3dData() []opts.Chart3DData {
	data := make([]opts.Chart3DData, 0)
	for i := 0; i < 80; i++ {
		data = append(data, opts.Chart3DData{Value: []interface{}{
			int(dataSeed.Int63()) % 100,
			int(dataSeed.Int63()) % 100,
			int(dataSeed.Int63()) % 100},
		})
	}
	return data
}

func scatter3DBase() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic Scatter3D example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        100,
			InRange:    &opts.VisualMapInRange{Color: scatter3DColor},
		}),
	)

	scatter3d.AddSeries("scatter3d", genScatter3dData())
	return scatter3d
}

func scatter3DDataItem() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "user-defined item style"}),
		charts.WithXAxis3DOpts(opts.XAxis3D{Name: "MY-X-AXIS", Show: true}),
		charts.WithYAxis3DOpts(opts.YAxis3D{Name: "MY-Y-AXIS"}),
		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "MY-Z-AXIS"}),
	)

	scatter3d.AddSeries("scatter3d", []opts.Chart3DData{
		{Name: "point1", Value: []interface{}{10, 10, 10}, ItemStyle: &opts.ItemStyle{Color: "green"}},
		{Name: "point2", Value: []interface{}{15, 15, 15}, ItemStyle: &opts.ItemStyle{Color: "blue"}},
		{Name: "point3", Value: []interface{}{20, 20, 20}, ItemStyle: &opts.ItemStyle{Color: "red"}},
	})

	return scatter3d
}

type Scatter3dExamples struct{}

func (Scatter3dExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		scatter3DBase(),
		scatter3DDataItem(),
	)

	f, err := os.Create("examples/html/scatter3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
