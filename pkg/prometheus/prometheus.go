package prometheus

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

func CreateConfig(path string) error {
	targets, err := os.Create(path)
	if err != nil {
		return err
	}
	defer targets.Close()

	_, err = targets.WriteString("[]")
	return err
}

func writeConfig(path string, configs []TargetConfig) error {
	marshal, _ := json.Marshal(configs)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(marshal)
	return err
}

func ReadConfigs(path string) ([]TargetConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = CreateConfig(path)
			return nil, err
		}
	}

	var configs []TargetConfig
	return configs, json.Unmarshal(bytes, &configs)
}

func ReadConfig(path string, configName string) (TargetConfig, error) {
	jobs, err := ReadConfigs(path)
	if err != nil {
		return TargetConfig{}, err
	}

	for _, t := range jobs {
		if t.Labels["job"] == configName {
			return t, err
		}
	}

	return TargetConfig{}, errors.New("未找到")
}

func AddConfig(path string, config TargetConfig) error {
	configs, err := ReadConfigs(path)
	if err != nil {
		return err
	}

	configs = append(configs, config)
	return writeConfig(path, configs)
}

func DeleteConfig(path, configName string) error {
	configs, err := ReadConfigs(path)
	if err != nil {
		return err
	}

	for i, t := range configs {
		if t.Labels["job"] == configName {
			configs = append(configs[:i], configs[i+1:]...)
			break
		}
	}

	return writeConfig(path, configs)
}

func AddTarget(path string, target EndPoint) error {
	configs, err := ReadConfigs(path)
	if err != nil {
		return err
	}

	for i, t := range configs {
		if t.Labels["job"] == target.JobName {
			configs[i].Targets = append(configs[i].Targets, target.Address)
			break
		}
	}

	return writeConfig(path, configs)
}

func DeleteTarget(path string, target EndPoint) error {
	configs, err := ReadConfigs(path)
	if err != nil {
		return err
	}

	for i, t := range configs {
		if t.Labels["job"] == target.JobName {
			for i2, s := range configs[i].Targets {
				if s == target.Address {
					configs[i].Targets = append(configs[i].Targets[:i2], configs[i].Targets[i2+1:]...)
					break
				}
			}
			break
		}
	}

	return writeConfig(path, configs)
}

type Config struct {
	BaseURL          string //prometheus 服务地址 例如: "http://127.0.0.1:9090"
	TargetConfigPath string //target 配置文件路径 为空则默认在程序根目录下生成`configs.json`作为配置文件
}

type TargetConfig struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

type TargetMetaData struct {
	Status string     `json:"status"`
	Data   []MetaData `json:"data"`
}

type Target struct {
	Instance string `json:"instance"`
	Job      string `json:"job"`
}

type MetaData struct {
	Target Target `json:"target"`
	Metric string `json:"metric"`
	Type   string `json:"type"`
	Help   string `json:"help"`
	Unit   string `json:"unit"`
}

type DiscoveredLabels struct {
	Address        string `json:"__address__"`
	MetaURL        string `json:"__meta_url"`
	MetricsPath    string `json:"__metrics_path__"`
	Scheme         string `json:"__scheme__"`
	ScrapeInterval string `json:"__scrape_interval__"`
	ScrapeTimeout  string `json:"__scrape_timeout__"`
	Job            string `json:"job"`
}

type Labels struct {
	Instance string `json:"instance"`
	Job      string `json:"job"`
}

type ActiveTargets struct {
	DiscoveredLabels   DiscoveredLabels `json:"discoveredLabels,omitempty"`
	Labels             Labels           `json:"labels"`
	ScrapePool         string           `json:"scrapePool"`
	ScrapeURL          string           `json:"scrapeUrl"`
	GlobalURL          string           `json:"globalUrl"`
	LastError          string           `json:"lastError"`
	LastScrape         time.Time        `json:"lastScrape"`
	LastScrapeDuration float64          `json:"lastScrapeDuration"`
	Health             string           `json:"health"`
	ScrapeInterval     string           `json:"scrapeInterval"`
	ScrapeTimeout      string           `json:"scrapeTimeout"`
}

type Data struct {
	ActiveTargets  []ActiveTargets `json:"activeTargets"`
	DroppedTargets []interface{}   `json:"droppedTargets"`
}

type EndPoint struct {
	JobName string
	Address string
}
