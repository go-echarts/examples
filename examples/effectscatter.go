package examples

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var player = []string{"Kobe", "Jordan", "Iverson", "LeBron", "Wade", "McGrady"}

func generateEffectScatterItems() []opts.EffectScatterData {
	items := make([]opts.EffectScatterData, 0)
	for i := 0; i < len(player); i++ {
		items = append(items, opts.EffectScatterData{Value: rand.Intn(100)})
	}
	return items
}

func esBase() *charts.EffectScatter {
	es := charts.NewEffectScatter()
	es.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic EffectScatter example"}),
	)

	es.SetXAxis(player).
		AddSeries("Dunk", generateEffectScatterItems())
	return es
}

func esEffectStyle() *charts.EffectScatter {
	es := charts.NewEffectScatter()
	es.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "wave style",
		}),
	)

	es.SetXAxis(player).
		AddSeries("Dunk", generateEffectScatterItems(),
			charts.WithRippleEffectOpts(opts.RippleEffect{
				Period:    4,
				Scale:     10,
				BrushType: "stroke",
			})).
		AddSeries("Shoot", generateEffectScatterItems(),
			charts.WithRippleEffectOpts(opts.RippleEffect{
				Period:    3,
				Scale:     6,
				BrushType: "fill",
			}),
		)
	return es
}

type EffectscatterExamples struct{}

func (EffectscatterExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		esBase(),
		esEffectStyle(),
	)

	f, err := os.Create("examples/html/effectscatter.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
