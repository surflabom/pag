package pag

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"pag/pkg/grafana"
	"strconv"
	"time"
)

type GrafanaClient struct {
	httpClient         http.Client
	baseUrl            string
	dashboardsEndPoint string
	datasourceEndPoint string
}

type authRoundTripper struct {
	rt   http.RoundTripper
	auth string
}

func (art authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+art.auth)
	req.Header.Set("Content-Type", "application/json")
	return art.rt.RoundTrip(req)
}

func NewGrafanaClient(config grafana.Config) (*GrafanaClient, error) {
	if config.BaseURL == "" || config.ApiKey == "" {
		return nil, errors.New("baseURL or apiKey cannot be empty")
	}

	resp, err := http.Get(config.BaseURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("BaseURL is not accessible, responses statusCode=" + strconv.Itoa(resp.StatusCode))
	}

	return &GrafanaClient{
		httpClient: http.Client{
			Transport: authRoundTripper{
				rt:   http.DefaultTransport,
				auth: config.ApiKey,
			},
			Timeout: time.Second * 5,
		},
		baseUrl:            config.BaseURL,
		dashboardsEndPoint: config.BaseURL + "/api/dashboards",
		datasourceEndPoint: config.BaseURL + "/api/datasources",
	}, err
}

func (g *GrafanaClient) GetReqWithContext(ctx context.Context, url string, i interface{}) error {
	return g.NewReqWithContext(ctx, "GET", url, nil, i)
}

func (g *GrafanaClient) NewReqWithContext(ctx context.Context, method string, url string, body io.Reader, i interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := g.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // 关闭响应体

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(i)
}

func (g *GrafanaClient) CreateDataSources(ctx context.Context, addDS grafana.AddDataSourceCommand) (grafana.DynMap, error) {
	marshal, _ := json.Marshal(addDS)
	res := make(grafana.DynMap)
	err := g.NewReqWithContext(ctx, "POST", g.datasourceEndPoint, bytes.NewBuffer(marshal), res)
	return res, err

}

func (g *GrafanaClient) GetDataSources(ctx context.Context) ([]grafana.DataSource, error) {
	var dss []grafana.DataSource
	err := g.GetReqWithContext(ctx, g.datasourceEndPoint, &dss)
	return dss, err
}

func (g *GrafanaClient) GetDataSourceByUID(ctx context.Context, uid string) (grafana.DtoDataSource, error) {
	var ds grafana.DtoDataSource
	err := g.GetReqWithContext(ctx, g.datasourceEndPoint+"/uid/"+uid, &ds)
	return ds, err
}

func (g *GrafanaClient) DeleteDataSourceByUID(ctx context.Context, uid string) error {
	err := g.NewReqWithContext(ctx, "DELETE", g.datasourceEndPoint+"/uid/"+uid, nil, nil)
	if err != io.EOF {
		return err
	}
	return nil
}

func (g *GrafanaClient) UpdateDataSourceByUID(ctx context.Context, upDS grafana.UpdateDataSourceCommand) (grafana.DynMap, error) {
	var dm grafana.DynMap
	marshal, _ := json.Marshal(upDS)
	err := g.NewReqWithContext(ctx, "PUT", g.datasourceEndPoint+"/uid/"+upDS.UID, bytes.NewBuffer(marshal), &dm)
	return dm, err
}

// CheckDatasourceHealthWithUID  todo 部分插件不支持
func (g *GrafanaClient) CheckDatasourceHealthWithUID(ctx context.Context, uid string) (grafana.DataSource, error) {
	//g.GetReqWithContext(ctx,g.datasourceEndPoint + "/uid/" + uid + "/health",nil)

	resp, err := g.httpClient.Get(g.datasourceEndPoint + "/uid/" + uid + "/health")
	if err != nil {
		return grafana.DataSource{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return grafana.DataSource{}, err
	}

	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))

	return grafana.DataSource{}, nil
	//var ds DataSource
	//return ds, json.NewDecoder(resp.Body).Decode(&ds)
}

func (g *GrafanaClient) GetDashBoardByUID(ctx context.Context, uid string) (*grafana.DashboardFullWithMeta, error) {
	var dbs grafana.DashboardFullWithMeta
	err := g.GetReqWithContext(ctx, g.dashboardsEndPoint+"/uid/"+uid, &dbs)
	return &dbs, err

}

// todo
func (g *GrafanaClient) PostDashBoard(ctx context.Context, command grafana.SaveDashboardCommand) error {
	marshal, _ := json.Marshal(command)
	err := g.NewReqWithContext(ctx, "POST", g.dashboardsEndPoint+"/db", bytes.NewBuffer(marshal), nil)
	return err

	//marshal, _ := json.Marshal(command)
	//request, _ := http.NewRequest("POST", g.dashboardsEndPoint+"/db", bytes.NewBuffer(marshal))
	//resp, err := g.httpClient.Do(request)
	//if err != nil {
	//	return err
	//}
	//
	//if resp.StatusCode != 200 {
	//	var err grafana.ResponseError
	//	json.NewDecoder(resp.Body).Decode(&err)
	//	return err
	//}
	//return err
}

func (g *GrafanaClient) DeleteDashBoard(ctx context.Context, uid string) error {
	request, _ := http.NewRequest("DELETE", g.dashboardsEndPoint+"/uid/"+uid, nil)
	resp, err := g.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	all, _ := io.ReadAll(resp.Body)
	log.Println(string(all))
	return nil
}
