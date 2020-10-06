package main

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
)

var geoData = []opts.GeoData{
	{"北京", []float64{116.40, 39.90, float64(rand.Intn(100))}},
	{"上海", []float64{121.47, 31.23, float64(rand.Intn(100))}},
	{"重庆", []float64{106.55, 29.56, float64(rand.Intn(100))}},
	{"武汉", []float64{114.31, 30.52, float64(rand.Intn(100))}},
	{"台湾", []float64{121.30, 25.03, float64(rand.Intn(100))}},
	{"香港", []float64{114.17, 22.28, float64(rand.Intn(100))}},
}

var guangdongData = []opts.GeoData{
	{"汕头", []float64{116.69, 23.39, float64(rand.Intn(100))}},
	{"深圳", []float64{114.07, 22.62, float64(rand.Intn(100))}},
	{"广州", []float64{113.23, 23.16, float64(rand.Intn(100))}},
}

func geoBase() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "geo-basic-example",
		}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map: "china",
			ItemStyle: opts.ItemStyle{
				Color: "#006666",
			},
		}),
	)

	geo.AddSeries("geo", types.ChartEffectScatter, geoData,
		charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     6,
			BrushType: "stroke",
		}),
	)
	return geo
}

func geoGuangdong() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "geo-Guangdong",
		}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map: "广东",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)

	geo.AddSeries("geo", types.ChartScatter, guangdongData)
	return geo
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		geoBase(),
		geoGuangdong(),
	)

	f, err := os.Create("html/geo.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(os.Stdout, f))
}
