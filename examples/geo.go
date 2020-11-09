package examples

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var geoData = []opts.GeoData{
	{Name: "北京", Value: []float64{116.40, 39.90, float64(rand.Intn(100))}},
	{Name: "上海", Value: []float64{121.47, 31.23, float64(rand.Intn(100))}},
	{Name: "重庆", Value: []float64{106.55, 29.56, float64(rand.Intn(100))}},
	{Name: "武汉", Value: []float64{114.31, 30.52, float64(rand.Intn(100))}},
	{Name: "台湾", Value: []float64{121.30, 25.03, float64(rand.Intn(100))}},
	{Name: "香港", Value: []float64{114.17, 22.28, float64(rand.Intn(100))}},
}

var guangdongData = []opts.GeoData{
	{Name: "汕头", Value: []float64{116.69, 23.39, float64(rand.Intn(100))}},
	{Name: "深圳", Value: []float64{114.07, 22.62, float64(rand.Intn(100))}},
	{Name: "广州", Value: []float64{113.23, 23.16, float64(rand.Intn(100))}},
}

func geoBase() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic geo example"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "china",
			ItemStyle: &opts.ItemStyle{Color: "#006666"},
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
		charts.WithTitleOpts(opts.Title{Title: "Guangdong province"}),
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

type GeoExamples struct{}

func (GeoExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		geoBase(),
		geoGuangdong(),
	)

	f, err := os.Create("examples/html/geo.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
