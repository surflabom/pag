package pag

import (
	"github.com/stretchr/testify/assert"
	"log"
	"pag/pkg/prometheus"
	"testing"
)

func TestPrometheus(t *testing.T) {
	client, _ := NewPrometheusClient("http://127.0.0.1:9090", "")
	testJobName := "testjob"
	EndPointAddress := "localhost:9090"

	client.Server(":8090", "/remoteTargets")

	t.Run("Add Prometheus Target", func(t *testing.T) {
		endpoint := prometheus.Config{
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
