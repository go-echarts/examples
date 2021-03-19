package examples

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var TreeNodes = []*opts.TreeData{
	{
		Name: "Node1",
		Children: []*opts.TreeData{
			{
				Name: "Chield1",
			},
		},
	},
	{
		Name: "Node2",
		Children: []*opts.TreeData{
			{
				Name: "Chield1",
			},
			{
				Name: "Chield2",
			},
			{
				Name: "Chield3",
			},
		},
	},
	{
		Name:      "Node3",
		Collapsed: true,
		Children: []*opts.TreeData{
			{
				Name: "Chield1",
			},
			{
				Name: "Chield2",
			},
			{
				Name: "Chield3",
			},
		},
	},
}

var Tree = []opts.TreeData{
	{
		Name:     "Root",
		Children: TreeNodes,
	},
}

func treeBase() *charts.Tree {
	graph := charts.NewTree()
	graph.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Width: "100%", Height: "95vh"}),
		charts.WithTitleOpts(opts.Title{Title: "basic tree example"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: false}),
	)
	graph.AddSeries("tree", Tree).
		SetSeriesOptions(
			charts.WithTreeOpts(
				opts.TreeChart{
					Layout:           "orthogonal",
					Orient:           "LR",
					InitialTreeDepth: -1,
					Leaves: &opts.TreeLeaves{
						Label: &opts.Label{Show: true, Position: "right", Color: "Black"},
					},
				},
			),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}),
		)
	return graph
}

type TreeExamples struct{}

func (TreeExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		treeBase(),
	)

	f, err := os.Create("examples/html/tree.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
