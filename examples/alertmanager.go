package main

import (
	"context"
	"log"
	"time"

	"github.com/surflabom/pag"
	"github.com/surflabom/pag/pkg/alertmanger"
)

func main() {

	//添加配置 初始化客户端
	config := alertmanger.Config{BaseURL: "http://127.0.0.1:9093"}
	client, _ := pag.NewAlertmanagerClient(config)

	ctx := context.Background()

	//新建一个 Alert
	alert := []alertmanger.Alert{
		{
			Annotations: map[string]string{
				"field1": "123",
				"field2": "123",
				"field3": "123",
			},
			Receivers: []alertmanger.Receiver{
				{Name: "123"},
			},
			Fingerprint: "Fingerprint",
			StartsAt:    time.Now().UTC(),
			EndsAt:      time.Now().UTC().Add(1 * time.Hour),
			Status: alertmanger.State{
				State: "active",
			},
			Labels: map[string]string{
				"field1": "value1",
				"field":  "value2",
			},
			GeneratorURL: "https://dawdaw.com",
		},
	}
	if err := client.NewAlert(ctx, alert); err != nil {
		log.Println(err)
	}

	//获取所有的 Alert
	alerts, err := client.GetAlerts(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(alerts)

	//新建一个 Silence
	sli := alertmanger.Silence{
		CreatedBy: "wwwww",
		Comment:   "new silence",
		StartsAt:  time.Now().UTC(),
		EndsAt:    time.Now().UTC().Add(1 * time.Hour),
		Matchers: []alertmanger.Matcher{
			{
				Name:    "env",
				Value:   "production",
				IsRegex: false,
				IsEqual: true,
			},
		},
		Status: alertmanger.State{
			State:       "active",
			InhibitedBy: []string{},
			SilencedBy:  []string{},
		},
	}
	silencesID, err := client.NewSilences(ctx, sli)
	log.Println(silencesID)

	//获取所有 Silence
	silences, err := client.GetAllSilences(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(silences)

}
