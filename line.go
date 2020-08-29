package main

import (
	"github.com/go-echarts/go-echarts/components"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/opts"
)

var (
	itemCntLine = 6
	nameItems   = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
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
		charts.WithTitleOpts(opts.Title{Title: "Line-示例图"}),
	)

	line.SetXAxis(nameItems).
		AddSeries("商家A", generateLineItems())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-显示 Label",
		}),
	)

	line.SetXAxis(nameItems).
		AddSeries("商家A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-标记点",
		}),
	)

	line.SetXAxis(nameItems).AddSeries("商家A", generateLineItems()).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "最大值", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "平均值", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "最小值", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
		)
	return line
}

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-显示分割线",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
	)

	line.SetXAxis(nameItems).AddSeries("商家A", generateLineItems(),
		charts.WithLabelOpts(
			opts.Label{Show: true},
		))
	return line
}

func lineStep() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-阶梯图",
		}),
	)

	line.SetXAxis(nameItems).AddSeries("商家A", generateLineItems()).
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
			Title: "Line-平滑曲线",
		}),
	)

	line.SetXAxis(nameItems).AddSeries("商家A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: true,
			}),
		)
	return line
}

func lineArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-填充区域",
		}),
	)

	line.SetXAxis(nameItems).AddSeries("商家A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: true,
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.2,
				}),
		)
	return line
}

func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-平滑区域",
		}),
	)

	line.SetXAxis(nameItems).AddSeries("商家A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.2,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: true,
			}),
		)
	return line
}

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line-多线",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)

	line.SetXAxis(nameItems).
		AddSeries("商家 A", generateLineItems()).
		AddSeries("商家 B", generateLineItems()).
		AddSeries("商家 C", generateLineItems()).
		AddSeries("商家 D", generateLineItems())
	return line
}

func lineDemo() *charts.Line {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "查询时间对比 哈希表 vs 二分查找",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "搜索时间(ns)",
			SplitLine: &opts.SplitLine{
				Show: false,
			},
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "元素数量",
		}),
	)

	line.SetXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
		AddSeries("map", generateLineItems(),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "bottom"})).
		AddSeries("slice", generateLineData([]float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131}),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top"})).
		SetSeriesOptions(
			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
				Name: "平均值",
				Type: "average",
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: true,
			}),
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show:      true,
					Formatter: "{a}: {b}",
				},
			}),
		)

	return line
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		lineShowLabel(),
		lineMarkPoint(),
		lineSplitLine(),
		lineStep(),
		lineSmooth(),
		lineArea(),
		lineSmoothArea(),
		lineMulti(),
		lineDemo(),
	)
	f, err := os.Create("line.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))
}
