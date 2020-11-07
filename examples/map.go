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
	baseMapData = []opts.MapData{
		{Name: "北京", Value: float64(rand.Intn(150))},
		{Name: "上海", Value: float64(rand.Intn(150))},
		{Name: "广东", Value: float64(rand.Intn(150))},
		{Name: "辽宁", Value: float64(rand.Intn(150))},
		{Name: "山东", Value: float64(rand.Intn(150))},
		{Name: "山西", Value: float64(rand.Intn(150))},
		{Name: "陕西", Value: float64(rand.Intn(150))},
		{Name: "新疆", Value: float64(rand.Intn(150))},
		{Name: "内蒙古", Value: float64(rand.Intn(150))},
	}

	guangdongMapData = map[string]float64{
		"深圳市": float64(rand.Intn(150)),
		"广州市": float64(rand.Intn(150)),
		"湛江市": float64(rand.Intn(150)),
		"汕头市": float64(rand.Intn(150)),
		"东莞市": float64(rand.Intn(150)),
		"佛山市": float64(rand.Intn(150)),
		"云浮市": float64(rand.Intn(150)),
		"肇庆市": float64(rand.Intn(150)),
		"梅州市": float64(rand.Intn(150)),
	}
)

func generateMapData(data map[string]float64) (items []opts.MapData) {
	items = make([]opts.MapData, 0)
	for k, v := range data {
		items = append(items, opts.MapData{Name: k, Value: v})
	}
	return
}

func mapBase() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic map example"}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "show label"}),
	)

	mc.AddSeries("map", baseMapData).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return mc
}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "VisualMap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("广东")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Guangdong province",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}},
		}),
	)

	mc.AddSeries("map", generateMapData(guangdongMapData))
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "macarons",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Map-theme",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        150,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

type MapExamples struct{}

func (MapExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapRegion(),
		mapTheme(),
	)

	f, err := os.Create("examples/html/map.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
