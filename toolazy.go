package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	port                 = "8080"    // Specify the port your server is running on
	path                 = "../cars" // Specify the path to your project directory
	modificationDetected = false     // Flag to track if any file modification has been detected
	restartCount         = 0
)

func main() {
	//run the server
	restartServer()
	// Create a channel to receive file change events
	fileChanges := make(chan string)

	// Start watching for file changes
	go watchFiles(fileChanges)

	// Watch for file change events and restart the server
	for {
		select {
		case <-fileChanges:
			// If a modification has been detected, restart the server
			if modificationDetected {
				restartServer()
				modificationDetected = false // Reset the flag after restarting
			}
		}
	}
}

// Function to restart the Go server
func restartServer() {

	// Kill any process listening on the specified port
	if err := killServerOnPort(port); err != nil {
		fmt.Println("Error killing server:", err)
	}
	time.Sleep(20 * time.Millisecond)
	// Start the server again in a separate goroutine
	go func() {
		goRun := exec.Command("go", "run", ".")
		goRun.Dir = path // Set the working directory for the command
		goRun.Stdout = os.Stdout
		goRun.Stderr = os.Stderr
		if err := goRun.Start(); err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
		if restartCount == 0 {
			fmt.Println("Server is running")
			restartCount++
		} else {
			fmt.Println("(" + strconv.Itoa(restartCount) + ") " + "server restarted")
		}

		// Wait for the server to start
		// After restarting refresh the page
	}()
}

// Function to kill a process listening on the specified port
func killServerOnPort(port string) error {
	cmd := exec.Command("fuser", "-k", "-n", "tcp", port)
	if err := cmd.Run(); err != nil {
		return nil
	}
	return nil
}

// Function to watch for file changes recursively in a directory
func watchFiles(changes chan<- string) {
	fileModTimes := getFileModTimes(path)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Continuously monitor for file changes
	for {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			checkFileModifications(path, info, fileModTimes, changes)
			return nil
		})
		if err != nil {
			fmt.Println("Error:", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func checkFileModifications(path string, info os.FileInfo, fileModTimes map[string]time.Time, changes chan<- string) {
	//fmt.Println("Checking modification for:", path)
	modTime := info.ModTime()
	//    fmt.Println("Current modification time:", modTime)
	lastModTime := fileModTimes[path]
	//    fmt.Println("Last modification time recorded:", lastModTime)
	if strings.HasSuffix(path, ".go") {
		if !lastModTime.IsZero() {
			if modTime.After(lastModTime) {
				fmt.Println("Go file modified:", path)
				changes <- path
				fileModTimes[path] = modTime
				modificationDetected = true // Set the flag indicating modification detected
				//            fmt.Println("Modification time updated for:", path, "to", modTime)
			}
		} else {
			fmt.Println("Initial detection of file:", path)
			changes <- path
			fileModTimes[path] = modTime
		}
	}
}

func getFileModTimes(path string) map[string]time.Time {
	fileModTimes := make(map[string]time.Time)
	return fileModTimes
}
