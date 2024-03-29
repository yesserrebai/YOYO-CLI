package cmd

import (
	Helpers "cli-go/internal/featureArch/helpers"
	"fmt"
	"strings"
)

func GenerateFeature(featureName string) {

	folders := []string{
		"src/features/" + featureName,
		"src/features/" + featureName + "/dtos",
		"src/features/" + featureName + "/tests",
		"src/features/" + featureName + "/tests/mocks",
	}
	// içn range of folder call the function
	for _, folder := range folders {
		Helpers.CreateFolder(folder)
	}
	capitalizedFeatureName := strings.Title(featureName)
	controllerFilePath := "src/features/" + featureName + "/" + featureName + ".controller.ts"
	controllerFileContent := fmt.Sprintf(`export default class %sController {}`, capitalizedFeatureName)
	Helpers.GenerateJavascriptFile(controllerFilePath, controllerFileContent)

	serviceFilePath := "src/features/" + featureName + "/" + featureName + ".service.ts"
	serviceFileContent := fmt.Sprintf(`export default class %sService {}`, capitalizedFeatureName)
	Helpers.GenerateJavascriptFile(serviceFilePath, serviceFileContent)

	helpersFilePath := "src/features/" + featureName + "/" + featureName + ".helpers.ts"
	helpersFileContent := fmt.Sprintf(`export default class %sHelper {}`, capitalizedFeatureName)
	Helpers.GenerateJavascriptFile(helpersFilePath, helpersFileContent)

	typesFilePath := "src/features/" + featureName + "/" + featureName + ".types.ts"
	typesFileContent := `// types and interfaces goes here`
	Helpers.GenerateJavascriptFile(typesFilePath, typesFileContent)

	modelFilePath := "src/features/" + featureName + "/" + featureName + ".model.ts"
	modelFileContent := `// mongo schema goes here`
	Helpers.GenerateJavascriptFile(modelFilePath, modelFileContent)

	constantsFilePath := "src/features/" + featureName + "/" + featureName + ".constants.ts"
	constantsFileContent := `// enums and constants goes here `
	Helpers.GenerateJavascriptFile(constantsFilePath, constantsFileContent)

	testFilePath := "src/features/" + featureName + "/tests/" + featureName + ".test.ts"
	testFileContent := "// test implementation goes here "
	Helpers.GenerateJavascriptFile(testFilePath, testFileContent)

	mockFilePath := "src/features/" + featureName + "/tests/mocks/" + featureName + ".mock.ts"
	mockFileContent := "// mocks goes here "
	Helpers.GenerateJavascriptFile(mockFilePath, mockFileContent)

	fmt.Printf("Feature generated: %s\n", featureName)

}
