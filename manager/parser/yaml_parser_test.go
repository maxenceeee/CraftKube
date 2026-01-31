package parser_test

import (
	"os"
	"testing"

	"github.com/goccy/go-yaml"
	"xamence.eu/craftkube/internal"
)

func TestCreateFileFromService(t *testing.T) {
	service := &internal.Service{
		Name:  "test-service",
		Image: "test-image",
		AutoScaling: internal.AutoScalingConfig{
			Enable:       true,
			MinInstances: 1,
			MaxInstances: 5,
			Cooldown:     60,
			TriggerLogic: "AND",
			Upscale: internal.UpscalePolicy{
				Policies: []internal.Policy{
					{
						Name: "cpu-policy",
						Type: "CPU",
						Inhibitor: internal.Inhibitor{
							Threshold: 80,
							Operator:  "GreaterThan",
							ValueType: "Percentage",
						},
						ScaleAmount: 1,
					},
				},
			},
			Downscale: internal.DownscalePolicy{
				Policies: []internal.Policy{
					{
						Name: "cpu-policy-down",
						Type: "CPU",
						Inhibitor: internal.Inhibitor{
							Threshold: 30,
							Operator:  "LessThan",
							ValueType: "Percentage",
						},
						ScaleAmount: 1,
					},
				},
			},
		},
	}

	bytes, err := yaml.Marshal(&service)
	if err != nil {
		t.Fatalf("Failed to marshal service to YAML: %v", err)
	}
	file, err := os.Create("yaml_test.yml")
	if err != nil {
		t.Fatalf("Failed to write YAML to file: %v", err)
	}

	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		t.Fatalf("Failed to write YAML to file: %v", err)
	}
}
