package examples

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCnt = 7
	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < itemCnt; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func barBasic() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic bar example", Subtitle: "This is the subtitle."}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barTitle() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and legend options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
			Right:    "40%",
		}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80%"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barTooltip() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "tooltip options"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80px"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barSetToolbox() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "toolbox options"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					// set the language
					// Chinese version: ["数据视图", "关闭", "刷新"]
					Lang: []string{"data view", "turn off", "refresh"},
				},
			}},
		),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barShowLabel() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "label options"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)
	return bar
}

func barXYName() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "display the axes name",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "XAxisName",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "YAxisName",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barXYFormatter() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "customized the xaxis/yaxis formatter"}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} x-unit"},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} y-unit"},
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barColor() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "set user-defined colors",
		}),
		charts.WithColorsOpts(opts.Colors{"blue", "pink"}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barSplitLine() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "splitline options",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "XAxisName",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "YAxisName",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barGap() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "set the gap of each bar",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.SetSeriesOptions(
		charts.WithBarChartOpts(opts.BarChart{
			BarGap: "150%",
		}),
	)
	return bar
}

func barDataZoomInside() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "datazoom options(inside)",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "inside",
			Start: 10,
			End:   50,
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barDataZoomSlider() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "datazoom options(slider)",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "slider",
			Start: 10,
			End:   50,
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barReverse() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "reverse xaxis and yaxis",
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.XYReversal()
	return bar
}

func barStack() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "stack style",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
			Stack: "stackA",
		}))
	return bar
}

func barMarkPoints() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options",
		}),
	)

	special := generateBarItems()
	special[0].Value = 100

	bar.SetXAxis(weeks).
		AddSeries("Category A", special, charts.WithMarkPointNameCoordItemOpts(opts.MarkPointNameCoordItem{
			Name:       "special mark",
			Coordinate: []interface{}{"Mon", 100},
			Label: &opts.Label{
				Show:     true,
				Color:    "pink",
				Position: "inside",
			},
		})).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithMarkPointNameTypeItemOpts(
			opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
			opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
		))
	return bar
}

func barMarkLines() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markline options",
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithMarkLineNameTypeItemOpts(
			opts.MarkLineNameTypeItem{Name: "Maximum", Type: "max"},
			opts.MarkLineNameTypeItem{Name: "Avg", Type: "average"},
		))
	return bar
}

func barOverlap() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts"}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		SetSeriesOptions(charts.WithMarkLineNameTypeItemOpts(
			opts.MarkLineNameTypeItem{Name: "Maximum", Type: "max"},
			opts.MarkLineNameTypeItem{Name: "Avg", Type: "average"},
		))
	bar.Overlap(lineBase())
	bar.Overlap(scatterBase())
	return bar
}

func barSize() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "adjust canvas size",
			Subtitle: "I want a bigger canvas size :)",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "600px",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

type Exampler interface {
	Examples()
}

type BarExamples struct{}

func (BarExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		barBasic(),
		barTitle(),
		barTooltip(),
		barSetToolbox(),
		barShowLabel(),
		barXYName(),
		barXYFormatter(),
		barColor(),
		barSplitLine(),
		barGap(),
		barDataZoomInside(),
		barDataZoomSlider(),
		barReverse(),
		barStack(),
		barMarkPoints(),
		barMarkLines(),
		barOverlap(),
		barSize(),
	)
	f, err := os.Create("examples/html/bar.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
