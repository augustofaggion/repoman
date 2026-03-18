package main

import (
	"repoman/helpers"
)

func main() {

	var projects []string
	var projectPaths []string

	// Create profile or reads with profile is already created
	helpers.CreateProfile(&projects, &projectPaths)
	
	// Make a list
	helpers.ListProjects(&projects, &projectPaths)

    // Open selected project
	helpers.OpenProject(projectPaths[0])
}