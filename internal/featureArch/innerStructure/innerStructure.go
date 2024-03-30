package innerStructure

func GenerateInnerStructure(projectName string) {
	// Generate folders
	folders := []string{
		"src",
		"src/config",
		"src/config/env",
		"src/exceptions",
		"src/features",
		"src/init",
		"src/middlewares",
		"src/providers",
		"src/shared",
		"src/shared/interfaces",
	}
	CreateFolders(projectName, folders)
	// Begin-- config folder
	CreateEnvFilesJSON(projectName, []string{"dev", "test", "stg", "preprod", "prod"})
	CreateConfigFile(projectName)
	CreateEnumFile(projectName)
	CreateEnvValidationFile(projectName)
	CreateEnvironmentVariablesFile(projectName)
	CreateIndexFile(projectName)
	// End-- config folder

	// Begin-- Exception folder
	CreateHttpExceptionFile(projectName)
	// End-- Exception folder

	// Begin-- Init folder
	CreateInitRoutesFile(projectName)
	// End-- Init folder

	// Begin -- Middlewares folder
	CreateDataValidatorFile(projectName)
	CreateErrorHanlderFile(projectName)
	CreateAuthFile(projectName)
	// End -- Middlewares folder

	// Begin -- Shared folder
	CreateControllerInterfaceFile(projectName)
	CreateLoggerFile(projectName)
	// End -- Shared folder
	CreateAppFile(projectName)
	CreateServerFile(projectName)
}
