
# Kinkku

## What is it?

Kinkku is a simple command line tool that simulates a live preview of your golang server by restarting it whenever a go file is modified.

I was annoyed by the cycle of manually shutting down and then running the server everytime I wanted to see changes, so I made kinkku to make working on my school projects more convenient.

ps. Don't judge my code. this was intended to be a tool just for myself so I didn't care much for good practices :)


## Installation

1. **Install via Go**:

Use the following command to install `kinkku` directly:

```bash
  go install github.com/headhilly/kinkku@latest
````
2. **Add to PATH**:

```bash
  export PATH=$PATH:$(go env GOPATH)/bin
```

After adding the line, run the following to apply the changes:

```bash
  source ~/.bashrc  # or source ~/.zshrc for zsh users
  # Optionally, you could close and restart the terminal
```

Now, you should be able to run kinkku from anywhere on your system.


## Usage

Now Kinkku is installed and ready to go ham! you're all set to use the command `kinkku`as so:

```bash
kinkku < server path > < port number >
```
For kinkku to be able to do it's job as intended, you need to provide the path to the server you are working on, followed by the port number your server is running on.

Here is an example, if you are using kinkku from the root folder of your server, you would use `.` as the path, followed by the port, as so:

```bash
kinkku . 6969
```
or if you are somewhere else, you can still navigate to the server as so:

```bash
kinkku ../example/notporn/server 6969
```
(Update: now if no extra arguments are added after kinkku it would use the current directory and port 8080 by default. If that's what you're using, then only use `kinkku` and you're good to go.)

that's it! kinkku now got your back and you won't need to manually restart the server anymore everytime you modify your go files!

## Here is everything kinkku does:
1. ### Starts your server
Kinkku will navigate to the path you have provided and execute "go run ." for you, so you do not need to run your server seperately when using kinkku.
If the server doesn't run at that point, im afraid it's a you problem.

2. ### Finds go files
After running the server, kinkku will scan the directory you have provided including all it's subdirectories for go files.
if no go files are found, Kinkku will exit.

3. ### Monitors the files
The moment kinkku finds a go file, it starts monitoring it.
Kinkku will also notice if you create a new go file in the directory or any of the subdirectories, and will start monitoring it.

4. ### Detects modifications
kinkku will be notified whenever any go file is modifed (bars).
When a modification is detected, a message will be printed in your terminal as so:

```bash
Go file modification detected:example/ham/burger.go
```
followed by kinkku shutting down any activity in the port you have provided and instantly restarting the server.

This will allow you to bascially have a live preview of your go server.

It's simple, all you need to do is save and refresh the page.


Keep in mind that whenever you save a go file, even if you haven't changed anything in the code, it registers as a modification.
Which means that if you have auto-save turned on you might get flooded with messages in your terminal. I guess I could do a flag that turns off the messages  but can't bother yet. Just turn off your auto-save for now :)

5. ### It minds it's own business.

Kinkku will NOT interfere with any messages that any other program is printing in your console and will just do it's own thing independently.
This could come in handy whenever you are debugging, you can instantly (when you save) see if any messages are being printed out in the terminal.

## Thats it!
If you end up using this tool, please let me know! It would make my day :) and also feel free to contact me if you have any suggestions.

Discord: headhilly

## Uninstall:

To uninstall kinkku, simply use the command:

```bash
sudo rm /usr/local/bin/kinkku
```
