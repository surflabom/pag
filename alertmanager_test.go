package pag

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/surflabom/pag/pkg/alertmanger"
)

func TestAlertmanager(t *testing.T) {
	var config alertmanger.Config
	config.BaseURL = "http://127.0.0.1:9093"
	client, _ := NewAlertmanagerClient(config)
	silenceUID := "ed1dcace-e5b0-4a04-a97b-2358a2d7af28"
	ctx := context.Background()

	t.Run("New Alert", func(t *testing.T) {
		alerts := []alertmanger.Alert{
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
		err := client.NewAlert(ctx, alerts)
		assert.NoError(t, err)
	})

	t.Run("Get Alerts", func(t *testing.T) {
		alerts, err := client.GetAlerts(ctx)
		assert.NoError(t, err)
		log.Println(alerts)
	})

	t.Run("Get Alert Group", func(t *testing.T) {
		group, err := client.GetAlertGroup(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, group)
		log.Println(group)
	})

	t.Run("New Silences", func(t *testing.T) {
		sli := alertmanger.Silence{
			CreatedBy: "dwadawdawdawd",
			Comment:   "dawdawdawdawdaw",
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
		silences, err := client.NewSilences(ctx, sli)

		assert.NoError(t, err)
		log.Println(silences)
		silenceUID = silences.SilenceID
	})

	t.Run("Get AllSilences", func(t *testing.T) {
		silences, err := client.GetAllSilences(ctx)

		assert.NoError(t, err)
		log.Printf("%+v", silences)
	})

	t.Run("Get Silence", func(t *testing.T) {
		silence, err := client.GetSilenceByID(ctx, silenceUID)
		assert.NoError(t, err)
		log.Println(silence)
	})

	t.Run("Delete Silence", func(t *testing.T) {
		err := client.DeleteSilenceByID(ctx, silenceUID)
		assert.NoError(t, err)
	})

}
