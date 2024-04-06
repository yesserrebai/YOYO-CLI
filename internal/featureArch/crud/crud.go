package crud

import (
	"bytes"
	"strings"
	"text/template"
)

func GenerateFileContent(featureName string, file string) string {

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

	tmpl := template.Must(template.New("controller").Parse(controllerTemplate))
	if file == "service" {
		tmpl = template.Must(template.New("service").Parse(serviceTemplate))
	}
	if file == "helper" {
		tmpl = template.Must(template.New("helper").Parse(helperTemplate))
	}
	// Execute the template with the provided data
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
