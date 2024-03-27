package feature

import (
	Helpers "cli-go/internal/featureArch/helpers"
	"fmt"
)

func GenerateFeature(featureName string) {
	// generate folder based on feature name
	// generate files absed on

	folders := []string{
		"src/" + featureName,
		"src/" + featureName, "/tests",
	}
	// içn range of folder call the function
	for _, folder := range folders {
		Helpers.CreateFolder(fmt.Sprintf("%s/%s", "aaaaa", folder))
	}
}
