package main

import (
	"io"
	"log"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
)

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genLine3dData() [][3]float64 {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{
				(1 + 0.25*math.Cos(75*float64(t))) * math.Cos(float64(t)),
				(1 + 0.25*math.Cos(75*float64(t))) * math.Sin(float64(t)),
				float64(t) + 2.0*math.Sin(75.0*float64(t)),
			},
		)
	}
	return data
}

func line3DBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line3D-example",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),
	)

	line3d.AddZAxis("line3D", genLine3dData())
	return line3d
}

func line3DAutoRotate() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Line3D-auto-rotate",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),

		charts.WithGrid3DOpts(opts.Grid3D{
			ViewControl: opts.ViewControl{
				AutoRotate: true,
			},
		}),
	)
	line3d.AddZAxis("line3D", genLine3dData())
	return line3d
}

func main() {

	page := components.NewPage()
	page.AddCharts(
		line3DBase(),
		line3DAutoRotate(),
	)

	f, err := os.Create("line3D.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))
}

