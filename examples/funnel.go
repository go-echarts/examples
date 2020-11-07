package examples

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var dimensions = []string{"Visit", "Add", "Order", "Payment", "Deal"}

func genFunnelKvItems() []opts.FunnelData {
	items := make([]opts.FunnelData, 0)
	for i := 0; i < len(dimensions); i++ {
		items = append(items, opts.FunnelData{Name: dimensions[i], Value: rand.Intn(50)})
	}
	return items
}
func funnelBase() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic funnel example"}),
	)

	funnel.AddSeries("Analytics", genFunnelKvItems())
	return funnel
}

// TODO: check the different from echarts side
func funnelShowLabel() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "show label"}),
	)

	funnel.AddSeries("Analytics", genFunnelKvItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:     true,
				Position: "left",
			},
		))
	return funnel
}

type FunnelExamples struct{}

func (FunnelExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		funnelBase(),
		funnelShowLabel(),
	)

	f, err := os.Create("examples/html/funnel.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
