package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var wcData = map[string]interface{}{
	"Sam S Club":               10000,
	"Macys":                    6181,
	"Amy Schumer":              4386,
	"Jurassic World":           4055,
	"Charter Communications":   2467,
	"Chick Fil A":              2244,
	"Planet Fitness":           1898,
	"Pitch Perfect":            1484,
	"Express":                  1689,
	"Home":                     1112,
	"Johnny Depp":              985,
	"Lena Dunham":              847,
	"Lewis Hamilton":           582,
	"KXAN":                     555,
	"Mary Ellen Mark":          550,
	"Farrah Abraham":           462,
	"Rita Ora":                 366,
	"Serena Williams":          282,
	"NCAA baseball tournament": 273,
	"Point Break":              265,
}

func generateWCData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wcBase() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic WordCloud example",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func wcCardioid() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "cardioid shape"}),
	)

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
		)
	return wc
}

func wcStar() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "star shape",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "star",
				}),
		)
	return wc
}

type WordcloudExamples struct{}

func (WordcloudExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		wcBase(),
		wcCardioid(),
		wcStar(),
	)

	f, err := os.Create("examples/html/wordcloud.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
