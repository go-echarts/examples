package main

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts-examples/examples"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	examplers := []examples.Exampler{
		examples.BarExamples{},
		examples.Bar3dExamples{},
		examples.BoxplotExamples{},
		examples.EffectscatterExamples{},
		examples.FunnelExamples{},
		examples.FunnelExamples{},
		examples.GaugeExamples{},
		examples.GeoExamples{},
		examples.GraphExamples{},
		examples.HeatmapExamples{},
		examples.KlineExamples{},
		examples.LineExamples{},
		examples.Line3dExamples{},
		examples.LiquidExamples{},
		examples.MapExamples{},
		examples.PageCenterLayoutExamples{},
		examples.PageFlexLayoutExamples{},
		examples.PageNoneLayoutExamples{},
		examples.ParallelExamples{},
		examples.PieExamples{},
		examples.RadarExamples{},
		examples.CustomizeExamples{},
		examples.SankeyExamples{},
		examples.ScatterExamples{},
		examples.Scatter3dExamples{},
		examples.Surface3dExamples{},
		examples.ThemeriverExamples{},
		examples.ThemeExamples{},
		examples.WordcloudExamples{},
	}

	for _, e := range examplers {
		e.Examples()
	}

	fs := http.FileServer(http.Dir("examples/html"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}
