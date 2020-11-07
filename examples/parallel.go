package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	parallelDataBJ = [][]interface{}{
		{1, 55, 9, 56, 0.46, 18, 6, "Moderate"},
		{2, 25, 11, 21, 0.65, 34, 9, "Good"},
		{3, 56, 7, 63, 0.3, 14, 5, "Moderate"},
		{4, 33, 7, 29, 0.33, 16, 6, "Good"},
		{5, 42, 24, 44, 0.76, 40, 16, "Good"},
		{6, 82, 58, 90, 1.77, 68, 33, "Moderate"},
		{7, 74, 49, 77, 1.46, 48, 27, "Moderate"},
		{8, 78, 55, 80, 1.29, 59, 29, "Moderate"},
		{9, 267, 216, 280, 4.8, 108, 64, "Heavily"},
		{10, 185, 127, 216, 2.52, 61, 27, "Moderately"},
		{11, 39, 19, 38, 0.57, 31, 15, "Good"},
		{12, 41, 11, 40, 0.43, 21, 7, "Good"},
		{13, 64, 38, 74, 1.04, 46, 22, "Moderate"},
		{14, 108, 79, 120, 1.7, 75, 41, "Lightly"},
		{15, 108, 63, 116, 1.48, 44, 26, "Lightly"},
		{16, 33, 6, 29, 0.34, 13, 5, "Good"},
		{17, 94, 66, 110, 1.54, 62, 31, "Moderate"},
		{18, 186, 142, 192, 3.88, 93, 79, "Moderately"},
		{19, 57, 31, 54, 0.96, 32, 14, "Moderate"},
		{20, 22, 8, 17, 0.48, 23, 10, "Good"},
		{21, 39, 15, 36, 0.61, 29, 13, "Good"},
	}

	parallelDataGZ = [][]interface{}{
		{1, 26, 37, 27, 1.163, 27, 13, "Good"},
		{2, 85, 62, 71, 1.195, 60, 8, "Moderate"},
		{3, 78, 38, 74, 1.363, 37, 7, "Moderate"},
		{4, 21, 21, 36, 0.634, 40, 9, "Good"},
		{5, 41, 42, 46, 0.915, 81, 13, "Good"},
		{6, 56, 52, 69, 1.067, 92, 16, "Moderate"},
		{7, 64, 30, 28, 0.924, 51, 2, "Moderate"},
		{8, 55, 48, 74, 1.236, 75, 26, "Moderate"},
		{9, 76, 85, 113, 1.237, 114, 27, "Moderate"},
		{10, 91, 81, 104, 1.041, 56, 40, "Moderate"},
		{11, 84, 39, 60, 0.964, 25, 11, "Moderate"},
		{12, 64, 51, 101, 0.862, 58, 23, "Moderate"},
		{13, 70, 69, 120, 1.198, 65, 36, "Moderate"},
		{14, 77, 105, 178, 2.549, 64, 16, "Moderate"},
		{15, 109, 68, 87, 0.996, 74, 29, "Lightly"},
		{16, 73, 68, 97, 0.905, 51, 34, "Moderate"},
		{17, 54, 27, 47, 0.592, 53, 12, "Moderate"},
		{18, 51, 61, 97, 0.811, 65, 19, "Moderate"},
		{19, 91, 71, 121, 1.374, 43, 18, "Moderate"},
		{20, 73, 102, 182, 2.787, 44, 19, "Moderate"},
		{21, 73, 50, 76, 0.717, 31, 20, "Moderate"},
	}

	parallelDataSH = [][]interface{}{
		{1, 91, 45, 125, 0.82, 34, 23, "Moderate"},
		{2, 65, 27, 78, 0.86, 45, 29, "Moderate"},
		{3, 83, 60, 84, 1.09, 73, 27, "Moderate"},
		{4, 109, 81, 121, 1.28, 68, 51, "Lightly"},
		{5, 106, 77, 114, 1.07, 55, 51, "Lightly"},
		{6, 109, 81, 121, 1.28, 68, 51, "Lightly"},
		{7, 106, 77, 114, 1.07, 55, 51, "Lightly"},
		{8, 89, 65, 78, 0.86, 51, 26, "Moderate"},
		{9, 53, 33, 47, 0.64, 50, 17, "Moderate"},
		{10, 80, 55, 80, 1.01, 75, 24, "Moderate"},
		{11, 117, 81, 124, 1.03, 45, 24, "Lightly"},
		{12, 99, 71, 142, 1.1, 62, 42, "Moderate"},
		{13, 95, 69, 130, 1.28, 74, 50, "Moderate"},
		{14, 116, 87, 131, 1.47, 84, 40, "Lightly"},
		{15, 108, 80, 121, 1.3, 85, 37, "Lightly"},
		{16, 134, 83, 167, 1.16, 57, 43, "Lightly"},
		{17, 79, 43, 107, 1.05, 59, 37, "Moderate"},
		{18, 71, 46, 89, 0.86, 64, 25, "Moderate"},
		{19, 97, 71, 113, 1.17, 88, 31, "Moderate"},
		{20, 84, 57, 91, 0.85, 55, 31, "Moderate"},
		{21, 87, 63, 101, 0.9, 56, 41, "Moderate"},
	}

	parallelAxisList = []opts.ParallelAxis{
		{Dim: 0, Name: "Date", Inverse: true, Max: 31, NameLocation: "start"},
		{Dim: 1, Name: "AQI"},
		{Dim: 2, Name: "PM2.5"},
		{Dim: 3, Name: "PM10"},
		{Dim: 4, Name: "CO"},
		{Dim: 5, Name: "NO2"},
		{Dim: 6, Name: "SO2"},
		{Dim: 7, Name: "Level", Type: "category", Data: []string{"Good", "Moderate", "Lightly", "Moderately", "Heavily", "Severely"}},
	}
)

func generateParallelData(data [][]interface{}) []opts.ParallelData {
	items := make([]opts.ParallelData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.ParallelData{Value: data[i]})
	}
	return items
}

func parallelBase() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic Parallel example",
		}),
		charts.WithParallelAxisList(parallelAxisList),
		charts.WithLegendOpts(opts.Legend{Show: true}),
	)

	parallel.AddSeries("Beijing", generateParallelData(parallelDataBJ))
	return parallel
}

func parallelComponent() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "with Component",
		}),
		charts.WithParallelComponentOpts(opts.ParallelComponent{
			Left:   "15%",
			Right:  "13%",
			Bottom: "10%",
			Top:    "20%",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithParallelAxisList(parallelAxisList),
	)

	parallel.AddSeries("Beijing", generateParallelData(parallelDataBJ))
	return parallel
}

func parallelMulti() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Multi Series",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithParallelAxisList(parallelAxisList),
	)

	parallel.AddSeries("Beijing", generateParallelData(parallelDataBJ)).
		AddSeries("Guangzhou", generateParallelData(parallelDataGZ)).
		AddSeries("Shanghai", generateParallelData(parallelDataSH))
	return parallel
}

type ParallelExamples struct{}

func (ParallelExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		parallelBase(),
		parallelComponent(),
		parallelMulti(),
	)
	f, err := os.Create("examples/html/parallel.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
