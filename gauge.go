package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
)

func gaugeBase() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Gauge-example"}),
	)
	gauge.AddSeries("ProjectA", []opts.GaugeData{{Name: "Work progress", Value: rand.Intn(50)}})
	return gauge
}

func gaugeTimer() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Gauge-timer"}),
	)
	gauge.AddSeries("ProjectB", []opts.GaugeData{{Name: "Work progress", Value: rand.Intn(50)}})
	fn := fmt.Sprintf(`setInterval(function () {
			option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
			myChart_%s.setOption(option_%s, true);
		}, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
	gauge.AddJSFuncs(fn)
	return gauge
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		gaugeBase(),
		gaugeTimer(),
	)

	f, err := os.Create("guage.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))

}
