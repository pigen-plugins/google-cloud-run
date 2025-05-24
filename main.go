package main

import (
	//"fmt"

	"github.com/hashicorp/go-plugin"
	shared "github.com/pigen-dev/shared"
	"github.com/pigen-plugins/google-cloud-run/pkg"
)

func main() {
	// plugin := shared.Plugin{
	// 	Label: "GOOGLE_CLOUD_RUN_DEMO",
	// 	Config: map[string]any{
	// 		"project_id":   "aidodev",
	// 		"location":     "europe-west1",
	// 		"service_name": "pigen-demo",
	// 		"unauthenticated": true,
	// 	},
	// }
	// cr := pkg.GoogleCloudRun{}
	// err := cr.SetupPlugin(plugin)
	// if err != nil {
	// 	fmt.Println("Error setting up plugin:", err)
	// }
	// output := cr.GetOutput(plugin)
	// fmt.Println("Cloud Run Output:", output)
	cr := &pkg.GoogleCloudRun{}
	pluginMap := map[string]plugin.Plugin{"pigenPlugin": &shared.PigenPlugin{Impl: cr}}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         pluginMap,
	})
}