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
	itemCntPie = 4
	seasons    = []string{"Spring", "Summer", "Autumn ", "Winter"}
)

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

func pieBase() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic pie example"}),
	)

	pie.AddSeries("pie", generatePieItems())
	return pie
}

func pieShowLabel() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "label options"}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
		)
	return pie
}

func pieRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Radius style"}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	return pie
}

func pieRoseArea() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Rose(Area)",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"40%", "75%"},
				RoseType: "area",
			}),
		)
	return pie
}

func pieRoseRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Rose(Radius)",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "radius",
			}),
		)
	return pie
}

func pieRoseAreaRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Rose(Area/Radius)",
		}),
	)

	pie.AddSeries("area", generatePieItems()).
		SetSeriesOptions(
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "area",
				Center:   []string{"25%", "50%"},
			}),
		)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "radius",
				Center:   []string{"75%", "50%"},
			}),
		)
	return pie
}

func pieInPie() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "pie in pie",
		}),
	)

	pie.AddSeries("area", generatePieItems(),
		charts.WithLabelOpts(opts.Label{
			Show:      opts.Bool(true),
			Formatter: "{b}: {c}",
		}),
		charts.WithPieChartOpts(opts.PieChart{
			Radius:   []string{"50%", "55%"},
			RoseType: "area",
		}),
	)

	pie.AddSeries("radius", generatePieItems(),
		charts.WithPieChartOpts(opts.PieChart{
			Radius:   []string{"0%", "45%"},
			RoseType: "radius",
		}),
	)
	return pie
}

func pieWithDispatchAction() *charts.Pie {
	const actionWithEchartsInstance = `
		let currentIndex = -1;
		setInterval(function() {
		  const myChart = %MY_ECHARTS%;
		  var dataLen = myChart.getOption().series[0].data.length;
		  myChart.dispatchAction({
			type: 'downplay',
			seriesIndex: 0,
			dataIndex: currentIndex
		  });
		  currentIndex = (currentIndex + 1) % dataLen;
		  myChart.dispatchAction({
			type: 'highlight',
			seriesIndex: 0,
			dataIndex: currentIndex
		  });
		  myChart.dispatchAction({
			type: 'showTip',
			seriesIndex: 0,
			dataIndex: currentIndex
		  });
		}, 1000);
`

	pie := charts.NewPie()
	pie.AddJSFuncStrs(actionWithEchartsInstance)
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "dispatchAction pie",
			Right: "40%",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger:   "item",
			Formatter: "{a} <br/>{b} : {c} ({d}%)",
		}),
		charts.WithLegendOpts(opts.Legend{
			Left:   "left",
			Orient: "vertical",
		}),
	)

	pie.AddSeries("pie action", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"55%"},
				Center: []string{"50%", "60%"},
			}),

			charts.WithEmphasisOpts(opts.Emphasis{
				ItemStyle: &opts.ItemStyle{
					ShadowBlur:    10,
					ShadowOffsetX: 0,
					ShadowColor:   "rgba(0, 0, 0, 0.5)",
				},
			}),
		)

	return pie

}

type PieExamples struct{}

func (PieExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		pieBase(),
		pieShowLabel(),
		pieRadius(),
		pieRoseArea(),
		pieRoseRadius(),
		pieRoseAreaRadius(),
		pieInPie(),
		pieWithDispatchAction(),
	)
	f, err := os.Create("examples/html/pie.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
