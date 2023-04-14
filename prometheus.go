package pag

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"pag/pkg/prometheus"
	"strconv"
	"time"
)

type PrometheusClient struct {
	httpClient *http.Client
	baseURL    string
	configPath string
}

func (p *PrometheusClient) handler(w http.ResponseWriter, r *http.Request) {
	configs, err := p.GetAllTarget()
	if err != nil {
		log.Println(err)
		return
	}

	marshal, _ := json.Marshal(configs)
	w.Header().Add("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(marshal))
}

func (p *PrometheusClient) Server(address, pattern string) {
	go func() {
		http.HandleFunc(pattern, p.handler)
		http.ListenAndServe(address, nil)
	}()
}

func NewPrometheusClient(baseURL, TargetConfigPath string) (*PrometheusClient, error) {
	if baseURL == "" {
		return nil, errors.New("baseURL 不能为空")
	}

	if TargetConfigPath == "" {
		TargetConfigPath = "configs.json"
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("baseURL 不可用, resp.StatusCode=" + strconv.Itoa(resp.StatusCode))
	}

	if _, err := os.Stat(TargetConfigPath); err != nil {
		if os.IsNotExist(err) {
			if err := prometheus.CreateConfig(TargetConfigPath); err != nil {
				panic(err)
			}
		}
	}

	return &PrometheusClient{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		baseURL:    baseURL,
		configPath: TargetConfigPath,
	}, err

}

func (p *PrometheusClient) GetTargetByName(targetName string) (prometheus.Config, error) {
	return prometheus.ReadConfig(p.configPath, targetName)
}

func (p *PrometheusClient) GetAllTarget() ([]prometheus.Config, error) {
	return prometheus.ReadConfigs(p.configPath)
}

func (p *PrometheusClient) AddTarget(config prometheus.Config) error {
	return prometheus.AddConfig(p.configPath, config)
}

func (p *PrometheusClient) DeleteTargetByName(configName string) error {
	return prometheus.DeleteConfig(p.configPath, configName)
}

func (p *PrometheusClient) AddTargetEndPoint(target prometheus.EndPoint) error {
	return prometheus.AddTarget(p.configPath, target)
}

func (p *PrometheusClient) DeleteTargetEndPoint(target prometheus.EndPoint) error {
	return prometheus.DeleteTarget(p.configPath, target)
}
