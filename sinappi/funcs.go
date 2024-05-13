package kinkku

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func StartUp() {
	banner, err := os.ReadFile("./sinappi/banner.txt")
	if err != nil {
		fmt.Println("Error printing my sweet banner:", err)
	}
	fmt.Println(FgMagenta + string(banner) + Reset)
	fmt.Println(FgCyan + Italic + "------------------------Ain't nobody got time for that !-------------------------" + Reset)
}

func GetArgs() {

	if len(os.Args) != 3 {
		fmt.Println(FgRed + "It's literally just 2 argument bro how can you fuck this up? Here's a usage example" + Reset)
		fmt.Println("$ go run . ./directory 6969")
		os.Exit(0)

	}
	path = os.Args[1]
	port = os.Args[2]
}

// Function to restart the Go server
func RestartServer() {

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
			fmt.Println(FgGreen + "Server is up!" + Reset)
			restartCount++
		} else if restartCount == 69 {
			fmt.Println(FgGreen + "(" + strconv.Itoa(restartCount) + ") " + "Server restarted." + Reset)
			noice, err := os.ReadFile("./sinappi/noice.txt")
			if err != nil {
				fmt.Println("Error printing my sweet banner:", err)
				fmt.Println(FgCyan + string(noice) + Reset)
				restartCount++
			}
		} else {
			fmt.Println(FgGreen + "(" + strconv.Itoa(restartCount) + ") " + "Server restarted." + Reset)
			restartCount++
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
func WatchFiles(changes chan<- string) {
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
				fmt.Println(FgMagenta + "Go file modification detected:" + Reset + path)
				changes <- path
				fileModTimes[path] = modTime
				ModificationDetected = true // Set the flag indicating modification detected
				//            fmt.Println("Modification time updated for:", path, "to", modTime)
			}
		} else {
			fmt.Println(FgYellow + "Go file found:" + Reset + path)
			changes <- path
			fileModTimes[path] = modTime
		}
	}
}

func getFileModTimes(path string) map[string]time.Time {
	fileModTimes := make(map[string]time.Time)
	return fileModTimes
}

/*
       /\   /\
         \_/
    __   / \   __
  -'  `. \_/ .'  `-
        \/ \/
   _.---(   )---._
_.'   _.-\_/-._   `._
     /   /_\   \
    /   /___\   \
   /   |_____|   \
_.'    | ___ |    `._                             not all bugs are bad!
        \___/                                               stop bug discrimination!




           \(")/
           -( )-
           /(_)\


*/
