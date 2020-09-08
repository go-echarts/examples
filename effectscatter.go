package main

import (
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
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
		charts.WithTitleOpts(opts.Title{
			Title: "EffectScatter-example",
		}),
	)

	es.SetXAxis(player).
		AddSeries("Dunk", generateEffectScatterItems())
	return es
}

func esEffectStyle() *charts.EffectScatter {
	es := charts.NewEffectScatter()
	es.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "EffectScatter-wave-style",
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

func main() {
	page := components.NewPage()
	page.AddCharts(
		esBase(),
		esEffectStyle(),
	)

	f, err := os.Create("effectscatter.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))

}
