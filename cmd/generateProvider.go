package cmd

import (
	Helpers "cli-go/internal/featureArch/helpers"
	"fmt"
	"strings"
)

func GenerateProvider(providerName string) {

	folders := []string{
		"src/providers/" + providerName,
		"src/providers/" + providerName + "/tests",
		"src/providers/" + providerName + "/tests/mocks",
	}
	// içn range of folder call the function
	for _, folder := range folders {
		Helpers.CreateFolder(folder)
	}
	capitalizedProviderName := strings.Title(providerName)
	providerFilePath := "src/providers/" + providerName + "/" + providerName + ".provider.ts"
	providerFileContent := fmt.Sprintf(`export default class %sProvider {}`, capitalizedProviderName)
	Helpers.GenerateJavascriptFile(providerFilePath, providerFileContent)

	typesFilePath := "src/providers/" + providerName + "/" + providerName + ".types.ts"
	typesFileContent := `// types and interfaces goes here`
	Helpers.GenerateJavascriptFile(typesFilePath, typesFileContent)

	constantsFilePath := "src/providers/" + providerName + "/" + providerName + ".constants.ts"
	constantsFileContent := `// enums and constants goes here `
	Helpers.GenerateJavascriptFile(constantsFilePath, constantsFileContent)

	testFilePath := "src/providers/" + providerName + "/tests/" + providerName + ".test.ts"
	testFileContent := "// test implementation goes here "
	Helpers.GenerateJavascriptFile(testFilePath, testFileContent)

	mockFilePath := "src/providers/" + providerName + "/tests/mocks/" + providerName + ".mock.ts"
	mockFileContent := "// mocks goes here "
	Helpers.GenerateJavascriptFile(mockFilePath, mockFileContent)

	fmt.Printf("Feature generated: %s\n", providerName)

}
