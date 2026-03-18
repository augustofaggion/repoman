package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Profile struct {
	RepoPath string `json:"profile"`
}

func Greet() string {
	fmt.Println("Hello, welcome to Repoman!")

	return ""
}

func GetProfilePath() string {
	profileFile := "profile.json"
	var profile Profile

	fmt.Println(Greet())

	if _, err := os.Stat(profileFile); os.IsNotExist(err) {
		fmt.Print("Enter your repo directory path: ")
		reader := bufio.NewReader(os.Stdin)
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)

		profile = Profile{RepoPath: path}

		file, _ := os.Create(profileFile)
		defer file.Close()
		json.NewEncoder(file).Encode(profile)
	} else {
		file, _ := os.Open(profileFile)
		defer file.Close()
		json.NewDecoder(file).Decode(&profile)
	}

	return profile.RepoPath
}

func CreateProfile(projects *[]string, projectPaths *[]string) {
	dir := GetProfilePath()
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			*projects = append(*projects, entry.Name())
			*projectPaths = append(*projectPaths, dir+"/"+entry.Name())
		}
	}

	if len(*projects) == 0 {
		fmt.Println("No projects found in the directory.")
		return
	}
}

func ListProjects(projects *[]string, projectPaths *[]string) {
	fmt.Println("Your projects:")

	for i, project := range *projects {
		fmt.Printf("%d: %s\n", i+1, project)
	}

	fmt.Print("Enter the number of the project to open: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // remove new line

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(*projects) {
		fmt.Println("Invalid selection. Please enter a valid number.")
		return
	}

	selectedProject := (*projects)[choice-1]
	selectedPath := (*projectPaths)[choice-1]

	fmt.Printf("Opening %s at %s...\n", selectedProject, selectedPath)
}

func OpenProject(path string) {
	cmd := exec.Command("code", path)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error opening project: %s\n", err)
	} else {
		fmt.Println("Project opened successfully!")
	}
}
