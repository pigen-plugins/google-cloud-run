package terraform

import (
	_ "embed"
	tfengine "github.com/pigen-dev/shared/tfengine"
)

//go:embed main.tf
var mainTf []byte
//go:embed variables.tf
var variablesTf []byte
//go:embed output.tf
var outputTf []byte

//Loading tf files
func LoadTFFiles()(tfengine.TerraformFiles) {
	tfFiles := tfengine.TerraformFiles{
		MainTf:      mainTf,
		VariablesTf: variablesTf,
		OutputTf:    outputTf,
	}
	return tfFiles
}