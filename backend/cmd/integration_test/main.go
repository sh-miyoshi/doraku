package main

import (
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
)

func main(){
	// Move to top dir
	os.Chdir("../../")
	projectDir, _ := filepath.Abs(".")
	fmt.Printf("Project Dir: %s\n", projectDir)

	// Run server as local
	if err := exec.Command("go","build").Run(); err != nil {
		fmt.Printf("Failed to build server: %v\n", err)
		return
	}
	fmt.Println("Build was finished successfully")

	cmd := exec.Command(
		"./doraku",
		"-alpha",
		"-logfile=server.log",
	)

	if err:= cmd.Start(); err!= nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer cmd.Process.Kill()

	// Test all API

	exec.Command("sleep", "5").Run()// TODO for debug
}