# kinkku v1.0

## Usage:

1. **Clone Repository**: Clone this repository to your local machine using the following command:

    ```bash
    git clone https://github.com/headhilly/kinkku.git
    ```

2. **Navigate to Project Directory**: Open a terminal and navigate to the directory `/kinkku`.

4. **Run the Program**: Execute the following command run the program, specifying your server's path and the port number it is running on:

    ```bash
    go run . /servers/myserver 8080
    ```

## Now Kinkku is running. Here's what it does:

1. It will start your server by performing `go run .` in the provided directory.
2. It will look for Go files in your server directory and start monitoring them for changes.
3. Whenever a modification is detected, it will automatically restart the server for you.
4. It will also notice any new Go files created and include them in the monitoring process.
