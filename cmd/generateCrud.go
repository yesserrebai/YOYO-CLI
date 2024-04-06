package cmd

import (
	"cli-go/internal/featureArch/crud"
	Helpers "cli-go/internal/helpers"
)

func GenerateCrud(featureName string) {
	generateControllerFileCode(featureName)
	generateServiceFileCode(featureName)
	generateHelperFileCode(featureName)
}
func generateControllerFileCode(featureName string) {
	controllerFilePath := "src/features/" + featureName + "/" + featureName + ".controller.ts"
	codeTemplate := crud.GenerateFileContent(featureName, "controller")
	Helpers.GenerateJavascriptFile(controllerFilePath, codeTemplate)
}
func generateServiceFileCode(featureName string) {
	serviceFilePath := "src/features/" + featureName + "/" + featureName + ".service.ts"
	codeTemplate := crud.GenerateFileContent(featureName, "service")
	Helpers.GenerateJavascriptFile(serviceFilePath, codeTemplate)
}
func generateHelperFileCode(featureName string) {
	helperFilePath := "src/features/" + featureName + "/" + featureName + ".helper.ts"
	codeTemplate := crud.GenerateFileContent(featureName, "helper")
	Helpers.GenerateJavascriptFile(helperFilePath, codeTemplate)
}
