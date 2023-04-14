package pag

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"pag/pkg/alertmanger"
	"strconv"
	"time"
)

type AlertmanagerClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewAlertmanagerClient(baseURL string) (*AlertmanagerClient, error) {
	if baseURL == "" {
		return nil, errors.New("baseURL 不能为空")
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("baseURL 不可用, resp.StatusCode=" + strconv.Itoa(resp.StatusCode))
	}

	return &AlertmanagerClient{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		baseUrl: baseURL + "/api/v2",
	}, nil
}

func (a *AlertmanagerClient) Status() (alertmanger.Status, error) {
	resp, err := a.httpClient.Get(a.baseUrl + "/status")
	if err != nil {
		return alertmanger.Status{}, err
	}
	defer resp.Body.Close() // 关闭响应体

	var status alertmanger.Status
	return status, json.NewDecoder(resp.Body).Decode(&status)

}

func (a *AlertmanagerClient) GetReceivers() ([]alertmanger.Receiver, error) {
	resp, err := a.httpClient.Get(a.baseUrl + "/receivers ")
	if err != nil {
		return nil, err
	}

	var receivers []alertmanger.Receiver
	return receivers, json.NewDecoder(resp.Body).Decode(&receivers)
}

// GetAllSilences  从API获取所有静默列表信息
func (a *AlertmanagerClient) GetAllSilences() ([]alertmanger.Silence, error) {
	resp, err := a.httpClient.Get(a.baseUrl + "/silences")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 关闭响应体

	var silences []alertmanger.Silence
	return silences, json.NewDecoder(resp.Body).Decode(&silences)
}

// NewSilences 创建新的静默信息
func (a *AlertmanagerClient) NewSilences(silence alertmanger.Silence) (alertmanger.SilenceID, error) {
	marshal, _ := json.Marshal(silence)
	resp, err := a.httpClient.Post(a.baseUrl+"/silences", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return alertmanger.SilenceID{}, err
	}
	defer resp.Body.Close() // 关闭响应体

	if resp.StatusCode != 200 {
		var err alertmanger.ResponseError
		json.NewDecoder(resp.Body).Decode(&err)
		return alertmanger.SilenceID{}, err
	}

	var silenceID alertmanger.SilenceID
	return silenceID, json.NewDecoder(resp.Body).Decode(&silenceID)
}

// GetSilenceByID 根据ID获取特定的静默信息
func (a *AlertmanagerClient) GetSilenceByID(id string) (alertmanger.Silence, error) {
	resp, err := http.DefaultClient.Get(a.baseUrl + "/silence/" + id)
	if err != nil {
		return alertmanger.Silence{}, err
	}
	defer resp.Body.Close() // 关闭响应体

	var silence alertmanger.Silence
	return silence, json.NewDecoder(resp.Body).Decode(&silence)
}

// DeleteSilenceByID 根据ID删除特定的静默信息
func (a *AlertmanagerClient) DeleteSilenceByID(id string) error {
	request, _ := http.NewRequest("DELETE", a.baseUrl+"/silence/"+id, nil)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // 关闭响应体

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}

// GetAlerts 从API获取所有告警列表信息
func (a *AlertmanagerClient) GetAlerts() ([]alertmanger.GettableAlert, error) {
	resp, err := http.DefaultClient.Get(a.baseUrl + "/alerts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 关闭响应体

	var alerts []alertmanger.GettableAlert
	return alerts, json.NewDecoder(resp.Body).Decode(&alerts)
}

// NewAlert 创建新的告警信息
func (a *AlertmanagerClient) NewAlert(alert []alertmanger.Alert) error {
	marshal, _ := json.Marshal(alert)
	resp, err := http.DefaultClient.Post(a.baseUrl+"/alerts", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	defer resp.Body.Close() // 关闭响应体

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}

// GetAlertGroup 从API获取所有告警分组信息
func (a *AlertmanagerClient) GetAlertGroup() ([]alertmanger.GettableAlertGroup, error) {
	resp, err := http.DefaultClient.Get(a.baseUrl + "/alerts/groups")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 关闭响应体

	if resp.StatusCode != 200 {
		var err alertmanger.ResponseError
		return nil, json.NewDecoder(resp.Body).Decode(&err)
	}

	var tableAlert []alertmanger.GettableAlertGroup
	return tableAlert, json.NewDecoder(resp.Body).Decode(&tableAlert)
}
