package alertmanger

import (
	"fmt"
	"time"
)

type SilenceID struct {
	SilenceID string `json:"silenceID"`
}

type Status struct {
	Cluster struct {
		Name   string `json:"name"`
		Status string `json:"status"`
		Peers  []struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"peers"`
	} `json:"cluster"`
	VersionInfo struct {
		Version   string `json:"version"`
		Revision  string `json:"revision"`
		Branch    string `json:"branch"`
		BuildUser string `json:"buildUser"`
		BuildDate string `json:"buildDate"`
		GoVersion string `json:"goVersion"`
	} `json:"versionInfo"`
	Config struct {
		Original string `json:"original"`
	} `json:"config"`
	Uptime time.Time `json:"uptime"`
}

type Receiver struct {
	Name string `json:"name"`
}

type Silence struct {
	ID        string    `json:"id,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	Comment   string    `json:"comment"`
	Status    State     `json:"status,omitempty"`
	StartsAt  time.Time `json:"startsAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Matchers  []Matcher `json:"matchers"`
	EndsAt    time.Time `json:"endsAt"`
}

type State struct {
	State       string   `json:"state"`
	InhibitedBy []string `json:"inhibitedBy"`
	SilencedBy  []string `json:"silencedBy"`
}

type Matcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
	IsEqual bool   `json:"isEqual"`
}

type Alert struct {
	Annotations  map[string]string `json:"annotations"`
	Receivers    []Receiver        `json:"receivers"`
	Fingerprint  string            `json:"fingerprint"`
	StartsAt     time.Time         `json:"startsAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	EndsAt       time.Time         `json:"endsAt"`
	Status       State             `json:"status"`
	Labels       map[string]string `json:"labels"`
	GeneratorURL string            `json:"generatorURL"`
}

type GettableAlert struct {
	Annotations  map[string]string   `json:"annotations"`
	Labels       map[string]string   `json:"labels"`
	Receivers    []map[string]string `json:"receivers"`
	EndsAt       time.Time           `json:"endsAt"`
	Fingerprint  string              `json:"fingerprint"`
	StartsAt     time.Time           `json:"startsAt"`
	Status       State               `json:"status"`
	UpdatedAt    time.Time           `json:"updatedAt"`
	GeneratorURL string              `json:"generatorURL"`
}

type GettableAlertGroup struct {
	Labels   map[string]string `json:"labels"`
	Receiver Receiver          `json:"receiver"`
	Alerts   []Alert           `json:"alerts"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r ResponseError) Error() string {
	return fmt.Sprintf("code:%v,message:%s", r.Code, r.Message)
}
