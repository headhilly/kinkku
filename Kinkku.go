package main

import kinkku "github.com/headhilly/kinkku/kinkku"

func main() {

	kinkku.GetArgs()
	kinkku.StartUp()
	//run the server
	kinkku.RestartServer()
	// Create a channel to receive file change events
	fileChanges := make(chan string)

	// Start watching for file changes
	go kinkku.WatchFiles(fileChanges)

	// Watch for file change events and restart the server
	for {
		select {
		case <-fileChanges:
			// If a modification has been detected, restart the server
			if kinkku.ModificationDetected {
				kinkku.RestartServer()
				kinkku.ModificationDetected = false // Reset the flag after restarting
			}
		}
	}
}
