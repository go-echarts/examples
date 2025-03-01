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
	itemCntLine = 6
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func generateLineItemsTwoAxis(points int, xFunc func(int) interface{}) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < points; i++ {
		items = append(items, opts.LineData{Value: []interface{}{xFunc(i), 100 + rand.Intn(20)}})
	}
	return items
}

func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

func lineBase() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and label options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "Average", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: opts.Bool(true)}}),
		)
	return line
}

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "splitline options",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
			},
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems(),
		charts.WithLabelOpts(
			opts.Label{Show: opts.Bool(true)},
		))
	return line
}

func lineNumerical() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "numerical X axis & accessories",
			Subtitle: "styled mark areas, mark lines and visual maps (area below line pieces)",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Max: 200,
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Type:      "piecewise",
			Dimension: "0",
			Pieces: []opts.Piece{
				{Gt: 1, Lt: 7, Color: "rgba(50,50,250,0.4)"},
				{Gt: 10, Lt: 15, Color: "rgba(50,50,250,0.5)"},
			},
			Show: opts.Bool(false),
		}),
	)

	line.AddSeries("Category A", generateLineItemsTwoAxis(30, func(i int) interface{} { return i }),
		charts.WithLineChartOpts(opts.LineChart{
			Symbol:     "triangle",
			SymbolSize: 10,
		}),
		charts.WithAreaStyleOpts(opts.AreaStyle{}),
		charts.WithLineStyleOpts(opts.LineStyle{
			Color: "rgba(50,50,250,0.7)",
		}),
		charts.WithItemStyleOpts(opts.ItemStyle{
			Color: "rgba(50,50,250,0.7)",
		}),
		charts.WithMarkAreaNameCoordItemOpts(
			opts.MarkAreaNameCoordItem{
				Name:        "Danger zone",
				Coordinate0: []interface{}{20},
				Coordinate1: []interface{}{25},
				Label:       &opts.Label{Show: opts.Bool(true), Position: "middle"},
				ItemStyle:   &opts.ItemStyle{Color: "rgba(255, 173, 177, 0.5)"},
			},
		),
		charts.WithMarkLineStyleOpts(opts.MarkLineStyle{
			Symbol:     []string{"square", "circle"},
			SymbolSize: 10,
		}),
		charts.WithMarkLineNameCoordItemOpts(opts.MarkLineNameCoordItem{
			Name:        "Danger level",
			Coordinate0: []interface{}{20, 10},
			Coordinate1: []interface{}{25, 50},
		}),
		charts.WithMarkLineNameXAxisItemOpts(
			opts.MarkLineNameXAxisItem{
				Name:  "Line of no return",
				XAxis: 28,
			},
		),
	)
	return line
}

func lineTime() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "temporal X axis",
			Subtitle: "time.Date as X axis values",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Min: 0,
			Max: 200,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Type: "time",
			Min:  time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
		}),
		charts.WithTooltipOpts(opts.Tooltip{ // Potential to string format tooltip here
			Show:    opts.Bool(true),
			Trigger: "axis",
		}),
	)

	line.AddSeries("Category A", generateLineItemsTwoAxis(50, func(i int) interface{} { return time.Date(2025, time.February, i, 0, 0, 0, 0, time.Local) }))
	return line
}

func lineStep() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "step style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Step: true,
			}),
		)
	return line
}

func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "smooth style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: opts.Bool(true),
			}),
		)
	return line
}

func lineArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "area options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: opts.Bool(true),
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.2,
				}),
			charts.WithMarkAreaNameCoordItemOpts(
				opts.MarkAreaNameCoordItem{
					Name:        "In stock",
					Coordinate0: []interface{}{2},
					Coordinate1: []interface{}{4},
					Label:       &opts.Label{Show: opts.Bool(true), Position: "middle"},
					ItemStyle:   &opts.ItemStyle{Color: "rgba(82, 228, 167, 0.5)"},
				},
			),
		)
	return line
}

func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "smooth area"}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.2,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: opts.Bool(true),
			}),
		)
	return line
}

func lineOverlap() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts"}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	line.Overlap(esEffectStyle())
	line.Overlap(scatterBase())
	return line
}

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category  A", generateLineItems()).
		AddSeries("Category  B", generateLineItems()).
		AddSeries("Category  C", generateLineItems()).
		AddSeries("Category  D", generateLineItems())
	return line
}

func lineDemo() *charts.Line {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Search Time: Hash table vs Binary search",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Cost time(ns)",
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
			},
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Elements",
		}),
	)

	line.SetXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
		AddSeries("map", generateLineItems(),
			charts.WithLabelOpts(opts.Label{Show: opts.Bool(true), Position: "bottom"})).
		AddSeries("slice", generateLineData([]float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131}),
			charts.WithLabelOpts(opts.Label{Show: opts.Bool(true), Position: "top"})).
		SetSeriesOptions(
			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
				Name: "Average",
				Type: "average",
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: opts.Bool(true),
			}),
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show:      opts.Bool(true),
					Formatter: "{a}: {b}",
				},
			}),
		)

	return line
}

func lineSymbols() *charts.Line {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "symbol options",
			Subtitle: "tooltip with 'axis' trigger",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}),
	)

	// Put data into instance
	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}

type LineExamples struct{}

func (LineExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		lineShowLabel(),
		lineSymbols(),
		lineMarkPoint(),
		lineSplitLine(),
		lineNumerical(),
		lineTime(),
		lineStep(),
		lineSmooth(),
		lineArea(),
		lineSmoothArea(),
		lineOverlap(),
		lineMulti(),
		lineDemo(),
	)
	f, err := os.Create("examples/html/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
