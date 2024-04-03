package cmd

import (
	"cli-go/internal/featureArch/crud"
	Helpers "cli-go/internal/featureArch/helpers"
)

func GenerateCrud(featureName string) {
	generateControllerFileCode(featureName)
}
func generateControllerFileCode(featureName string) {
	//	capitalizedFeatureName := strings.Title(featureName)
	controllerFilePath := "src/features/" + featureName + "/" + featureName + ".controller.ts"
	codeTemplate := crud.GenerateControllerFileContent(featureName)
	Helpers.GenerateJavascriptFile(controllerFilePath, codeTemplate)
}
func generateServiceFileCode(featureName string) {

}
func generateHelperFileCode(featureName string) {

}
