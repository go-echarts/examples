package examples

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var TreeMap = []opts.TreeMapNode{
	{
		Name:     "d1",
		Children: []opts.TreeMapNode{{Name: "f1", Value: 1000}},
	},
	{
		Name:  "d2",
		Children: []opts.TreeMapNode{
			{Name: "f1", Value: 100},
			{Name: "f2", Value: 300},
			{Name: "f3", Value: 200},
		},
	},
	{
		Name:  "d3",
		// Children populated later.
	},
	{
		Name:  "f1",
		Value: 450,
	},
}

var ToolTipFormatter = `
function (info) {
	var formatUtil = echarts.format;
	var value = info.value;
	var treePathInfo = info.treePathInfo;
	var treePath = [];
	for (var i = 1; i < treePathInfo.length; i++) {
		treePath.push(treePathInfo[i].name);
	}
	return ['<div class="tooltip-title">' + formatUtil.encodeHTML(treePath.join('/')) + '</div>',
		'Disk Usage: ' + formatUtil.addCommas(value) + ' KB',
		].join('');
}
`

func treeMapBase() *charts.TreeMap {
	graph := charts.NewTreeMap()
	graph.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeMacarons}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Basic treemap example",
			Subtitle: "File system usage",
			Left:     "center",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Formatter: opts.FuncOpts(ToolTipFormatter),
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:   true,
			Orient: "horizontal",
			Left:   "right",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show: true, Title: "Save as image"},
				Restore: &opts.ToolBoxFeatureRestore{
					Show: true, Title: "Reset"},
			}}),
	)
	// Populate "d3" node with large number of small-sized files.
	d3Index := 2
	d3NumFiles := 40
	TreeMap[d3Index].Children = make([]opts.TreeMapNode, d3NumFiles)
	for i := range TreeMap[d3Index].Children {
		TreeMap[d3Index].Children[i] = opts.TreeMapNode{
			Name:  fmt.Sprintf("f%v", i),
			Value: 5 + rand.Intn(15-5),
		}
	}
	// Add initialized data to graph.
	graph.AddSeries("Root FS", TreeMap).
		SetSeriesOptions(
			charts.WithTreeMapOpts(
				opts.TreeMapChart{
					Animation:  true,
					Roam:       true,
					UpperLabel: &opts.UpperLabel{Show: true},
					Levels: &[]opts.TreeMapLevel{
						{ // Series
							ItemStyle: &opts.ItemStyle{
								BorderColor: "#777",
								BorderWidth: 1,
								GapWidth:    1},
							UpperLabel: &opts.UpperLabel{Show: false},
						},
						{ // Level
							ItemStyle: &opts.ItemStyle{
								BorderColor: "#666",
								BorderWidth: 2,
								GapWidth:    1},
							Emphasis: &opts.Emphasis{
								ItemStyle: &opts.ItemStyle{BorderColor: "#555"},
							},
						},
						{ // Node
							ColorSaturation: []float32{0.35, 0.5},
							ItemStyle: &opts.ItemStyle{
								GapWidth:              1,
								BorderWidth:           0,
								BorderColorSaturation: 0.6,
							},
						},
					},
				},
			),
			charts.WithItemStyleOpts(opts.ItemStyle{BorderColor: "#fff"}),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "inside", Color: "White"}),
		)
	return graph
}

type TreeMapExamples struct{}

func (TreeMapExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		treeMapBase(),
	)

	f, err := os.Create("examples/html/treemap.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
