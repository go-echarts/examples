package examples

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for sunburst chart
func generateItems() []opts.SunBurstData {
	items := make([]opts.SunBurstData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.SunBurstData{
			Value: rand.Float64(),
			Name:  "11111",
			Children: []*opts.SunBurstData{
				{
					Value: rand.Float64(),
					Name:  "2222",
				},
			},
		})
	}
	return items
}

func main() {
	sunburst := charts.NewSunburst()
	sunburst.SetGlobalOptions(
		charts.WithTitleOpts((opts.Title{Title: "Sunburst"})),
	)
	sunburst.AddSeries("sunburst", generateItems()).SetSeriesOptions(
		charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "12345s",
			},
		),
		charts.WithSunburstOpts(
			opts.SunburstChart{
				Animation: true,
			},
		),
	)
	// Where the magic happens
	f, _ := os.Create("test.html")
	sunburst.Render(f)
}
