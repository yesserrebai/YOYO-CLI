package main

import (
	"cli-go/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func printBox(message string, width int, height int) {
	if width < len(message)+4 {
		panic("Width is too small to fit the message")
	}

	topBorder := strings.Repeat("*", width)
	sideBorder := "*" + strings.Repeat(" ", width-2) + "*"

	fmt.Println(topBorder)
	for i := 0; i < height-2; i++ {
		fmt.Println(sideBorder)
	}

	// Center the message horizontally within the box
	messageIndent := (width - len(message)) / 2
	messageLine := sideBorder[:messageIndent] + message + sideBorder[messageIndent+len(message):]

	fmt.Println(messageLine)

	fmt.Println(topBorder)
}

var projectSetup = &cobra.Command{
	Use:   "init",
	Short: "generate project",
	Run: func(command *cobra.Command, args []string) {

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
		printBox("WELCOME TO YOYO-CLI", 60, 5)
		cmd.InitProject(projectName)

	},
}
var generateFeature = &cobra.Command{
	Use:   "g-feature",
	Short: "generate feature",
	Run: func(command *cobra.Command, args []string) {
		// Create a prompt
		promptFeatureName := promptui.Prompt{
			Label: "Enter Feature name (🟡 name should be singular not plural)",
		}
		// Get Feature name
		featureName, err := promptFeatureName.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			os.Exit(1)
		}
		cmd.GenerateFeature(featureName)
	},
}
var generateProvider = &cobra.Command{
	Use:   "g-provider",
	Short: "generate provider",
	Run: func(command *cobra.Command, args []string) {
		// Create a prompt
		promptProviderName := promptui.Prompt{
			Label: "Enter Provider name(🟡 name should be singular not plural)",
		}
		// Get Provider name
		providerName, err := promptProviderName.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			os.Exit(1)
		}
		cmd.GenerateProvider(providerName)
	},
}
var generateCrud = &cobra.Command{
	Use:   "g-crud",
	Short: "generate crud",
	Run: func(command *cobra.Command, args []string) {
		// Create a prompt
		promptFeatureName := promptui.Prompt{
			Label: "what feature name you want to create crud for ?",
		}

		featureName, err := promptFeatureName.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			os.Exit(1)
		}
		cmd.GenerateCrud(featureName)
	},
}

func main() {
	var rootcommand = &cobra.Command{Use: "yoyo"}

	// Add the  command to the root command
	rootcommand.AddCommand(projectSetup, generateFeature, generateProvider, generateCrud)

	// Execute the root command
	if err := rootcommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
