package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/pigen-plugins/google-cloud-run/helpers"
	"github.com/pigen-plugins/google-cloud-run/pkg/terraform"
	shared "github.com/pigen-dev/shared"
	tfengine "github.com/pigen-dev/shared/tfengine"
)


type GoogleCloudRun struct {
	Label string `yaml:"label" json:"label"`
	Config Config `yaml:"config" json:"config"`
	Output Output `yaml:"output" json:"output"`
}



type Config struct {
	ProjectId string `yaml:"project_id" json:"project_id"`
	Location string `yaml:"location" json:"location"`
	ServiceName string `yaml:"service_name" json:"service_name"`
	Image string `yaml:"image" json:"image" omitempty:"true"`
	Ingress string `yaml:"ingress" json:"ingress" omitempty:"true"`
	Unauthenticated bool `yaml:"unauthenticated" json:"unauthenticated" omitempty:"true"`
}

type Output struct {
	CloudRunUrl string `yaml:"cloud_run_url" json:"cloud_run_url"`
}


func (cr *GoogleCloudRun) Initializer(plugin shared.Plugin) (*tfengine.Terraform ,error) {
	config := Config{}
	err:= helpers.YamlConfigParser(plugin.Config, &config)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse YAML config: %v", err)
	}
	if config.Image == "" {
		config.Image = "us-docker.pkg.dev/cloudrun/container/hello"
	}
	if config.Ingress == "" {
		config.Ingress = "INGRESS_TRAFFIC_ALL"
	}
	cr.Config = config
	cr.Label = plugin.Label
	fmt.Println("Parsed config:", cr)
	// Initialize Terraform
	files := terraform.LoadTFFiles()
	tfVars, err := helpers.StructToMap(cr.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert struct to map: %v", err)
	}
	fmt.Println("Terraform variables:", tfVars)
	t, err := tfengine.NewTF(tfVars, files, cr.Label)
	if err != nil {
		return nil, fmt.Errorf("Failed to setup Terraform executor: %v", err)
	}
	
	return t, nil
}



func (cr *GoogleCloudRun) SetupPlugin(plugin shared.Plugin) error {
	tf, err := cr.Initializer(plugin)
	ctx := context.Background()
	if err != nil {
		return fmt.Errorf("Failed to initialize plugin: %v", err)
	}

	// 1. Initialize Terraform
	fmt.Println(cr.Label)
	if err := tf.TerraformInit(ctx, cr.Config.ProjectId, cr.Label); err != nil {
		return fmt.Errorf("Error during Terraform init: %v", err)
	}

	// 2. Plan Terraform changes
	if err := tf.TerraformPlan(ctx); err != nil {
		return fmt.Errorf("Error during Terraform plan: %v", err)
	}

	
	if err := tf.TerraformApply(ctx); err != nil {
		return fmt.Errorf("Error during Terraform apply: %v", err)
	}
	log.Println("Terraform apply completed.")
	return nil
}


func (cr *GoogleCloudRun) GetOutput(plugin shared.Plugin) shared.GetOutputResponse {
	tf, err := cr.Initializer(plugin)
	ctx := context.Background()
	if err != nil {
		return shared.GetOutputResponse{Output: nil, Error: fmt.Errorf("Failed to initialize plugin: %v", err)}
	}
	// 1. Initialize Terraform
	if err := tf.TerraformInit(ctx, cr.Config.ProjectId, cr.Label); err != nil {
		return shared.GetOutputResponse{Output: nil, Error: fmt.Errorf("Error during Terraform init: %v", err)}
	}

	output, err := tf.TerraformOutput(ctx)
	if err != nil {
		return shared.GetOutputResponse{Output: nil, Error: fmt.Errorf("Error during Terraform output: %v", err)}
	}
	log.Println("Terraform output retrieved successfully.")
	return shared.GetOutputResponse{Output: output, Error: nil}
}


func (cr *GoogleCloudRun) Destroy(plugin shared.Plugin) error {
	tf, err := cr.Initializer(plugin)
	if err != nil {
		return fmt.Errorf("Failed to initialize plugin: %v", err)
	}
	ctx := context.Background()
	// 1. Initialize Terraform
	if err := tf.TerraformInit(ctx, cr.Config.ProjectId, cr.Label); err != nil {
		return fmt.Errorf("Error during Terraform init: %v", err)
	}

	if err := tf.TerraformDestroy(ctx); err != nil {
		return fmt.Errorf("Error during Terraform destroy: %v", err)
	}
	log.Println("Terraform destroy completed.")
	return nil
}