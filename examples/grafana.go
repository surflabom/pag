package main

import (
	"context"
	"log"

	"github.com/surflabom/pag"
	"github.com/surflabom/pag/pkg/grafana"
)

func main() {

	//使用 grafana 管理页面生成的 apiKey 来初始化 client
	congfig := grafana.Config{
		BaseURL: "http://127.0.0.1:3000",
		ApiKey:  "eyJrIjoiT3hacE5YclNSa29BWnJsemU5TjRndnZTSVZzWWdSY00iLCJuIjoiZGF3IiwiaWQiOj",
	}
	client, _ := pag.NewGrafanaClient(congfig)
	ctx := context.Background()

	//新添加一个数据源
	DS := grafana.DataSourceAdd{
		Name:      "Prometheus",
		Type:      "prometheus",
		Access:    "proxy",
		URL:       "http://127.0.0.1:9090",
		IsDefault: true,
		UID:       "6fDMIrLVz",
		OrgID:     1,
	}
	addDataSources, _ := client.CreateDataSources(ctx, DS)
	log.Println(addDataSources)

	//获取所有的数据源
	sources, _ := client.GetDataSources(ctx)
	log.Println(sources)

	//删除某个数据源
	if err := client.DeleteDataSourceByUID(ctx, "6fDMIrLVz"); err != nil {
		log.Println(err)
	}

	//新建一个 DashBoard
	panelUID := "swtQMN-Vz"
	sds := grafana.SaveDashboardCommand{

		Dashboard: map[string]interface{}{
			"title":         "zyl3",
			"timezone":      "browser",
			"schemaVersion": 16,
			"version":       1,
			"refresh":       "5s",
			"panels": []grafana.Panels{
				{
					ID: 1, // 每一个panel 的唯一ID
					Datasource: grafana.Datasource{
						Type: "prometheus",
						UID:  panelUID,
					},
					FieldConfig: grafana.FieldConfig{
						Defaults: grafana.Defaults{
							Color: grafana.Color{Mode: "palette-classic"},
							Custom: grafana.Custom{
								AxisCenteredZero: false,
								AxisColorMode:    "text",
								AxisLabel:        "",
								AxisPlacement:    "auto",
								BarAlignment:     0,
								DrawStyle:        "line",
								FillOpacity:      0,
								GradientMode:     "none",
								HideFrom: grafana.HideFrom{
									Legend:  false,
									Tooltip: false,
									Viz:     false,
								},
								LineInterpolation: "linear",
								LineWidth:         1,
								PointSize:         5,
								ScaleDistribution: grafana.ScaleDistribution{Type: "linear"},
								ShowPoints:        "auto",
								SpanNulls:         false,
								Stacking: grafana.Stacking{
									Group: "A",
									Mode:  "none",
								},
								ThresholdsStyle: grafana.ThresholdsStyle{Mode: "off"},
							},
							Mappings: nil,
							Thresholds: grafana.Thresholds{Mode: "absolute", Steps: []grafana.Steps{
								{Color: "blue", Value: nil},
								{Color: "red", Value: "80"},
							}},
						},
						Overrides: nil,
					},
					GridPos: grafana.GridPos{
						H: 9,
						W: 12,
						X: 0,
						Y: 0,
					},
					Options: grafana.Options{
						Legend: grafana.Legend{
							Calcs:       nil,
							DisplayMode: "list",
							Placement:   "bottom",
							ShowLegend:  true,
						},
						Tooltip: grafana.Tooltip{
							Mode: "single",
							Sort: "none",
						},
					},
					Targets: []grafana.Targets{
						{
							Datasource: grafana.Datasource{
								Type: "prometheus",
								UID:  "swtQMN-Vz",
							},
							EditorMode:   "builder",
							Expr:         "go_goroutines",
							LegendFormat: "__auto",
							Range:        true,
							RefID:        "A",
						},
					},
					Type:  "timeseries",
					Title: "Panel Title",
				},
				{
					ID: 3, // 每一个panel 的唯一ID
					Datasource: grafana.Datasource{
						Type: "prometheus",
						UID:  "swtQMN-Vz",
					},
					FieldConfig: grafana.FieldConfig{
						Defaults: grafana.Defaults{
							Color: grafana.Color{Mode: "palette-classic"},
							Custom: grafana.Custom{
								AxisCenteredZero: false,
								AxisColorMode:    "text",
								AxisLabel:        "",
								AxisPlacement:    "auto",
								BarAlignment:     0,
								DrawStyle:        "line",
								FillOpacity:      0,
								GradientMode:     "none",
								HideFrom: grafana.HideFrom{
									Legend:  false,
									Tooltip: false,
									Viz:     false,
								},
								LineInterpolation: "linear",
								LineWidth:         1,
								PointSize:         5,
								ScaleDistribution: grafana.ScaleDistribution{Type: "linear"},
								ShowPoints:        "auto",
								SpanNulls:         false,
								Stacking: grafana.Stacking{
									Group: "A",
									Mode:  "none",
								},
								ThresholdsStyle: grafana.ThresholdsStyle{Mode: "off"},
							},
							Mappings: nil,
							Thresholds: grafana.Thresholds{Mode: "absolute", Steps: []grafana.Steps{
								{Color: "blue", Value: nil},
								{Color: "red", Value: "80"},
							}},
						},
						Overrides: nil,
					},
					GridPos: grafana.GridPos{
						H: 9,
						W: 12,
						X: 0,
						Y: 0,
					},
					Options: grafana.Options{
						Legend: grafana.Legend{
							Calcs:       nil,
							DisplayMode: "list",
							Placement:   "bottom",
							ShowLegend:  true,
						},
						Tooltip: grafana.Tooltip{
							Mode: "single",
							Sort: "none",
						},
					},
					Targets: []grafana.Targets{
						{
							Datasource: grafana.Datasource{
								Type: "prometheus",
								UID:  "swtQMN-Vz",
							},
							EditorMode:   "builder",
							Expr:         "go_memstats_heap_objects",
							LegendFormat: "__auto",
							Range:        true,
							RefID:        "A",
						},
					},
					Type:  "timeseries",
					Title: "go_memstats_heap_objects",
				},
			},
		},
		Overwrite: true, //如果为 false 则不会更新 dashboard
		//FolderUID: "l3KqBxCMz",
		//IsFolder:  false,
		Message: "Made changes to xyz",
	}

	if err := client.PostDashBoard(ctx, sds); err != nil {
		log.Println(err)
	}

	//通过 DashBoard UID 来获得对应的 DashBoard 详细信息
	DashBoard, err := client.GetDashBoardByUID(ctx, panelUID)
	if err != nil {
		log.Println(err)
	}

	log.Println(DashBoard)

}
