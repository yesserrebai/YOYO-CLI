package main

import (
	"cli-go/cmd"
	innerStructure "cli-go/internal/featureArch/innerStructure"
	outterStructure "cli-go/internal/featureArch/outterStructure"
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
		// to YOYO cli
		// Create outter structure
		outterStructure.GenerateOutterStructure(projectName)
		// create inner structure
		innerStructure.GenerateInnerStructure(projectName)

		fmt.Println("✅ Project successfully created!")
		fmt.Printf("To get started:\n")
		fmt.Printf("\t 👉cd your_project_name and yarn install👈\n")

	},
}
var generateFeature = &cobra.Command{
	Use:   "g-feature",
	Short: "generate feature",
	Run: func(command *cobra.Command, args []string) {
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
		cmd.GenerateFeature(featureName)
	},
}
var generateProvider = &cobra.Command{
	Use:   "g-provider",
	Short: "generate provider",
	Run: func(command *cobra.Command, args []string) {
		// Create a prompt
		promptProviderName := promptui.Prompt{
			Label: "Enter Provider name",
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
var generateJwtMiddleware = &cobra.Command{
	Use:   "g-jwt",
	Short: "generate jwt middleware",
	Run: func(command *cobra.Command, args []string) {
		cmd.GenerateJwtMiddleware()
	},
}

func main() {
	var rootcommand = &cobra.Command{Use: "yoyo"}

	// Add the  command to the root command
	rootcommand.AddCommand(projectSetup, generateFeature, generateProvider, generateJwtMiddleware)

	// Execute the root command
	if err := rootcommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
