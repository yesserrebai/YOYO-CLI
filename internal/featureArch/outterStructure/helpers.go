package outterStructure

import (
	Helpers "cli-go/internal/featureArch/helpers"
	"fmt"
	"os"
	"strings"
)

func CreatePackageJsonFile(projectName string) {
	// Define the file path
	filePath := fmt.Sprintf("%s/package.json", projectName)

	// Define the JSON data
	jsonData := map[string]interface{}{
		"name":        projectName,
		"version":     "1.0.0",
		"description": "",
		"main":        "index.js",
		"dependencies": map[string]interface{}{
			"axios":                 "^0.26.1",
			"body-parser":           "^1.19.0",
			"class-transformer":     "^0.3.1",
			"class-validator":       "^0.13.2",
			"config":                "^3.2.2",
			"cors":                  "^2.8.5",
			"express":               "^4.17.1",
			"express-async-errors":  "^3.1.1",
			"google-auth-library":   "^8.7.0",
			"jsonwebtoken":          "^8.5.1",
			"mongodb":               "^4.4.0",
			"mongodb-memory-server": "^8.13.0",
			"mongoose":              "^6.3.0",
			"swagger-ui-express":    "^4.1.6",
			"winston":               "^3.2.1",
		},
		"devDependencies": map[string]interface{}{
			"@types/config":                    "^0.0.34",
			"@types/cors":                      "^2.8.6",
			"@types/express":                   "^4.17.1",
			"@types/glob":                      "^8.0.1",
			"@types/jest":                      "^24.0.18",
			"@types/jsonwebtoken":              "^8.3.5",
			"@types/node":                      "12.7.2",
			"@types/supertest":                 "^2.0.12",
			"@types/swagger-ui-express":        "^4.1.2",
			"@types/validator":                 "^13.7.2",
			"@typescript-eslint/eslint-plugin": "^2.1.0",
			"@typescript-eslint/parser":        "^2.1.0",
			"eslint":                           "^6.3.0",
			"eslint-config-prettier":           "^6.2.0",
			"eslint-plugin-import":             "^2.18.2",
			"eslint-plugin-prettier":           "^3.1.0",
			"jest":                             "^27.5.1",
			"nodemon":                          "^2.0.4",
			"prettier":                         "^1.18.2",
			"supertest":                        "^6.3.3",
			"ts-jest":                          "^27.1.4",
			"ts-node":                          "^8.3.0",
			"typescript":                       "4.4.2",
		},
		"scripts": map[string]interface{}{
			"start":         "nodemon -L",
			"start:prod":    "export NODE_CONFIG_DIR=./src/config && export NODE_ENV=prod && node dist/server.js",
			"start:stg":     "export NODE_CONFIG_DIR=./src/config && export NODE_ENV=staging && node dist/server.js",
			"start:preprod": "export NODE_CONFIG_DIR=./src/config && export NODE_ENV=preprod && node dist/server.js",
			"build":         "tsc",
			"lint":          "eslint --fix src/**/*.ts ",
			"test":          "export NODE_CONFIG_DIR=./src/config && export NODE_ENV=test && jest --detectOpenHandles --coverage --forceExit",
		},
	}
	Helpers.GenerateJsonFile(filePath, jsonData)
}
func CreateTsConfigFile(projectName string) {
	filePath := fmt.Sprintf("%s/tsconfig.json", projectName)
	jsonData := map[string]interface{}{
		"include": []string{"src/**/*"},
		"compilerOptions": map[string]interface{}{
			"module":                       "commonjs",
			"target":                       "es2017",
			"lib":                          []string{"es2017", "DOM"},
			"skipLibCheck":                 true,
			"esModuleInterop":              true,
			"experimentalDecorators":       true,
			"emitDecoratorMetadata":        true,
			"outDir":                       "dist",
			"strict":                       true,
			"sourceMap":                    true,
			"resolveJsonModule":            true,
			"strictPropertyInitialization": false,
		},
	}
	Helpers.GenerateJsonFile(filePath, jsonData)
}
func CreateReadmeFile(projectName string) {
	filePath := fmt.Sprintf("%s/README.md", projectName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating README file: %v\n", err)
		return
	}
	defer file.Close()
}
func CreateNodemonFile(projectName string) {
	filePath := fmt.Sprintf("%s/nodemon.json", projectName)
	jsonData := map[string]interface{}{
		"watch": []string{"src"},
		"ext":   "ts",
		"exec":  "ts-node ./src/server.ts",
		"env": map[string]interface{}{
			"NODE_ENV":        "dev",
			"NODE_CONFIG_DIR": "./src/config",
		},
	}
	Helpers.GenerateJsonFile(filePath, jsonData)
}
func CreateJestConfigFile(projectName string) {
	filepath := fmt.Sprintf("%s/jest.config.json", projectName)
	jsonData := map[string]interface{}{
		"globals": map[string]interface{}{
			"ts-jest": map[string]interface{}{
				"tsconfig": "tsconfig.json",
			},
		},
		"moduleFileExtensions": []string{"ts", "js"},
		"transform": map[string]interface{}{
			"^.+\\.ts$": "ts-jest",
		},
		"testMatch":       []string{"'<rootDir>/src/providers/*/tests/**/*.test.(ts|js)','<rootDir>/src/features/*/tests/**/*.test.(ts|js)'"},
		"testEnvironment": "node",
	}
	Helpers.GenerateJsonFile(filepath, jsonData)
}
func CreateStgDockerFile(projectName string) {
	filePath := fmt.Sprintf("%s/stg.Dockerfile", projectName)
	dockerfileContent := `
FROM node:16-alpine as build

WORKDIR /usr/src/app

ENV NODE_CONFIG_DIR=./src/config

# Install app dependencies
COPY package.json ./

RUN yarn install

# Bundle app source
COPY . .
RUN yarn build -v

EXPOSE 8084
CMD [ "yarn", "start:stg" ]
`
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating docker file: %v\n", err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(dockerfileContent))
	if err != nil {
		fmt.Printf("Error writing docker file content data to file: %v\n", err)
		return
	}
}
func CreatePreprodDockerFile(projectName string) {
	filePath := fmt.Sprintf("%s/preprod.Dockerfile", projectName)
	dockerfileContent := `
FROM node:16-alpine as build

WORKDIR /usr/src/app

ENV NODE_CONFIG_DIR=./src/config

# Install app dependencies
COPY package.json ./

RUN yarn install

# Bundle app source
COPY . .
RUN yarn build -v

EXPOSE 8084
CMD [ "yarn", "start:preprod" ]
`
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating docker file: %v\n", err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(dockerfileContent))
	if err != nil {
		fmt.Printf("Error writing docker file content data to file: %v\n", err)
		return
	}
}
func CreateProdDockerFile(projectName string) {
	filePath := fmt.Sprintf("%s/prod.Dockerfile", projectName)
	dockerfileContent := `
# This file is for staging. DO NOT use it for production env
FROM node:16-alpine as build

WORKDIR /usr/src/app

ENV NODE_CONFIG_DIR=./src/config

# Install app dependencies
COPY package.json ./

RUN yarn install

# Bundle app source
COPY . .
RUN yarn build -v

EXPOSE 8084
CMD [ "yarn", "start:prod" ]
`
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating docker file: %v\n", err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(dockerfileContent))
	if err != nil {
		fmt.Printf("Error writing docker file content data to file: %v\n", err)
		return
	}
}
func CreatePrettierrcFile(projectName string) {
	filepath := fmt.Sprintf("%s/.prettierrc", projectName)
	jsonData := map[string]interface{}{
		"printWidth":    80,
		"tabWidth":      2,
		"useTabs":       false,
		"semi":          true,
		"singleQuote":   true,
		"trailingComma": "all",
	}
	Helpers.GenerateJsonFile(filepath, jsonData)
}
func CreateGitIgnoreFile(projectName string) {
	filePath := fmt.Sprintf("%s/.gitignore", projectName)
	gitIngoreFileContent := `
node_modules
dist
env.prod.json
routes.ts
swagger.json
*.log
TODO
.idea
.vscode
.DS_Store
config/default.json
config/production.json
dep
.env
`
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating gitIgnore file: %v\n", err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(gitIngoreFileContent))
	if err != nil {
		fmt.Printf("Error writing gitIgnore file content data to file: %v\n", err)
		return
	}
}

func CreateEnvFile(projectName string) {
	filePath := fmt.Sprintf("%s/.env", projectName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error writing env file content data to file: %v\n", err)
		return
	}
	defer file.Close()
	envfileContent := `NODE_ENV=dev
	PORT=8080`

	lines := strings.Split(envfileContent, "\n")

	for _, line := range lines {
		_, err = file.Write([]byte(line))
		if err != nil {
			fmt.Printf("Error writing .env file content data to file: %v\n", err)
			return
		}
	}

}

// replace docker file with stg.Dockerfile add preprod.dokckerfile and prod.file
