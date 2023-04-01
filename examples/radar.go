package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	radarDataBJ = [][]float32{
		{55, 9, 56, 0.46, 18, 6, 1},
		{25, 11, 21, 0.65, 34, 9, 2},
		{56, 7, 63, 0.3, 14, 5, 3},
		{33, 7, 29, 0.33, 16, 6, 4},
		{42, 24, 44, 0.76, 40, 16, 5},
		{82, 58, 90, 1.77, 68, 33, 6},
		{74, 49, 77, 1.46, 48, 27, 7},
		{78, 55, 80, 1.29, 59, 29, 8},
		{267, 216, 280, 4.8, 108, 64, 9},
		{185, 127, 216, 2.52, 61, 27, 10},
		{39, 19, 38, 0.57, 31, 15, 11},
		{41, 11, 40, 0.43, 21, 7, 12},
		{64, 38, 74, 1.04, 46, 22, 13},
		{108, 79, 120, 1.7, 75, 41, 14},
		{108, 63, 116, 1.48, 44, 26, 15},
		{33, 6, 29, 0.34, 13, 5, 16},
		{94, 66, 110, 1.54, 62, 31, 17},
		{186, 142, 192, 3.88, 93, 79, 18},
		{57, 31, 54, 0.96, 32, 14, 19},
		{22, 8, 17, 0.48, 23, 10, 20},
		{39, 15, 36, 0.61, 29, 13, 21},
	}

	radarDataGZ = [][]float32{
		{26, 37, 27, 1.163, 27, 13, 1},
		{85, 62, 71, 1.195, 60, 8, 2},
		{78, 38, 74, 1.363, 37, 7, 3},
		{21, 21, 36, 0.634, 40, 9, 4},
		{41, 42, 46, 0.915, 81, 13, 5},
		{56, 52, 69, 1.067, 92, 16, 6},
		{64, 30, 28, 0.924, 51, 2, 7},
		{55, 48, 74, 1.236, 75, 26, 8},
		{76, 85, 113, 1.237, 114, 27, 9},
		{91, 81, 104, 1.041, 56, 40, 10},
		{84, 39, 60, 0.964, 25, 11, 11},
		{64, 51, 101, 0.862, 58, 23, 12},
		{70, 69, 120, 1.198, 65, 36, 13},
		{77, 105, 178, 2.549, 64, 16, 14},
		{109, 68, 87, 0.996, 74, 29, 15},
		{73, 68, 97, 0.905, 51, 34, 16},
		{54, 27, 47, 0.592, 53, 12, 17},
		{51, 61, 97, 0.811, 65, 19, 18},
		{91, 71, 121, 1.374, 43, 18, 19},
		{73, 102, 182, 2.787, 44, 19, 20},
		{73, 50, 76, 0.717, 31, 20, 21},
	}

	radarDataSH = [][]float32{
		{91, 45, 125, 0.82, 34, 23, 1},
		{65, 27, 78, 0.86, 45, 29, 2},
		{83, 60, 84, 1.09, 73, 27, 3},
		{109, 81, 121, 1.28, 68, 51, 4},
		{106, 77, 114, 1.07, 55, 51, 5},
		{109, 81, 121, 1.28, 68, 51, 6},
		{106, 77, 114, 1.07, 55, 51, 7},
		{89, 65, 78, 0.86, 51, 26, 8},
		{53, 33, 47, 0.64, 50, 17, 9},
		{80, 55, 80, 1.01, 75, 24, 10},
		{117, 81, 124, 1.03, 45, 24, 11},
		{99, 71, 142, 1.1, 62, 42, 12},
		{95, 69, 130, 1.28, 74, 50, 13},
		{116, 87, 131, 1.47, 84, 40, 14},
		{108, 80, 121, 1.3, 85, 37, 15},
		{134, 83, 167, 1.16, 57, 43, 16},
		{79, 43, 107, 1.05, 59, 37, 17},
		{71, 46, 89, 0.86, 64, 25, 18},
		{97, 71, 113, 1.17, 88, 31, 19},
		{84, 57, 91, 0.85, 55, 31, 20},
		{87, 63, 101, 0.9, 56, 41, 21},
	}
)

func generateRadarItems(radarData [][]float32) []opts.RadarData {
	items := make([]opts.RadarData, 0)
	for i := 0; i < len(radarData); i++ {
		items = append(items, opts.RadarData{Value: radarData[i]})
	}
	return items

}

