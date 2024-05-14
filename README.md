
# Kinkku

Kinkku is a simple command line tool that simulates a live preview of your golang server by restarting it whenever a go file is modified.
I was annoyed by having to shut then run the server manually over and over again, so I made kinkku to make working on my school projects more convenient.

ps. Don't judge my code. this was intended to be a tool just for myself and didn't care much for good practices :)

### This version has only been tested on Linux for now

## Installation

1. **Clone Repository**:

Clone this repository to your local machine using the following command:

```bash
git clone https://github.com/headhilly/kinkku.git
```

2. **Navigate to Cloned Directory**:


```bash
cd kinkku
```

3. **Build and Install**:

Build the kinkku binary and move it to the `/usr/local/bin` directory using the following commands:

```bash
go build
```

Then


```bash
sudo mv kinkku /usr/local/bin
```

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

# All you need to do is save and refresh the page.


Keep in mind that whenever you save a go file, even if you haven't changed anything in the code, it registers as a modification.
Which means that if you have auto-save turned on you might get flooded with messages in your terminal. I guess I could do a flag that turns off the messages  but can't bother yet. Just turn off your auto-save for now :)

5. ### It minds it's own business.

Kinkku will NOT interfere with any messages that any other program is printing in your console and will just do it's own thing independently.
This could come in handy whenever you are debugging, you can instantly (when you save) see if any messages are being printed out in the terminal.


## Uninstall:

To uninstall kinkku, simply use the command:

```bash
sudo rm /usr/local/bin/kinkku
```

If you end up using this, please let me know! It would make my day :) and also feel free to contact me if you have any suggestions.