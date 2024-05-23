package kinkku

var (
	port                 = "8080"    // Specify the port your server is running on
	path                 = "../cars" // Specify the path to your project directory
	ModificationDetected = false     // Flag to track if any file modification has been detected
	restartCount         = 0
)

//style vars

var (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m"
	Hidden    = "\033[8m"

	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

var (
	banner = `
 _   __  _____   _   _   _   __  _   __  _   _                __        _____ 
| | / / |_   _| | \ | | | | / / | | / / | | | |              /  |      |  _  |
| |/ /    | |   |  \| | | |/ /  | |/ /  | | | |     __   __   | |      | |/| |
|    \    | |   |     | |    \  |    \  | | | |     \ \ / /   | |      | | | |
| |\  \  _| |_  | |\  | | |\  \ | |\  \ | |_| |      \ V /   _| |_  _  \ |_/ /
\_| \_/  \___/  \_| \_/ \_| \_/ \_| \_/  \___/        \_/    \___/ (_)  \___/ `
	slogan = `----------------------- Go do your thing, I Go do mine! -------------------------`
	noice  = ` _   _   _____   _____   _____   _____   _ 
| \ | | |  _  | |_   _| /  __ \ |  ___| | |
|  \| | | | | |   | |   | /  \/ | |__   | |
| .   | | | | |   | |   | |     |  __|  | |
| |\  | \ \_/ /  _| |_  | \__/\ | |___  |_|
\_| \_/  \___/   \___/   \____/ \____/  (_)	
`
)
