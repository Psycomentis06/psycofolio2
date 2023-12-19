package services_test

import (
	"testing"

	"github.com/psycomentis/psycofolio++/src/services"
)

func TestExportToJson(t *testing.T) {
	cnf := services.CreateDefaultConfig()
	path := "/tmp/config.json"
	err := cnf.ExportToJson(path)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Saved to " + path)
}
