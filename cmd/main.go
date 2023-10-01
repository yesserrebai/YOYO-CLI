package main

import (
	outterstrcture "cli-go/internal/featureArch/outterStructure"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectSetup = &cobra.Command{
	Use:   "project-setup",
	Short: "generate project",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a prompt
		promptProjectName := promptui.Prompt{
			Label: "Enter project name",
		}
		// Get project name
		projectName, err := promptProjectName.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			os.Exit(1)
		}
		dirPath := projectName
		errMkdir := os.Mkdir(dirPath, 0700)
		if errMkdir != nil {
			fmt.Println("Error creating directory", errMkdir)
			return
		}
		fmt.Println("generating your project", projectName)
		outterstrcture.CreatePackageJsonFile(projectName)
		outterstrcture.CreateDockerFile(projectName)
		outterstrcture.CreateEnvFile(projectName)
		outterstrcture.CreateGitIgnoreFile(projectName)
		outterstrcture.CreateJestConfigFile(projectName)
		outterstrcture.CreateNodemonFile(projectName)
		outterstrcture.CreatePrettierrcFile(projectName)
		outterstrcture.CreateReadmeFile(projectName)
		outterstrcture.CreateTsConfigFile(projectName)
		fmt.Println("Now cd your project and type yarn install")

	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "yoyo"}

	// Add the  command to the root command
	rootCmd.AddCommand(projectSetup)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
