package pag

import (
	"bytes"
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

func NewGrafanaClient(baseURL, apiKey string) (*GrafanaClient, error) {
	if baseURL == "" || apiKey == "" {
		return nil, errors.New("baseURL 或 apiKey 不能为空")
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("baseURL 不可用, resp.StatusCode=" + strconv.Itoa(resp.StatusCode))
	}

	return &GrafanaClient{
		httpClient: http.Client{
			Transport: authRoundTripper{
				rt:   http.DefaultTransport,
				auth: apiKey,
			},
			Timeout: time.Second * 5,
		},
		baseUrl:            baseURL,
		dashboardsEndPoint: baseURL + "/api/dashboards",
		datasourceEndPoint: baseURL + "/api/datasources",
	}, err
}

func (g *GrafanaClient) CreateDataSources(addDS grafana.AddDataSourceCommand) (grafana.DynMap, error) {
	marshal, _ := json.Marshal(addDS)
	resp, err := g.httpClient.Post(g.datasourceEndPoint, "", bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return nil, err
	}

	res := make(grafana.DynMap)
	return res, json.NewDecoder(resp.Body).Decode(&res)
}

func (g *GrafanaClient) GetDataSources() ([]grafana.DataSource, error) {
	resp, err := g.httpClient.Get(g.datasourceEndPoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return nil, err
	}

	var dss []grafana.DataSource
	return dss, json.NewDecoder(resp.Body).Decode(&dss)
}

func (g *GrafanaClient) GetDataSourceByUID(uid string) (grafana.DtoDataSource, error) {
	resp, err := g.httpClient.Get(g.datasourceEndPoint + "/uid/" + uid)
	if err != nil {
		return grafana.DtoDataSource{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return grafana.DtoDataSource{}, err
	}

	var ds grafana.DtoDataSource
	return ds, json.NewDecoder(resp.Body).Decode(&ds)
}

func (g *GrafanaClient) DeleteDataSourceByUID(uid string) error {
	request, _ := http.NewRequest("DELETE", g.datasourceEndPoint+"/uid/"+uid, nil)
	resp, err := g.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return err
	}

	return err
}

func (g *GrafanaClient) UpdateDataSourceByUID(upDS grafana.UpdateDataSourceCommand) (grafana.DynMap, error) {

	marshal, _ := json.Marshal(upDS)
	request, _ := http.NewRequest("PUT", g.datasourceEndPoint+"/uid/"+upDS.UID, bytes.NewBuffer(marshal))
	resp, err := g.httpClient.Do(request)
	if err != nil {
		return grafana.DynMap{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return nil, err
	}

	var dm grafana.DynMap
	return dm, json.NewDecoder(resp.Body).Decode(&dm)

}

// CheckDatasourceHealthWithUID  todo
func (g *GrafanaClient) CheckDatasourceHealthWithUID(uid string) (grafana.DataSource, error) {
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

func (g *GrafanaClient) GetDashBoardByUID(uid string) (*grafana.DashboardFullWithMeta, error) {
	resp, err := g.httpClient.Get(g.dashboardsEndPoint + "/uid/" + uid)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dbs grafana.DashboardFullWithMeta
	return &dbs, json.NewDecoder(resp.Body).Decode(&dbs)
}

func (g *GrafanaClient) PostDashBoard(command grafana.SaveDashboardCommand) error {
	marshal, _ := json.Marshal(command)
	request, _ := http.NewRequest("POST", g.dashboardsEndPoint+"/db", bytes.NewBuffer(marshal))
	resp, err := g.httpClient.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		var err grafana.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return err
	}
	return err
}

func (g *GrafanaClient) DeleteDashBoard(uid string) error {
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
