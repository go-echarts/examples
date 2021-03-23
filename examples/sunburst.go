package examples

import (
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateSunburstItems() []opts.SunBurstData {
	items := make([]opts.SunBurstData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.SunBurstData{
			Value: rand.Float64(),
			Name:  "parent-" + strconv.Itoa(i),
			Children: []*opts.SunBurstData{
				{
					Value: rand.Float64(),
					Name:  "child-" + strconv.Itoa(i),
				},
			},
		})
	}
	return items
}

func sunburstBase() *charts.Sunburst {
	sunburst := charts.NewSunburst()
	sunburst.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic sunburst example"}),
	)
	sunburst.AddSeries("sunburst", generateSunburstItems())
	return sunburst
}

type SunburstExample struct{}

func (SunburstExample) Examples() {
	page := components.NewPage()
	page.AddCharts(
		sunburstBase(),
	)
	f, err := os.Create("examples/html/sunburst.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))

}
