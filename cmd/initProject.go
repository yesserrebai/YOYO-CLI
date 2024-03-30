package cmd

import (
	"cli-go/internal/featureArch/innerStructure"
	"cli-go/internal/featureArch/outterStructure"
	"fmt"
)

func InitProject(projectName string) {
	// to YOYO cli
	// Create outter structure
	outterStructure.GenerateOutterStructure(projectName)
	// create inner structure
	innerStructure.GenerateInnerStructure(projectName)

	fmt.Println("✅ Project successfully created!")
	fmt.Printf("To get started:\n")
	fmt.Printf("\t 👉cd your_project_name and yarn install👈\n")
}
