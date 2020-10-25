package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func genLiquidItems(data []float32) []opts.LiquidData {
	items := make([]opts.LiquidData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LiquidData{Value: data[i]})
	}
	return items
}

func liquidBase() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-basic-example",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
			}),
		)
	return liquid
}

func liquidShowLabel() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-show-label",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
			}),
		)
	return liquid
}

func liquidOutline() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-show-outline",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
				IsShowOutline:   true,
			}),
		)
	return liquid
}

func liquidWaveAnimation() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-close-wave-animation",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: false,
				IsShowOutline:   true,
			}),
		)
	return liquid
}

func liquidDiamond() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-shape(Diamond)",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
				Shape:           "diamond",
			}),
		)
	return liquid
}

func liquidPin() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-shape(Pin)",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
				Shape:           "pin",
			}),
		)
	return liquid
}

func liquidArrow() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-shape(Arrow)",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
				Shape:           "arrow",
			}),
		)
	return liquid
}

func liquidTriangle() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Liquid-shape(Triangle)",
		}),
	)

	liquid.AddSeries("liquid", genLiquidItems([]float32{0.3, 0.4, 0.5})).
		SetSeriesOptions(
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
				Shape:           "triangle",
			}),
		)
	return liquid
}

type LiquidExamples struct{}

func (LiquidExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		liquidBase(),
		liquidShowLabel(),
		liquidOutline(),
		liquidWaveAnimation(),
		liquidDiamond(),
		liquidPin(),
		liquidArrow(),
		liquidTriangle(),
	)

	f, err := os.Create("examples/html/liquid.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