var indicators = []*opts.Indicator{
	{Name: "AQI", Max: 300},
	{Name: "PM2.5", Max: 250},
	{Name: "PM10", Max: 300},
	{Name: "CO", Max: 5},
	{Name: "NO2", Max: 200},
	{Name: "SO2", Max: 100},
}

func radarBase() *charts.Radar {
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic radar example",
		}),
		charts.WithRadarComponentOpts(opts.RadarComponent{
			Indicator: indicators,
			SplitArea: &opts.SplitArea{Show: true},
			SplitLine: &opts.SplitLine{Show: true},
		}),
	)
	radar.AddSeries("Beijing", generateRadarItems(radarDataBJ))
	return radar
}

func radarStyle() *charts.Radar {
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:      "style options",
			Right:      "center",
			TitleStyle: &opts.TextStyle{Color: "#eee"},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "#161627",
		}),
		charts.WithRadarComponentOpts(opts.RadarComponent{
			Indicator:   indicators,
			Shape:       "circle",
			SplitNumber: 5,
			SplitLine: &opts.SplitLine{
				Show: true,
				LineStyle: &opts.LineStyle{
					Opacity: 0.1,
				},
			},
		}),

		charts.WithLegendOpts(opts.Legend{
			Show:   true,
			Bottom: "5px",
			TextStyle: &opts.TextStyle{
				Color: "#eee",
			},
		}),
	)

	radar.AddSeries("Beijing", generateRadarItems(radarDataBJ)).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{
				Width:   1,
				Opacity: 0.5,
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.1,
			}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color: "#F9713C",
			}),
		)
	return radar
}

func radarLegendMulti() *charts.Radar {
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Legend(Multi)",
			Right: "center",
			TitleStyle: &opts.TextStyle{
				Color: "#eee",
			},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "#161627",
		}),
		charts.WithRadarComponentOpts(opts.RadarComponent{
			Indicator:   indicators,
			Shape:       "circle",
			SplitNumber: 5,
			SplitLine: &opts.SplitLine{
				Show: true,
				LineStyle: &opts.LineStyle{
					Opacity: 0.1,
				},
			},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show:   true,
			Bottom: "5px",
			TextStyle: &opts.TextStyle{
				Color: "#eee",
			},
		}),
	)

	radar.AddSeries("Beijing", generateRadarItems(radarDataBJ), charts.WithItemStyleOpts(opts.ItemStyle{Color: "#F9713C"})).
		AddSeries("Guangzhou", generateRadarItems(radarDataGZ), charts.WithItemStyleOpts(opts.ItemStyle{Color: "#B3E4A1"})).
		AddSeries("Shanghai", generateRadarItems(radarDataSH), charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(238, 197, 102)"})).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{
				Width:   1,
				Opacity: 0.5,
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.1,
			}),
		)
	return radar
}

func radarLegendSingle() *charts.Radar {
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Legend(Single)",
			Right: "center",
			TitleStyle: &opts.TextStyle{
				Color: "#eee",
			},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "#161627",
		}),
		charts.WithRadarComponentOpts(opts.RadarComponent{
			Indicator:   indicators,
			Shape:       "circle",
			SplitNumber: 5,
			SplitLine: &opts.SplitLine{
				Show: true,
				LineStyle: &opts.LineStyle{
					Opacity: 0.1,
				},
			},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show:   true,
			Bottom: "5px",
			TextStyle: &opts.TextStyle{
				Color: "#eee",
			},
			SelectedMode: "single",
		}),
	)
	radar.AddSeries("Beijing", generateRadarItems(radarDataBJ),
		charts.WithItemStyleOpts(opts.ItemStyle{Color: "#F9713C"})).
		AddSeries("Guangzhou", generateRadarItems(radarDataGZ),
			charts.WithItemStyleOpts(opts.ItemStyle{Color: "#B3E4A1"})).
		AddSeries("Shanghai", generateRadarItems(radarDataSH),
			charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(238, 197, 102)"})).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{Width: 1, Opacity: 0.5}),
			charts.WithAreaStyleOpts(opts.AreaStyle{Opacity: 0.1}),
		)
	return radar
}

type RadarExamples struct{}

func (RadarExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		radarBase(),
		radarStyle(),
		radarLegendMulti(),
		radarLegendSingle(),
	)
	f, err := os.Create("examples/html/radar.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
