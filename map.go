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

var (
	baseMapData = map[string]float32{
		"北京":  float32(rand.Intn(150)),
		"上海":  float32(rand.Intn(150)),
		"广东":  float32(rand.Intn(150)),
		"辽宁":  float32(rand.Intn(150)),
		"山东":  float32(rand.Intn(150)),
		"山西":  float32(rand.Intn(150)),
		"陕西":  float32(rand.Intn(150)),
		"新疆":  float32(rand.Intn(150)),
		"内蒙古": float32(rand.Intn(150)),
	}

	guangdongMapData = map[string]float32{
		"深圳市": float32(rand.Intn(150)),
		"广州市": float32(rand.Intn(150)),
		"湛江市": float32(rand.Intn(150)),
		"汕头市": float32(rand.Intn(150)),
		"东莞市": float32(rand.Intn(150)),
		"佛山市": float32(rand.Intn(150)),
		"云浮市": float32(rand.Intn(150)),
		"肇庆市": float32(rand.Intn(150)),
		"梅州市": float32(rand.Intn(150)),
	}
)

func generateMapData(data map[string]float32) (items []opts.MapData) {
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
		charts.WithTitleOpts(opts.Title{
			Title: "Map-example",
		}),
	)

	mc.AddSeries("map", generateMapData(baseMapData))
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Map-show-label",
		}),
	)

	mc.AddSeries("map", generateMapData(baseMapData)).
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
			Title: "Map-VisualMap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)

	mc.AddSeries("map", generateMapData(baseMapData))
	return mc
}

func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("广东")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Map-religion-Guangdong",
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

	mc.AddSeries("map", generateMapData(baseMapData))
	return mc
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapRegion(),
		mapTheme(),
	)

	f, err := os.Create("map.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))
}
