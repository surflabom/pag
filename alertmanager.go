package pag

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/surflabom/pag/pkg/alertmanger"
)

type AlertmanagerClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewAlertmanagerClient(config alertmanger.Config) (*AlertmanagerClient, error) {
	if config.BaseURL == "" {
		return nil, errors.New("BaseURL  cannot be empty")
	}

	resp, err := http.Get(config.BaseURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("BaseURL is not accessible, responses statusCode=" + strconv.Itoa(resp.StatusCode))
	}

	return &AlertmanagerClient{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		baseUrl: config.BaseURL + "/api/v2",
	}, nil
}

func (a *AlertmanagerClient) GetReqWithContext(ctx context.Context, url string, i interface{}) error {
	return a.NewReqWithContext(ctx, "GET", url, nil, i)
}

func (a *AlertmanagerClient) NewReqWithContext(ctx context.Context, method string, url string, body io.Reader, i interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // 关闭响应体

	if resp.StatusCode != 200 {
		var err alertmanger.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(i)
}

func (a *AlertmanagerClient) Status(ctx context.Context) (alertmanger.Status, error) {
	var status alertmanger.Status
	err := a.GetReqWithContext(ctx, a.baseUrl+"/status", &status)
	return status, err
}

func (a *AlertmanagerClient) GetReceivers(ctx context.Context) ([]alertmanger.Receiver, error) {
	var receivers []alertmanger.Receiver
	err := a.GetReqWithContext(ctx, a.baseUrl+"/receivers", receivers)
	return receivers, err
}

// GetAllSilences  从API获取所有静默列表信息
func (a *AlertmanagerClient) GetAllSilences(ctx context.Context) ([]alertmanger.Silence, error) {
	var silences []alertmanger.Silence
	err := a.GetReqWithContext(ctx, a.baseUrl+"/silences", &silences)
	return silences, err

}

// NewSilences 创建新的静默信息
func (a *AlertmanagerClient) NewSilences(ctx context.Context, silence alertmanger.Silence) (alertmanger.SilenceID, error) {
	marshal, _ := json.Marshal(silence)
	var silenceID alertmanger.SilenceID
	err := a.NewReqWithContext(ctx, "POST", a.baseUrl+"/silences", bytes.NewBuffer(marshal), &silenceID)
	return silenceID, err
}

// GetSilenceByID 根据ID获取特定的静默信息
func (a *AlertmanagerClient) GetSilenceByID(ctx context.Context, id string) (alertmanger.Silence, error) {
	var silence alertmanger.Silence
	err := a.GetReqWithContext(ctx, a.baseUrl+"/silence/"+id, &silence)
	return silence, err

}

// DeleteSilenceByID 根据ID删除特定的静默信息
func (a *AlertmanagerClient) DeleteSilenceByID(ctx context.Context, id string) error {
	if err := a.NewReqWithContext(ctx, "DELETE", a.baseUrl+"/silence/"+id, nil, nil); err != io.EOF {
		return err
	}
	return nil

}

// GetAlerts 从API获取所有告警列表信息
func (a *AlertmanagerClient) GetAlerts(ctx context.Context) ([]alertmanger.GettableAlert, error) {
	var alerts []alertmanger.GettableAlert
	err := a.GetReqWithContext(ctx, a.baseUrl+"/alerts", &alerts)
	return alerts, err
}

// NewAlert 创建新的告警信息
func (a *AlertmanagerClient) NewAlert(ctx context.Context, alert []alertmanger.Alert) error {
	marshal, _ := json.Marshal(alert)
	err := a.NewReqWithContext(ctx, "POST", a.baseUrl+"/alerts", bytes.NewBuffer(marshal), nil)
	if err != io.EOF {
		return err
	}
	return nil
}

// GetAlertGroup 从API获取所有告警分组信息
func (a *AlertmanagerClient) GetAlertGroup(ctx context.Context) ([]alertmanger.GettableAlertGroup, error) {
	var tableAlert []alertmanger.GettableAlertGroup
	err := a.GetReqWithContext(ctx, a.baseUrl+"/alerts/groups", &tableAlert)
	return tableAlert, err
}
