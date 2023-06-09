package pag

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/surflabom/pag/pkg/prometheus"
)

func TestPrometheus(t *testing.T) {
	var config prometheus.Config
	config.BaseURL = "http://127.0.0.1"
	config.TargetConfigPath = ""
	client, _ := NewPrometheusClient(config)
	testJobName := "testjob"
	EndPointAddress := "localhost:9090"

	client.Server(":8090", "/remoteTargets")

	t.Run("Add Prometheus Target", func(t *testing.T) {
		endpoint := prometheus.TargetConfig{
			Targets: []string{},
			Labels: map[string]string{
				"Job": testJobName,
			},
		}

		err := client.AddTarget(endpoint)
		assert.NoError(t, err)
	})

	t.Run("Get Prometheus Configs", func(t *testing.T) {
		targets, err := client.GetAllTarget()
		assert.NoError(t, err)
		log.Println(targets)
	})

	t.Run("Delete Prometheus Target ByName", func(t *testing.T) {
		err := client.DeleteTargetByName(testJobName)
		assert.NoError(t, err)
	})

	t.Run("Get Prometheus Target ByName", func(t *testing.T) {
		target, err := client.GetTargetByName(testJobName)
		assert.NoError(t, err)
		log.Println(target)
	})

	t.Run("Add Target  EndPoint", func(t *testing.T) {
		endpoint := prometheus.EndPoint{
			JobName: testJobName,
			Address: EndPointAddress,
		}
		err := client.AddTargetEndPoint(endpoint)
		assert.NoError(t, err)
	})

	t.Run("Delete Target EndPoint", func(t *testing.T) {
		endpoint := prometheus.EndPoint{
			JobName: testJobName,
			Address: EndPointAddress,
		}
		err := client.DeleteTargetEndPoint(endpoint)
		assert.NoError(t, err)

	})

}
