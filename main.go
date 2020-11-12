package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"os"
)

type teamdata struct {
	Week string
	Data map[string]int
}

var (
	valsc = []teamdata{
		{Week: "Week 1", Data: map[string]int{"Person-1": 2, "Person-2": 2, "Person-3": 2, "Person-4": 2, "P-5": 2, "P-6": 2}},
		{Week: "Week 2", Data: map[string]int{"Person-1": 0, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 3", Data: map[string]int{"Person-1": 5, "Person-2": 1, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 1}},
		{Week: "Week 4", Data: map[string]int{"Person-1": 1, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 5", Data: map[string]int{"Person-1": 0, "Person-2": 0, "Person-3": 2, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 6", Data: map[string]int{"Person-1": 3, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 7", Data: map[string]int{"Person-1": 5, "Person-2": 5, "Person-3": 0, "Person-4": 1, "P-5": 0, "P-6": 0}},
		{Week: "Week 8", Data: map[string]int{"Person-1": 2, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 9", Data: map[string]int{"Person-1": 4, "Person-2": 0, "Person-3": 1, "Person-4": 0, "P-5": 0, "P-6": 2}},
		{Week: "Week 10", Data: map[string]int{"Person-1": 3, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 11", Data: map[string]int{"Person-1": 3, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 12", Data: map[string]int{"Person-1": 0, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
	}
	valst = []teamdata{
		{Week: "Week 1", Data: map[string]int{"Person-1": 100, "Person-2": 100, "Person-3": 100, "Person-4": 100, "P-5": 100, "P-6": 100}},
		{Week: "Week 2", Data: map[string]int{"Person-1": 100, "Person-2": 100, "Person-3": 100, "Person-4": 100, "P-5": 100, "P-6": 75}},
		{Week: "Week 3", Data: map[string]int{"Person-1": 100, "Person-2": 50, "Person-3": 50, "Person-4": 50, "P-5": 50, "P-6": 50}},
		{Week: "Week 4", Data: map[string]int{"Person-1": 100, "Person-2": 100, "Person-3": 50, "Person-4": 100, "P-5": 50, "P-6": 100}},
		{Week: "Week 5", Data: map[string]int{"Person-1": 100, "Person-2": 100, "Person-3": 0, "Person-4": 100, "P-5": 0, "P-6": 50}},
		{Week: "Week 6", Data: map[string]int{"Person-1": 100, "Person-2": 100, "Person-3": 100, "Person-4": 100, "P-5": 0, "P-6": 50}},
		{Week: "Week 7", Data: map[string]int{"Person-1": 100, "Person-2": 50, "Person-3": 0, "Person-4": 100, "P-5": 0, "P-6": 100}},
		{Week: "Week 8", Data: map[string]int{"Person-1": 100, "Person-2": 50, "Person-3": 100, "Person-4": 50, "P-5": 50, "P-6": 100}},
		{Week: "Week 9", Data: map[string]int{"Person-1": 100, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 10", Data: map[string]int{"Person-1": 70, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 11", Data: map[string]int{"Person-1": 100, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
		{Week: "Week 12", Data: map[string]int{"Person-1": 0, "Person-2": 0, "Person-3": 0, "Person-4": 0, "P-5": 0, "P-6": 0}},
	}
	order = []string{"Person-1", "Person-2", "Person-3", "Person-4", "P-5", "P-6"}
)

func generateLineItems(who string, arr []teamdata) []opts.LineData {
	res := make([]opts.LineData, 0)
	for i := 0; i < len(arr); i++ {
		res = append(res, opts.LineData{Value: arr[i].Data[who]})
	}
	return res
}

func getweeks(arr []teamdata) []string { //should be same for both
	res := []string{}
	for _, val := range arr {
		res = append(res, val.Week)
	}
	return res
}

func lineSmooth(which bool) *charts.Line {
	line, title := charts.NewLine(), ""
	if which {
		title = "Git Commit Graph"
	} else {
		title = "Task Graph"
		line.SetGlobalOptions(
			charts.WithYAxisOpts(opts.YAxis{
				AxisLabel: &opts.Label{Formatter: "{value}%"},
			}),
		)
	}
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			PageTitle: "Team X Stats",
			Width:     "800px",
			Height:    "500px",
			Theme:     "dark", //switch to available themes listed here or create your own from the templates, https://github.com/go-echarts/go-echarts-assets/tree/master/assets/themes
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger:   "axis",
			TriggerOn: "click",
			Show:      true,
			Formatter: "Date: {b}<br/>Value: {c}",
		}),
		charts.WithToolboxOpts(opts.Toolbox{ //toolbox postioning and options
			Show:   true,
			Orient: "vertical",
			Left:   "1%",
			Top:    "5%",
			Feature: &opts.ToolBoxFeature{
				&opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type: "png",
					Name: "team_graphs",
					Title: "Save",
				},
				nil,
				&opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "View Data",
					Lang:  []string{"Data View", "Exit"},
				},
				nil,
			},
		}),
	)

	var whichlist []teamdata
	if which == true {
		whichlist = valsc
	} else {
		whichlist = valst
	}

	line.SetXAxis(getweeks(whichlist))

	for _, k := range order {
		line.AddSeries(k, generateLineItems(k, whichlist))
	}
	line.SetSeriesOptions(charts.WithLineChartOpts(
		opts.LineChart{
			Smooth: true,
		}),
	)

	return line
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		lineSmooth(true),
		lineSmooth(false),
	)
	f, err := os.Create("./line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
