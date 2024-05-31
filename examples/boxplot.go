package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/montanaflynn/stats"
)

var (
	bpX = [...]string{"expr1", "expr2", "expr3", "expr4", "expr5"}
	bpY = [][]float64{
		{850, 740, 900, 1070, 930, 850, 950, 980, 980, 880, 1000, 980, 930, 650, 760, 810, 1000, 1000, 960, 960},
		{960, 940, 960, 940, 880, 800, 850, 880, 900, 840, 830, 790, 810, 880, 880, 830, 800, 790, 760, 800},
		{880, 880, 880, 860, 720, 720, 620, 860, 970, 950, 880, 910, 850, 870, 840, 840, 850, 840, 840, 840},
		{890, 810, 810, 820, 800, 770, 760, 740, 750, 760, 910, 920, 890, 860, 880, 720, 840, 850, 850, 780},
		{890, 840, 780, 810, 760, 810, 790, 810, 820, 850, 870, 870, 810, 740, 810, 940, 950, 800, 810, 870},
	}
)

func createBoxPlotData(data []float64) []float64 {
	min, _ := stats.Min(data)
	max, _ := stats.Max(data)
	q, _ := stats.Quartile(data)

	return []float64{
		min,
		q.Q1,
		q.Q2,
		q.Q3,
		max,
	}
}

func generateBoxPlotItems(boxPlotData [][]float64) []opts.BoxPlotData {
	items := make([]opts.BoxPlotData, 0)
	for i := 0; i < len(boxPlotData); i++ {
		items = append(items, opts.BoxPlotData{Value: createBoxPlotData(boxPlotData[i])})
	}
	return items

}

func boxPlotBase() *charts.BoxPlot {
	bp := charts.NewBoxPlot()
	bp.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic boxplot example"}),
	)

	bp.SetXAxis(bpX).AddSeries("boxplot", generateBoxPlotItems(bpY))
	return bp
}

func boxPlotMulti() *charts.BoxPlot {
	bp := charts.NewBoxPlot()
	bp.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "boxplot with multi-series"}),
	)

	bp.SetXAxis(bpX[:2]).
		AddSeries("series1", generateBoxPlotItems(bpY[:2])).
		AddSeries("series2", generateBoxPlotItems(bpY[2:]))
	return bp
}

type BoxplotExamples struct{}

func (BoxplotExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		boxPlotBase(),
		boxPlotMulti(),
	)
	f, err := os.Create("examples/html/boxplot.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
