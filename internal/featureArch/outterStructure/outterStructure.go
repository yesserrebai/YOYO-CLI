package outterStructure

func GenerateOutterStructure(projectName string) {
	CreateEnvFile(projectName)
	CreateGitIgnoreFile(projectName)
	CreateJestConfigFile(projectName)
	CreateNodemonFile(projectName)
	CreatePackageJsonFile(projectName)
	CreatePreprodDockerFile(projectName)
	CreatePrettierrcFile(projectName)
	CreateProdDockerFile(projectName)
	CreateReadmeFile(projectName)
	CreateStgDockerFile(projectName)
	CreateTsConfigFile(projectName)
}
