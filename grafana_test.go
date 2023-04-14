package pag

import (
	"PAG/pkg/grafana"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDataSource(t *testing.T) {
	authKey := "eyJrIjoiT3hacE5YclNSa29BWnJsemU5TjRndnZTSVZzWWdSY00iLCJuIjoiZGF3IiwiaWQiOj"
	client, _ := NewGrafanaClient(authKey, "http://127.0.0.1:3000")

	t.Run("create datasource", func(t *testing.T) {

		DS := grafana.AddDataSourceCommand{
			Name:      "Prometheus",
			Type:      "prometheus",
			Access:    "proxy",
			URL:       "http://127.0.0.1:9090",
			IsDefault: true,
			UID:       "6fDMIrLVz",
			OrgID:     1,
		}
		addDataSources, err := client.CreateDataSources(DS)
		assert.NoError(t, err)
		assert.NotEmpty(t, addDataSources)
	})

	t.Run("get all datasource", func(t *testing.T) {
		sources, err := client.GetDataSources()
		assert.NoError(t, err)
		assert.NotEmpty(t, sources)
	})

	t.Run("get datasouces by uid", func(t *testing.T) {
		dataSource, err := client.GetDataSourceByUID("6fDMIrLVz")
		assert.NoError(t, err)
		assert.NotEmpty(t, dataSource)
	})

	t.Run("delete datasouce by uid", func(t *testing.T) {
		err := client.DeleteDataSourceByUID("6fDMIrLVz")
		assert.NoError(t, err)
	})

	t.Run("Check Datasource Health", func(t *testing.T) {
		dsh, err := client.CheckDatasourceHealthWithUID("6fDMIrLVz")
		assert.NoError(t, err)
		log.Println(dsh)
	})
}

func TestDashBoard(t *testing.T) {
	authKey := "eyJrIjoiT3hacE5YclNSa29BWnJsemU5TjRndnZTSVZzWWdSY00iLCJuIjoiZGF3IiwiaWQiOj"
	client, _ := NewGrafanaClient(authKey, "http://127.0.0.1:3000")
	panelUID := "swtQMN-Vz"

	t.Run("create dashboard", func(t *testing.T) {
		sds := grafana.SaveDashboardCommand{
			//Dashboard: grafana.NewFromAny(map[string]interface{}{
			//	"title":         "Production Overview",
			//	"timezone":      "browser",
			//	"schemaVersion": 16,
			//	"version":       1,
			//	"refresh":       "30s",
			//}),
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

		err := client.PostDashBoard(sds)

		assert.NoError(t, err)
	})

	t.Run("Get DashBoard", func(t *testing.T) {
		dashboard, err := client.GetDashBoardByUID(panelUID)

		assert.NoError(t, err)
		assert.NotEmpty(t, dashboard)
	})

	t.Run("Delete DataSource", func(t *testing.T) {
		err := client.DeleteDataSourceByUID(panelUID)
		assert.NoError(t, err)
	})
}
