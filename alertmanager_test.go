package pag

import (
	"PAG/pkg/alertmanger"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestAlertmanager(t *testing.T) {
	client, _ := NewAlertmanagerClient("http://127.0.0.1:9093")
	silenceUID := ""

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
		err := client.NewAlert(alerts)
		assert.NoError(t, err)
	})

	t.Run("Get Alerts", func(t *testing.T) {
		alerts, err := client.GetAlerts()
		assert.NoError(t, err)
		log.Println(alerts)
	})

	t.Run("Get Alert Group", func(t *testing.T) {
		group, err := client.GetAlertGroup()
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
		silences, err := client.NewSilences(sli)

		assert.NoError(t, err)
		log.Println(silences)
		silenceUID = silences.SilenceID
	})

	t.Run("Get AllSilences", func(t *testing.T) {
		silences, err := client.GetAllSilences()

		assert.NoError(t, err)
		log.Printf("%+v", silences)
	})

	t.Run("Get Silence", func(t *testing.T) {
		silence, err := client.GetSilenceByID(silenceUID)
		assert.NoError(t, err)
		assert.NotEmpty(t, silence)
		log.Println(silence)
	})

	t.Run("Delete Silence", func(t *testing.T) {
		err := client.DeleteSilenceByID(silenceUID)
		assert.NoError(t, err)
	})

}
