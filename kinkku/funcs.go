package kinkku

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func StartUp() {

	fmt.Println(FgMagenta + banner + Reset)
	fmt.Println(FgCyan + Italic + slogan + Reset)
}

func GetArgs() {

	if len(os.Args) != 3 && len(os.Args) != 1 {
		fmt.Println(FgRed + "Oops, skill issue: You got the ham, but where is the mustard?" + Reset)
		fmt.Println("Wrong number of arguments.")
		fmt.Println("kinkku usage example:")
		fmt.Println("$ kinkku ./directory 6969")
		os.Exit(0)

	}

	if len(os.Args) == 1 {
		path = "."
		port = "8080"
		fmt.Println()
		fmt.Println(FgCyan + "Using current directory and port 8080 by default." + Reset)
	} else {

		path = os.Args[1]
		port = os.Args[2]
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(FgRed + "Oops, skill issue: Can't find your fridge." + Reset)
		fmt.Println("Provided directory path does not exist.")
		fmt.Println("kinkku usage example:")
		fmt.Println("$ kinkku ./directory 6969")

		os.Exit(0)

		// Handle the case where the directory does not exist
	}
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
			fmt.Println(FgRed + "Oops, skill issue: there is no ham to serve." + Reset)
			fmt.Println("Error starting server:", err)

			os.Exit(0)
		}
		if restartCount == 0 {
			fmt.Println(FgGreen + "Server is up! Let's Go!" + Reset)
			restartCount++
		} else if restartCount == 69 {
			fmt.Println(FgGreen + "(" + strconv.Itoa(restartCount) + ") " + "Server restarted." + Reset)
			fmt.Println(FgCyan + string(noice) + Reset)
			restartCount++
		} else {
			fmt.Println(FgGreen + "(" + strconv.Itoa(restartCount) + ") " + "Server restarted." + Reset)
			restartCount++
		}

		// Wait for the server to start
		// After restarting refresh the page
	}()
}

// Function to find servers pid number on mac.
func getServerPID(port string) (string, error) {
	cmd := exec.Command("lsof", "-ti", "tcp:"+port)
	output, err := cmd.Output()
	if err != nil {
		return "0", err
	}
	pid := strings.TrimSpace(string(output))
	if len(pid) > 5 { //this part is for error correction. Not sure if it gets rid of all problems tho.
		pid = pid[len(pid)-6:]
	}
	return pid, nil
}

// Function to kill a process listening on the specified port
func killServerOnPort(port string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "darwin" {
		pid, err := getServerPID(port)
		if err != nil {
			return err
		}
		if pid == "0" {
			return errors.New("Server PID not found, unable to kill server. Please restart kinkku.")
		}
		// Kill process using pid-number on mac
		cmd = exec.Command("kill", "-9", pid)
	} else {
		// Kill process on other systems.
		cmd = exec.Command("fuser", "-k", "-n", "tcp", port)
	}
	if err := cmd.Run(); err != nil {
		return nil
	}
	return nil
}

// Function to watch for file changes recursively in a directory
func WatchFiles(changes chan<- string) {
	fileModTimes := getFileModTimes()
	goFilesFound := false
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			goFilesFound = true // Set to true if Go file is found
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !goFilesFound {
		fmt.Println(FgRed + "Oops, skill issue: there is no ham to serve." + Reset)
		fmt.Println("No Go files found in provided directory.")
		os.Exit(0)
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

func getFileModTimes() map[string]time.Time {
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
