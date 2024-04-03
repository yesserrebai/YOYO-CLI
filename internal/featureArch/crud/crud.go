package crud

import (
	"bytes"
	"strings"
	"text/template"
)

func GenerateControllerFileContent(featureName string) string {

	// Convert featureName to lowercase for path and method names
	featureNameLower := strings.ToLower(featureName)
	capitalizedFeatureName := strings.Title(featureName)

	// Define data to be passed to the template
	data := struct {
		FeatureName            string
		FeatureNameLower       string
		CapitalizedFeatureName string
	}{
		FeatureName:            featureName,
		FeatureNameLower:       featureNameLower,
		CapitalizedFeatureName: capitalizedFeatureName,
	}

	// Create a new template and parse the template string
	tmpl := template.Must(template.New("controller").Parse(controllerTemplate))

	// Execute the template with the provided data
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
