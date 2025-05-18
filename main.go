package main

import (
	"fmt"

	shared "github.com/pigen-dev/shared"
	"github.com/pigen-plugins/google-cloud-run/pkg"
)

func main() {
	plugin := shared.Plugin{
		Label: "google-cloud-run",
		Config: map[string]any{
			"project_id":   "aidodev",
			"location":     "europe-west1",
			"service_name": "your-service-name",
			"unauthenticated": true,
		},
	}
	cr := pkg.GoogleCloudRun{}
	err := cr.SetupPlugin(plugin)
	if err != nil {
		fmt.Println("Error setting up plugin:", err)
	}
	output := cr.GetOutput(plugin)
	fmt.Println("Cloud Run URL:", output)
}