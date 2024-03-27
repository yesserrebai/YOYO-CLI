package main

import (
	"cli-go/internal/featureArch/feature"
	innerStructure "cli-go/internal/featureArch/innerStructure"
	outterStructure "cli-go/internal/featureArch/outterStructure"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectSetup = &cobra.Command{
	Use:   "init",
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
		// Create outter structure
		outterStructure.GenerateOutterStructure(projectName)
		// create inner structure
		innerStructure.GenerateInnerStructure(projectName)

		fmt.Println("Now cd your project and type yarn install")

	},
}
var generateFeature = &cobra.Command{
	Use:   "g-feature",
	Short: "generate feature",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a prompt
		promptFeatureName := promptui.Prompt{
			Label: "Enter Feature name",
		}
		// Get Feature name
		featureName, err := promptFeatureName.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			os.Exit(1)
		}
		feature.GenerateFeature(featureName)
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "yoyo"}

	// Add the  command to the root command
	rootCmd.AddCommand(projectSetup, generateFeature)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
