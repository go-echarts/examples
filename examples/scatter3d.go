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

func genScatter3dData() [][3]int {
	data := make([][3]int, 0)
	for i := 0; i < 80; i++ {
		data = append(data, [3]int{int(dataSeed.Int63()) % 100, int(dataSeed.Int63()) % 100, int(dataSeed.Int63()) % 100})
	}
	return data
}

func scatter3DBase() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Scatter3D-example",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        100,
			InRange:    &opts.VisualMapInRange{Color: scatter3DColor},
		}),
	)

	scatter3d.AddZAxis("scatter3d", genScatter3dData())
	return scatter3d
}

type Scatter3dExamples struct{}

func (Scatter3dExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		scatter3DBase(),
	)

	f, err := os.Create("examples/html/scatter3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
