package main

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
)

var (
	itemCntScatter = 6
	sports         = []string{"Swimming", "Surfing", "Shooting ", "Skating", "Wrestling", "Diving"}
)

func generateScatterItems() []opts.ScatterData {
	items := make([]opts.ScatterData, 0)
	for i := 0; i < itemCntScatter; i++ {
		items = append(items, opts.ScatterData{
			Value:        rand.Intn(100),
			Symbol:       "roundRect",
			SymbolSize:   20,
			SymbolRotate: 10,
		})
	}
	return items
}
func scatterBase() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Scatter-basic-example",
		}),
	)

	scatter.SetXAxis(sports).
		AddSeries("Category A", generateScatterItems()).
		AddSeries("Category B", generateScatterItems())

	return scatter
}

func scatterShowLabel() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "Scatter-with-label-opts",
		}),
	)

	scatter.SetXAxis(sports).
		AddSeries("Category A", generateScatterItems()).
		AddSeries("Category B", generateScatterItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:     true,
				Position: "right",
			}),
		)
	return scatter
}

func scatterSplitLine() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Scatter-splitline",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Sports",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Score",
			SplitLine: &opts.SplitLine{
				Show: true,
			}}),
	)

	scatter.SetXAxis(sports).
		AddSeries("Player A", generateScatterItems()).
		AddSeries("Player B", generateScatterItems())
	return scatter
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		scatterBase(),
		scatterShowLabel(),
		scatterSplitLine(),
	)
	f, err := os.Create("html/scatter.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(os.Stdout, f))
}
